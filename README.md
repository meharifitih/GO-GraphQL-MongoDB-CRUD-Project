

# Golang GraphQL MongoDB CRUD Project

Welcome! This project demonstrates a full-stack CRUD API using Go, GraphQL (gqlgen), and MongoDB Atlas. You can use this as a reference for your own projects or follow along with the related YouTube tutorial.

---

## 🚀 Features
- Create, read, update, and delete job listings
- GraphQL Playground UI for easy API testing
- MongoDB Atlas integration for cloud database storage
- Clean, idiomatic Go code using gqlgen

---

## 🛠️ Quick Start

### 1. Project Setup

```sh

# Clone this repository and enter the project directory
git clone https://github.com/meharifitih/GO-GraphQL-MongoDB-CRUD-Project.git
cd GO-GraphQL-MongoDB-CRUD-Project

# Initialize a Go module (if not already initialized)
go mod init github.com/meharifitih/GO-GraphQL-MongoDB-CRUD-Project

# Install gqlgen
go get github.com/99designs/gqlgen

# Add gqlgen to tools.go for reproducible builds
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go

# Download all dependencies
go mod tidy

# Initialize gqlgen project structure
go run github.com/99designs/gqlgen init

# After writing your GraphQL schema, generate code
go run github.com/99designs/gqlgen generate

# Start the server
go run server.go
```

### 2. Configure MongoDB Atlas
- Update your MongoDB connection string in `database/database.go`.
- Make sure your cluster is running and your IP is whitelisted.

### 3. Open the Playground
- Visit [http://localhost:8080/](http://localhost:8080/) in your browser.

---

## 🧪 Example GraphQL Queries & Mutations

### Get All Jobs
```graphql
query GetAllJobs {
  jobs {
    _id
    title
    description
    company
    url
  }
}
```

---

### Create a Job
```graphql
mutation CreateJobListing($input: CreateJobListingInput!) {
  createJobListing(input: $input) {
    _id
    title
    description
    company
    url
  }
}
```
**Variables:**
```json
{
  "input": {
    "title": "Software Development Engineer - I",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
    "company": "Google",
    "url": "https://www.google.com/"
  }
}
```

---

### Get Job By ID
```graphql
query GetJob($id: ID!) {
  job(id: $id) {
    _id
    title
    description
    url
    company
  }
}
```
**Variables:**
```json
{
  "id": "638051d7acc418c13197fdf7"
}
```

---

### Update Job By ID
```graphql
mutation UpdateJob($id: ID!, $input: UpdateJobListingInput!) {
  updateJobListing(id: $id, input: $input) {
    _id
    title
    description
    company
    url
  }
}
```
**Variables:**
```json
{
  "id": "638051d3acc418c13197fdf6",
  "input": {
    "title": "Software Development Engineer - III"
  }
}
```

---

### Delete Job By ID
```graphql
mutation DeleteQuery($id: ID!) {
  deleteJobListing(id: $id) {
    deleteJobId
  }
}
```
**Variables:**
```json
{
  "id": "638051d3acc418c13197fdf6"
}
```

---

## ℹ️ Notes
- Replace the IDs in the variables with actual IDs from your database.
- All fields marked with `!` in the schema are required.
- Make sure your MongoDB Atlas cluster is running and accessible from your IP address.
- The GraphQL Playground is available at `/` (root URL) when the server is running.

---

## 📺 Video Tutorial
Check out the full walkthrough on YouTube: [Golang GraphQL MongoDB CRUD Project](#) *(insert your video link here)*

---

Feel free to fork, modify, and extend this project for your own needs!
