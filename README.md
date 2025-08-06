# Chirpy

This is a learning project for the golang backend path of the [Boot.dev](https_boot.dev) course.

## Description

Chirpy is a simple, Twitter-like social media application. Users can create an account, post "chirps" (short messages), and view chirps from other users.

## Features

*   **User authentication:** Users can sign up and log in to their accounts.
*   **Create and view chirps:** Authenticated users can post new chirps and see a timeline of all chirps.
*   **Simple and clean API:** The application provides a RESTful API for all its functionalities.

## Technologies Used

*   **Go:** The backend of the application is written in Go.
*   **PostgreSQL:** A PostgreSQL database is used to store user and chirp data.
*   **net/http:** The application uses the standard `net/http` package to handle HTTP requests.
*   **JSON Web Tokens (JWT):** JWTs are used for user authentication.

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

*   Go
*   PostgreSQL
*   [sqlc](https://github.com/sqlc-dev/sqlc)

### Installation

1.  Clone the repo
    ```sh
    git clone https://github.com/lucasrodlima/chirpy.git
    ```
2.  Install Go packages
    ```sh
    go mod tidy
    ```
3.  Set up the database
    *   Create a `.env` file in the root directory and add your PostgreSQL connection string:
        ```
        DB_URL="your_postgres_connection_string"
        ```
    *   Run the database migrations:
        ```sh
        sqlc generate
        ```
4.  Run the application
    ```sh
    go run .
    ```

## API Endpoints

The following are the main API endpoints provided by the application:

*   `POST /api/users`: Create a new user.
*   `POST /api/login`: Log in a user.
*   `POST /api/chirps`: Create a new chirp.
*   `GET /api/chirps`: Get all chirps.
*
## Project Structure

*   `main.go`: The entry point of the application.
*   `handler_*.go`: These files contain the HTTP handlers for the different API endpoints.
*   `internal/`: This directory contains the internal logic of the application, such as authentication and database access.
*   `sql/`: This directory contains the SQL schema and queries for the database.

