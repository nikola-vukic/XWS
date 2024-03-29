package startup

import (
	"accommodation_booking/common/auth"
	"accommodation_booking/common/client"
	accommodation "accommodation_booking/common/proto/accommodation_service"
	grade "accommodation_booking/common/proto/grade_service"
	profile "accommodation_booking/common/proto/profile_service"
	reservation "accommodation_booking/common/proto/reservation_service"
	saga "accommodation_booking/common/saga/messaging"
	"accommodation_booking/common/saga/messaging/nats"
	"accommodation_booking/grade_service/application"
	"accommodation_booking/grade_service/domain"
	"accommodation_booking/grade_service/infrastructure/api"
	"accommodation_booking/grade_service/infrastructure/persistence"
	"accommodation_booking/grade_service/startup/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {
	const gradeServicePath = "/grade.GradeService/"

	return map[string][]string{
		gradeServicePath + "GetAll": {"unexposed"},
		gradeServicePath + "Get":    {"unexposed"},
		gradeServicePath + "Create": {"guest"},
		gradeServicePath + "Update": {"guest"},
		gradeServicePath + "Delete": {"guest"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	gradeStore := server.initGradeStore(mongoClient)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	accommodationClient, err := client.NewAccommodationClient(fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	reservationClient, err := client.NewReservationClient(fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	gradeService := server.initGradeService(gradeStore)
	gradeHandler := server.initGradeHandler(gradeService, profileClient, accommodationClient, reservationClient)

	server.startGrpcServer(gradeHandler, jwtManager)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.GradeDBHost, server.config.GradeDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initGradeHandler(service *application.GradeService, profileClient profile.ProfileServiceClient, accommodationClient accommodation.AccommodationServiceClient, reservationClient reservation.ReservationServiceClient) *api.GradeHandler {
	return api.NewGradeHandler(service, profileClient, accommodationClient, reservationClient)
}

func (server *Server) initGradeStore(client *mongo.Client) domain.GradeStore {
	store := persistence.NewGradeMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, Grade := range grades {
		grade, err := store.Create(context.TODO(), Grade)
		if err != nil {
			log.Fatal(err)
			log.Println(grade)
		}
	}
	return store
}

func (server *Server) initGradeService(store domain.GradeStore) *application.GradeService {
	return application.NewGradeService(store)
}

func (server *Server) startGrpcServer(gradeHandler *api.GradeHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(serverOptions...)
	grade.RegisterGradeServiceServer(grpcServer, gradeHandler)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
