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
	jobCollection := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err := cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}
	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := jobCollection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"url":         jobInfo.URL,
		"company":     jobInfo.Company,
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := inserted.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{
		ID:          insertedID,
		Title:       jobInfo.Title,
		Company:     jobInfo.Company,
		Description: jobInfo.Description,
		URL:         jobInfo.URL,
	}

	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobID string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobListing := bson.M{}
	if jobInfo.Title != nil {
		updateJobListing["title"] = jobInfo.Title
	}

	if jobInfo.Description != nil {
		updateJobListing["description"] = jobInfo.Description
	}

	if jobInfo.URL != nil {
		updateJobListing["url"] = jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobID)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobListing}

	result := jobCollection.FindOneAndUpdate(ctx, filter, update,
		options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing
	if err := result.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing
}

func (db *DB) DeleteJobListing(jobID string) *model.DeleteJobListingResponse {
	jobCollection := db.client.Database("graphql-job-board").Collection("job")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobID)
	filter := bson.M{"_id": _id}

	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return &model.DeleteJobListingResponse{DeleteJobID: jobID}
}
