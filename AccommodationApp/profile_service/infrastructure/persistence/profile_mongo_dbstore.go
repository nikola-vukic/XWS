package persistence

import (
	"accommodation_booking/profile_service/domain"
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	DATABASE   = "profile_service"
	COLLECTION = "profile"
)

type ProfileMongoDBStore struct {
	profiles *mongo.Collection
}

func NewProfileMongoDBStore(client *mongo.Client) domain.ProfileStore {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	_, err := profiles.Indexes().DropOne(context.TODO(), "fullName_text")
	index := mongo.IndexModel{
		Keys: bson.D{{"fullName", "text"}},
		//Options: options.Index().SetUnique(true),
	}
	opts := options.CreateIndexes().SetMaxTime(20 * time.Second)
	_, err = profiles.Indexes().CreateOne(context.TODO(), index, opts)

	if err != nil {
		panic(err)
	}

	return &ProfileMongoDBStore{
		profiles: profiles,
	}
}

func (store *ProfileMongoDBStore) Get(ctx context.Context, profileId string) (*domain.Profile, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ProfileMongoDBStore) GetAll(ctx context.Context, search string) ([]*domain.Profile, error) {
	filter := bson.D{{"fullName", bson.M{"$regex": "^.*" + search + ".*$"}}}
	return store.filter(filter, search)
}

func (store *ProfileMongoDBStore) Create(ctx context.Context, profile *domain.Profile) error {
	result, err := store.profiles.InsertOne(ctx, profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ProfileMongoDBStore) Update(ctx context.Context, profileId string, profile *domain.Profile) error {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return err
	}
	result, err := store.profiles.ReplaceOne(
		ctx,
		bson.M{"_id": id},
		profile,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(profile.Id.String())
	}
	return nil
}

func (store *ProfileMongoDBStore) DeleteAll(ctx context.Context) error {
	_, err := store.profiles.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ProfileMongoDBStore) filter(filter interface{}, search string) ([]*domain.Profile, error) {
	cursor, err := store.profiles.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, errors.New(search)
	}
	return decode(cursor)
}

func (store *ProfileMongoDBStore) filterOne(filter interface{}) (profile *domain.Profile, err error) {
	result := store.profiles.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func (store *ProfileMongoDBStore) Delete(ctx context.Context, id string) error {
	profileId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.profiles.DeleteOne(ctx, bson.M{"_id": profileId})
	if err != nil {
		return err
	}
	return nil
}

func decode(cursor *mongo.Cursor) (profiles []*domain.Profile, err error) {
	for cursor.Next(context.TODO()) {
		var Profile domain.Profile
		err = cursor.Decode(&Profile)
		if err != nil {
			return
		}
		profiles = append(profiles, &Profile)
	}
	err = cursor.Err()
	return
}

func (store *ProfileMongoDBStore) GetByToken(ctx context.Context, token string) (*domain.Profile, error) {
	filter := bson.M{"agentToken": token}
	return store.filterOne(filter)
}

func (store *ProfileMongoDBStore) GenerateToken(ctx context.Context, id primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": id}
	profile, err := store.filterOne(filter)
	if err != nil {
		return "", err
	}
	token := getToken(10)
	profile.Token = token
	err = store.Update(ctx, profile.Id.Hex(), profile)
	if err != nil {
		return "", err
	}
	return token, nil
}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
