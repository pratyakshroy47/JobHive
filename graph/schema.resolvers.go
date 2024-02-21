// resolver.go
package graph

import (
	"errors"

	"github.com/pratyakshroy47/gql-go/database"
	"github.com/pratyakshroy47/gql-go/graph/model"
	"github.com/pratyakshroy47/gql-go/logger"
)

type Resolver struct {
	DB *database.JobRepository
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateJobListing(ctx Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	jobListing, err := r.DB.CreateJobListing(input)
	if err != nil {
		logger.Error("Failed to create job listing", err)
		return nil, err
	}

	return jobListing, nil
}

func (r *mutationResolver) UpdateJobListing(ctx Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	jobListing, err := r.DB.UpdateJobListing(id, input)
	if err != nil {
		logger.Error("Failed to update job listing", err)
		return nil, err
	}

	if jobListing == nil {
		return nil, errors.New("job listing not found")
	}

	return jobListing, nil
}

func (r *mutationResolver) DeleteJobListing(ctx Context, id string) (*model.DeleteJobResponse, error) {
	response, err := r.DB.DeleteJobListing(id)
	if err != nil {
		logger.Error("Failed to delete job listing", err)
		return nil, err
	}

	if response == nil {
		return nil, errors.New("job listing not found")
	}

	return response, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Jobs(ctx Context, limit *int64, offset *int64) ([]*model.JobListing, error) {
	l := maxDocumentLimit
	o := int64(0)

	if limit != nil {
		l = *limit
	}

	if offset != nil {
		o = *offset
	}

	jobListings, err := r.DB.GetJobs(l, o)
	if err != nil {
		logger.Error("Failed to fetch job listings", err)
		return nil, err
	}

	return jobListings, nil
}

func (r *queryResolver) Job(ctx Context, id string) (*model.JobListing, error) {
	jobListing, err := r.DB.GetJob(id)
	if err != nil {
		logger.Error("Failed to fetch job listing", err)
		return nil, err
	}

	if jobListing == nil {
		return nil, errors.New("job listing not found")
	}

	return jobListing, nil
}