package repository

import (
	"FastGourmet/internal/core/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoDbRepository struct {
	client *mongo.Client
}

func NewMongoDbRepository(client *mongo.Client) *MongoDbRepository {
	return &MongoDbRepository{client: client}
}

func (m MongoDbRepository) Save(order *domain.Order) error {
	//TODO implement me
	collection := m.client.Database("FastGourmet").Collection("orders")
	result, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	log.Println("Inserted a single document: ", result.InsertedID)
	return nil
}

func (m MongoDbRepository) List() ([]*domain.Order, error) {
	filter := bson.D{{}}
	collection := m.client.Database("FastGourmet").Collection("orders")
	var orders []*domain.Order
	ctx := context.TODO()
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return orders, err
	}
	for cur.Next(ctx) {
		var order domain.Order
		err := cur.Decode(&order)
		if err != nil {
			return orders, err
		}
		orders = append(orders, &order)
	}
	if err := cur.Err(); err != nil {
		return orders, err
	}
	cur.Close(ctx)
	return orders, nil
}

func (m MongoDbRepository) Get(id string) (*domain.Order, error) {
	//TODO implement me
	filter := bson.D{{"id", id}}
	collection := m.client.Database("FastGourmet").Collection("orders")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var results []domain.Order
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return &results[0], nil

}

func (m MongoDbRepository) Update(id string, order *domain.Order) error {
	filter := bson.D{{"id", id}}
	collection := m.client.Database("FastGourmet").Collection("orders")
	update := bson.D{{"$set", bson.D{{"priority", order.Priority}, {"status", order.Status}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
