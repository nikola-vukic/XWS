package startup

import (
	"accommodation_booking/common/auth"
	user "accommodation_booking/common/proto/user_service"
	"accommodation_booking/user_service/application"
	"accommodation_booking/user_service/domain"
	"accommodation_booking/user_service/infrastructure/api"
	"accommodation_booking/user_service/infrastructure/persistence"
	"accommodation_booking/user_service/startup/config"
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
	const userServicePath = "/user.UserService/"

	return map[string][]string{
		userServicePath + "GetAll": {"*"},
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)
	userService := server.initUserService(userStore)

	jwtManager := auth.NewJWTManager("secretKey", 60*time.Minute)

	userHandler := server.initUserHandler(userService, jwtManager)

	server.startGrpcServer(userHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatalf("MGF: %v", err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	err := store.DeleteAll(context.TODO())
	if err != nil {
		return nil
	}
	for _, User := range users {
		_, err := store.Register(context.TODO(), User)
		if err != nil {
			log.Fatalf("RUF: %v", err)
		}
	}

	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService,
	jwtManager *auth.JWTManager) *api.UserHandler {
	return api.NewUserHandler(service, jwtManager)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler, jwtManager *auth.JWTManager) {
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
	user.RegisterUserServiceServer(grpcServer, userHandler)
	reflection.Register(grpcServer)
	//serveMux := http.NewServeMux()
	//serveMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
	//	promhttp.Handler().ServeHTTP(w, r)
	//})
	//err = http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), serveMux)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//	return
	//}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}