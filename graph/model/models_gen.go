// models_gen.go

package model

type CreateJobListingInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type DeleteJobResponse struct {
	DeleteJobID string `json:"deleteJobId"`
}

type JobListing struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateJobListingInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Company     *string `json:"company,omitempty"`
	URL         *string `json:"url,omitempty"`
}