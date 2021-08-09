# Go REST API Example
A lightweight example API written in go. Orginally designed to help people practice deploying an API to a Kubernetes cluster. 

Note:  that the playbooks in this repository are for demonstration and learning. They are not intended for production setups.


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
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /users
* `GET` : Get all users

#### /user/{#username}
* `GET` : Get a specific user based on their username 

#### /comments 
* `GET` : Get each user's latest comment 

#### /health
* `GET` : Get the healthcheck endpoint, `OK` 

## Todo

- [ ] Organize the code with packages
- [ ] Write documentation on how to find and use 
- [ ] Write the tests for all endpoints.
- [ ] Build a deployment process to push Docker image to a container registry 