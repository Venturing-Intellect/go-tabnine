Title: Customer Feedback Application

## Table of Contents
1. Introduction
2. Installation
3. Usage
4. Components Diagram
5. API Endpoints
6. Contributing
7. License

## 1. Introduction

This is a simple customer feedback application built using Go and React. The application allows users to submit feedback through a web form, which is then stored in a PostgreSQL database.

## 2. Installation

To set up the application, follow these steps:

### Backend:

1. Install Go: Download and install the latest version of Go from https://golang.org/dl/.
2. Clone the backend repository: `git clone git@github.com:Venturing-Intellect/go-tabnine.git`.
3. Navigate to the backend directory: `cd go-tabnine`.
4. Install dependencies: `go mod tidy && go mod download`.
5. Create a `.env` file in the backend directory and add the following environment variables:
   - `DB_USER`: PostgreSQL database user
   - `DB_PASSWORD`: PostgreSQL database password
   - `DB_NAME`: PostgreSQL database name
   - `DB_HOST`: PostgreSQL database host (default: localhost)
   - `DB_PORT`: PostgreSQL database port (default: 5432)
6. Build and run the application: `docker compose up --build`.

### Frontend:

1. Install Node.js: Download and install the latest version of Node.js from https://nodejs.org/.
2. Clone the frontend repository: `git clone git@github.com:Venturing-Intellect/go-tabnine.git`.
3. Navigate to the frontend directory: `cd go-tabnine/frontend`.
4. Install dependencies: `npm install`.
5. Start the frontend server: `npm start`.

## 3. Usage

To use the application, open your web browser and navigate to http://localhost:3000. You will see the customer feedback form. Fill in the required fields and submit the form to store the feedback in the PostgreSQL database.

## 4. Components Diagram

Here is a diagram of the main components of the customer feedback application:

```
+------------+     +-------------+     +---------------+
| React App  |     | Go API      |     | PostgreSQL DB |
+------------+     +-------------+     +---------------+
     | POST /feedback |
     |----------------|
     |               |
     |               |
     |               |
     |               |
     |               |
     |               |
     |               |
     |               |
     |               |
     +-------------+
```

## 5. API Endpoints

The customer feedback application has the following API endpoint:

- POST /feedback: Submits a new customer feedback.

## 6. Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request on the respective GitHub repositories.

## 7. License

This project is licensed under the MIT License. See the LICENSE file for more details.