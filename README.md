# API Documentation

[![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)](https://app.swaggerhub.com/apis-docs/campyukuser/technical-test/1.0) [![Postman Collection](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://documenter.getpostman.com/view/19389812/2s93JtRPqh)

# Project structure

    .
    ├── config
    │   └── config.go
    ├── docs
    │   ├── app.env.example
    │   └── swagger.yaml
    ├── internal
    │   ├── catalogue
    │   │   ├── handler
    │   │   ├── repository
    │   │   ├── service
    │   │   └── entity.go
    │   └── schedule
    │   │   ├── handler
    │   │   ├── repository
    │   │   ├── service
    │   │   └── entity.go
    ├── pkg // Third party libraries,  utilities, and databases.
    │   ├── cloud
    │   │   └── cloudinary.go
    │   ├── database
    │   │   └── mysql.go
    │   └── helper
    ├── README.md
    ├── app.env
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── script.sql

</div>
