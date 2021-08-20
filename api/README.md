# Go API Example
A lightweight example API written in go that was designed to help people practice deploying an API to a Kubernetes cluster. 

## Installation & Run
```bash
# Download this project
go get https://github.com/do-community/example-k8s-workloads/
```

```bash
# Build and Run
cd api
go mod init 
go run main.go

# API Endpoint : http://127.0.0.1:4000
```

## Structure
```
├── api
|   └── vendor
│   └── users
│       └── users.json // Sample data for the application
└── main.go
```

## API

#### api/v1/users
* `GET` : Get all users

#### api/v1/user/{#username}
* `GET` : Get a specific user with their username 

#### api/v1/comments 
* `GET` : Get each user's latest comment 

#### api/v1/health
* `GET` : Get the healthcheck endpoint, `OK` 

## Todo
- [ ] Add yaml manifest for Kubernetes Deployment 
- [ ] Write the tests for all endpoints
- [ ] Build a deployment process to push Docker image to a container registry 
