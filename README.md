
# Click Counter Service (Pre-interview test task)

This is a simple API service built in Go that allows counting and retrieving statistics of clicks on banners. The service keeps track of clicks in a PostgreSQL database and supports querying click statistics for banners over specified time ranges.

## Features

- **Count Clicks**: Increment a click counter for a banner when a request is made to the `/counter/{bannerID}` endpoint.
- **Retrieve Stats**: Retrieve banner click statistics within a specific time range using the `/stats/{bannerID}` endpoint.

---

## Table of Contents

- [Installation](#installation)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [How to Run the Application](#how-to-run-the-application)
- [Testing](#testing)
- [Error Handling](#error-handling)
- [Contributing](#contributing)

---

## Installation

To install and run the service locally, follow these steps:

### Prerequisites

- **Go** (1.18+)
- **PostgreSQL** (for database storage)
- **Curl** (for making requests from the command line)

### Clone the repository

```bash
git clone https://github.com/yourusername/click-counter.git
cd click-counter
```

### Install Dependencies

Install Go dependencies:

```bash
go mod tidy
```

### Set up PostgreSQL

1. Create a PostgreSQL database and user.
2. Create a table for storing clicks:

```sql
CREATE TABLE clicks (
  banner_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
```

3. Update the PostgreSQL connection string in `config.yaml` or environment variables for your local setup.

---

## API Endpoints

### 1. **GET /counter/{bannerID}** - Increment Click Counter

- **Description**: This endpoint increments the click counter for a given banner. When the endpoint is hit, a new click is logged for the specified `bannerID`.
- **Method**: `GET`
- **URL**: `/counter/{bannerID}`
- **Response**: Returns `HTTP 200 OK` if the click is successfully logged.

#### Example Request:

```bash
curl -X GET "http://localhost:8080/api/v1/counter/634dfb86-f492-4f12-b524-2b3d35f2c5a3"
```

#### Example Response:

```json
{
  "status": 201,
  "message": "Click successfully generated",
  "payload": null
}
```

---

### 2. **POST /stats/{bannerID}** - Get Click Stats in a Time Range

- **Description**: This endpoint retrieves the click statistics for a specific banner within a given time range.
- **Method**: `POST`
- **URL**: `/stats/{bannerID}`
- **Body**:
  - `tsFrom`: Start timestamp (in `ISO 8601` format, e.g., `2024-11-12T00:00:00Z`).
  - `tsTo`: End timestamp (in `ISO 8601` format, e.g., `2024-11-14T00:00:00Z`).
  
- **Response**: Returns a JSON object with an array of clicks recorded for the banner within the specified time range.

#### Example Request:

```bash
curl -X POST "http://localhost:8080/api/v1/stats/634dfb86-f492-4f12-b524-2b3d35f2c5a3" \
        -H "Content-Type: application/json"      
        -d '{
           "tsFrom": "2024-11-12T00:00:00Z", 
           "tsTo": "2024-11-14T00:00:00Z"
         }'
```

#### Example Response:

```json
{
  "status": 200,
  "message": "Click successfully",
  "payload": {
    "count": 3545,
    "data": [
      {
        "created_at": "2024-11-13T17:07:48.846793Z",
        "banner_id": "634dfb86-f492-4f12-b524-2b3d35f2c5a3"
      },
      {
        "created_at": "2024-11-13T17:07:48.848502Z",
        "banner_id": "634dfb86-f492-4f12-b524-2b3d35f2c5a3"
      },
      {
        "created_at": "2024-11-13T17:07:48.849619Z",
        "banner_id": "634dfb86-f492-4f12-b524-2b3d35f2c5a3"
      },
      ...
    ]
  }
}
```

---

## Database Schema

The database stores click data in a table named `clicks`. The schema consists of the following fields:

- `id`: **Primary Key** – Auto-incremented identifier for the click.
- `banner_id`: **VARCHAR(255)** – The ID of the banner.
- `created_at`: **TIMESTAMP** – The timestamp when the click occurred.

```sql
CREATE TABLE clicks (
  id SERIAL PRIMARY KEY,
  banner_id VARCHAR(255),
  created_at TIMESTAMP
);
```

---

## How to Run the Application

### Running Locally

1. Set up your PostgreSQL database as described in the **Installation** section.
2. Make sure the necessary environment variables are set for the database connection:
   - `DB_HOST`: Database host.
   - `DB_PORT`: Database port.
   - `DB_USER`: Database user.
   - `DB_PASSWORD`: Database password.
   - `DB_NAME`: Database name.
3. Run the Go application:

```bash
go run cmd/migration/main.go -up # runs migration up
go run cmd/server/main.go 
```

4. The server should now be running on `http://localhost:8080`.

---

## Testing

You can use `curl` to test the API endpoints.

### To Record a Click:

```bash
curl -X GET "http://localhost:8080/api/v1/counter/634dfb86-f492-4f12-b524-2b3d35f2c5a3"
```

### To Get Click Statistics:

```bash
curl -X POST "http://localhost:8080/api/v1/stats/634dfb86-f492-4f12-b524-2b3d35f2c5a3" \     
          -H "Content-Type: application/json" \     
          -d '{
               "tsFrom": "2024-11-12T00:00:00Z", 
               "tsTo": "2024-11-14T00:00:00Z"
             }'
```

---

## Error Handling

The API returns standard HTTP error codes for common errors:

- **400 Bad Request**: Invalid input or missing parameters.
- **404 Not Found**: Banner ID not found.
- **500 Internal Server Error**: General error when processing the request.

Example error response:

```json
{
  "error": "Invalid request payload"
}
```

## Load Testing with k6

[k6](https://k6.io/) is a modern load testing tool that allows you to simulate traffic against your web applications. You can use k6 to test how well the **Click Counter Service** performs under load, and to ensure the service can handle the expected traffic.

### Prerequisites

Before running the k6 tests, make sure you have one of the following:

- **Docker**: If you prefer not to install k6 locally.
- **k6 installed**: If you prefer to run it directly on your system.

### Running k6 without Installing (Docker)

If you don't want to install k6 on your local machine, you can use Docker to run it as a container. Ensure that Docker is installed and running on your machine.

1. Open a terminal and run the following command to execute your test script:

```bash
   k6 run k6/counter_get.js
```

```js
import http from 'k6/http';

export const options = {
    vus: 500,              // Virtual users to simulate
    duration: '30s',       // Duration of the test
    rps: 500,              // Requests per second (limit to 500 RPS)
};

export default function () {
    const bannerId = "634dfb86-f492-4f12-b524-2b3d35f2c5a3";
    const url = `http://localhost:8080/api/v1/counter/${bannerId}`;

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.get(url, params); // Sending GET request to the counter endpoint
}

```

This script configures 500 virtual users to simulate traffic over a duration of 30 seconds, sending 500 requests per second to the counter API.

---

## Contributing

Feel free to fork this repository and create pull requests. If you have any suggestions or improvements, please submit an issue or open a pull request. Make sure to follow proper commit message conventions.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
