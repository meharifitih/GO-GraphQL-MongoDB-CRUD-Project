package database

import (
	"context"
	"log"
	"time"

	"github.com/meharifiti/graphQL-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString = "mongodb+srv://meharifitih:XBTMG8Em14Ih8fI7@cluster0.rdeuiwm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	var jobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var jobListings []*model.JobListing
	return jobListings
}

func (db *DB) CreateJobListing(input model.CreateJobListingInput) *model.JobListing {
	return &model.JobListing{}
}

func (db *DB) UpdateJobListing(jobID string, input model.UpdateJobListingInput) *model.JobListing {
	return &model.JobListing{}
}

func (db *DB) DeleteJobListing(jobID string) *model.DeleteJobListingResponse {
	return &model.DeleteJobListingResponse{DeleteJobID: jobID}
}
