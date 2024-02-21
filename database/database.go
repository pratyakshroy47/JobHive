package database

import (
	"context"
	"errors"
	"time"

	"github.com/pratyakshroy47/gql-go/graph/model"
	"github.com/pratyakshroy47/gql-go/mongo" // Corrected import path
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName     = "graphql-job-board"
	collectionName   = "jobs"
	defaultTimeout   = 10 * time.Second
	maxDocumentLimit = 100
)

type JobRepository struct {
	db *mongo.Database
}

func NewJobRepository(client *mongo.Client) *JobRepository {
	return &JobRepository{
		db: client.Database(databaseName),
	}
}

func (r *JobRepository) GetJob(id string) (*model.JobListing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var jobListing model.JobListing

	err = r.db.Collection(collectionName).FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &jobListing, nil
}

func (r *JobRepository) GetJobs(limit, offset int64) ([]*model.JobListing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	opts := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := r.db.Collection(collectionName).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var jobListings []*model.JobListing
	if err = cursor.All(ctx, &jobListings); err != nil {
		return nil, err
	}

	return jobListings, nil
}

func (r *JobRepository) CreateJobListing(input model.CreateJobListingInput) (*model.JobListing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	jobListing := model.JobListing{
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         input.URL,
	}

	result, err := r.db.Collection(collectionName).InsertOne(ctx, jobListing)
	if err != nil {
		return nil, err
	}

	jobListing.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return &jobListing, nil
}

func (r *JobRepository) UpdateJobListing(id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{}

	if input.Title != nil {
		update["title"] = *input.Title
	}

	if input.Description != nil {
		update["description"] = *input.Description
	}

	if input.Company != nil {
		update["company"] = *input.Company
	}

	if input.URL != nil {
		update["url"] = *input.URL
	}

	if len(update) == 0 {
		return nil, errors.New("no fields to update")
	}

	updateOpts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var jobListing model.JobListing

	err = r.db.Collection(collectionName).FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, updateOpts).Decode(&jobListing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &jobListing, nil
}

func (r *JobRepository) DeleteJobListing(id string) (*model.DeleteJobResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	result, err := r.db.Collection(collectionName).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, nil
	}

	return &model.DeleteJobResponse{DeleteJobID: id}, nil
}