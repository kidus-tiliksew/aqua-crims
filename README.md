# Aqua Cloud Resource Inventory Management System

Aqua Cloud Resource Inventory Management System is a service for managing customers, cloud resources, and notifications. It provides both HTTP and gRPC APIs.

## Features

- Manage customers and their cloud resources.
- Send and manage notifications.
- Seed initial data using a CLI command.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/kidus-tiliksew/aqua-crims.git
   cd aqua-crims
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Create a `.env` file with the following environment variables:
   ```env
   DATABASE_DSN=host=0.0.0.0 user=postgres password=password dbname=aqua-crims
   AMPQ_DSN=your_ampq_dsn
   PORT=8080
   GRPC_PORT=9090
   ```

## Usage

### Running the HTTP and gRPC Servers

To start the HTTP and gRPC servers, run:

```sh
go run .
```

### Seeding the Database

To seed the database with initial cloud resources, run:

```sh
go run cmd/cli.go seed
```

## HTTP API Endpoints

- `POST /customers`: Create a new customer.
- `POST /customers/:id/cloud-resources`: Create cloud resources for a customer.
- `GET /customers/:id/cloud-resources`: Get cloud resources for a customer.
- `GET /customers/:id/notifications`: Get notifications for a customer.
- `DELETE /customers/:id/notifications`: Delete notifications for a customer.
- `POST /cloud-resources`: Create a new cloud resource.
- `PUT /cloud-resources/:id`: Update a cloud resource.
- `DELETE /cloud-resources/:id`: Delete a cloud resource.
- `DELETE /notifications/:id`: Delete a notification.

## gRPC API

The gRPC server provides the following methods:

- `DeleteNotification`: Deletes a notification by ID.
- `DeleteNotificationByUser`: Deletes all notifications for a user.
- `GetNotificationsByUser`: Get's all notifications for a user.

## Design 

The core application follows the principles of Clean Architecture, also known as Hexagonal Architecture or Onion Architecture. Emphasizing separation of concerns and dependency inversion, making the application more modular, testable, and maintainable. 

## Structure

The file structure follows a CQRS pattern.

### Commands:

- **Purpose**: Encapsulate all the write operations (commands) of the application.
- **Files**: `customer_create.go`, `cloud_resource_create.go`, `notification_create.go`, etc.
- **Handlers**: Each command has a corresponding handler that contains the business logic for executing the command.

### Queries:

- **Purpose**: Encapsulate all the read operations (queries) of the application.
- **Files**: `cloud_resource_get.go`, `notification_get_by_user.go`, etc.
- **Handlers**: Each query has a corresponding handler that contains the business logic for executing the query.

### Application:

- **Purpose**: Acts as a facade that exposes the commands and queries to the rest of the application.
- **Files**: `application.go`
- Components:
   - **Commands Interface**: Defines all the command methods.
   - **Queries Interface**: Defines all the query methods.
   - **Application Struct**: Implements the App interface by embedding appCommands and appQueries.