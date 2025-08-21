package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://meharifitih:XBTMG8Em14Ih8fI7@cluster0.rdeuiwm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
}
