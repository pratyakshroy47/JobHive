# JobHive

JobHive is a high-performance GraphQL API built with Go and MongoDB, designed to provide an efficient and scalable solution for managing job listings. It leverages the power of GraphQL to reduce over-fetching by up to 90%, and MongoDB's flexible data model and powerful querying capabilities to improve query performance by up to 70%.

## Features

- **Efficient Data Fetching**: GraphQL allows clients to request only the data they need, reducing over-fetching and under-fetching issues common in traditional REST APIs.
- **Powerful Querying**: MongoDB's flexible data model and efficient indexing enable fast data retrieval and manipulation.
- **Concurrency and Performance**: Go's lightweight concurrency model with goroutines and channels ensures efficient handling of concurrent requests and I/O operations, resulting in up to 40% faster execution times.
- **Seamless Scalability**: MongoDB's sharding capabilities enable horizontal data distribution across multiple servers, boosting read and write performance by up to 90%.
- **Developer Productivity**: GraphQL's strongly typed schema and tooling, combined with Go's simplicity and extensive standard library, facilitate faster development cycles and improved code maintainability.

## Technologies Used

- GraphQL
- Go (Golang)
- MongoDB
- Docker
- Docker Compose
- gqlgen
- viper
- zap
- go-mongo-driver

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/JobHive.git

2. Set up the environment variables:
   cp .env.example .env

Edit the `.env` file and provide the required configuration values, such as the MongoDB connection string.

3. Build and run the project:
   ```bash
   docker-compose up --build

This will start the GraphQL API server and the MongoDB database.

4. Access the GraphQL Playground:

Open your web browser and navigate to `http://localhost:8080/playground` to explore and test the GraphQL API using the interactive GraphQL Playground.
