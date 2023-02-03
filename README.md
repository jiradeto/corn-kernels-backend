# CornKernels



## Overview
CornKernels is a lightweight inventory management system that provides a comprehensive solution for tracking stock.


## Features
- Inventory Management
- Stock Management
- REST API for CRUD operation on the inventory (corn kernels)
- REST API for CRUD operation on the stock movement

# Getting Started

## Building and Running

### 1. Prerequisite
If it's your first time running this application, you need to setup app configuration with:
```
make setup
```
This will setup `.env` and configure docker network used by all gh-scanner containers for you.


## Development
To run application locally, you need to first install project Go dependencies. At the root of project, run:
```
go mod download
```

Then you need to start start API Service with command:
```
go run cmd/main.go
```

## Project Structures
The project is made easy for maintenance and for future changes by implementing clean architecture where the code is organized as a layer with a specific responsibility. The layer in this project can be defined from the innermost layer to the outermost layer as follows:
- Presenter (`app/presenters`): receive incoming requests and pass them to use case, format response sent by use case to end user
- Use cases (`app/usecases`): handle code business logic
- Entities (`app/entities`): define the data model used in the app and connect to the database