/blog-app
├── /cmd
│   ├── /user-service
│   │   └── main.go                # Entry point for User Service
│   ├── /post-service
│   │   └── main.go                # Entry point for Post Service
│   ├── /comment-service
│   │   └── main.go                # Entry point for Comment Service
│   └── /auth-service
│       └── main.go                # Entry point for Auth Service
├── /internal
│   ├── /user-service
│   │   ├── /handlers
│   │   ├── /models
│   │   ├── /repository
│   │   └── /services
│   ├── /post-service
│   │   ├── /handlers
│   │   ├── /models
│   │   ├── /repository
│   │   └── /services
│   ├── /comment-service
│   │   ├── /handlers
│   │   ├── /models
│   │   ├── /repository
│   │   └── /services
│   └── /auth-service
│       ├── /handlers
│       ├── /models
│       ├── /repository
│       └── /services
├── /pkg
│   ├── /config                  # Shared config between services
│   ├── /middleware               # Common middleware (e.g., logging, CORS, etc.)
│   ├── /utils                    # Shared utility functions (e.g., error handling)
├── /docker
│   ├── /user-service.Dockerfile  # Dockerfile for User Service
│   ├── /post-service.Dockerfile  # Dockerfile for Post Service
│   ├── /comment-service.Dockerfile # Dockerfile for Comment Service
│   └── /auth-service.Dockerfile  # Dockerfile for Auth Service
├── /k8s                          # Kubernetes manifests for deployments and services
│   ├── user-service-deployment.yaml
│   ├── post-service-deployment.yaml
│   ├── comment-service-deployment.yaml
│   └── auth-service-deployment.yaml
├── /go.mod                       # Go modules file
├── /go.sum                       # Go modules checksum file
└── /README.md                    # Documentation
