# DevPulse Backend 🚀

DevPulse is a visual monitoring tool designed for developers to keep track of their cloud-deployed projects' health and performance.

This repository contains the **Go-based backend API** that handles project management, automated health checks, and statistics aggregation.

## ✨ Features

- **Automated Monitoring**: Background workers perform health checks every minute.
- **Project Management**: CRUD operations to manage your monitored applications.
- **Latency Tracking**: Records and reports latency for each health check.
- **Uptime Calculation**: Aggregates data to provide 24h uptime statistics.
- **CORS Support**: Configured for seamless integration with the DevPulse frontend.

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/) (v1.23+)
- **Framework**: [Fiber v2](https://gofiber.io/)
- **Database**: [PostgreSQL](https://www.postgresql.org/) with [SQLx](https://github.com/jmoiron/sqlx)
- **Environment**: [Godotenv](https://github.com/joho/godotenv)
- **ID Generation**: [Google UUID](https://github.com/google/uuid)

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher installed.
- A running PostgreSQL instance.

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/me-dev-pulse/backend.git
   cd backend
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up Environment Variables**:
   Create a `.env` file in the root directory (use the provided `.env` as a template):
   ```env
   DATABASE_URL=postgresql://user:password@localhost:5432/devpulse?sslmode=disable
   ```

4. **Run the application**:
   ```bash
   go run src/main.go
   ```
   The server will start on `http://localhost:3000`.

## 🛣️ API Routes

| Method   | Endpoint                | Description                      |
| -------- | ----------------------- | -------------------------------- |
| `GET`    | `/api/health`           | API Health Check                 |
| `GET`    | `/api/projects`         | List all monitored projects      |
| `POST`   | `/api/projects`         | Add a new project to monitor     |
| `DELETE` | `/api/projects/:id`     | Remove a project                 |
| `GET`    | `/api/projects/:id/stats` | Get statistics for a project     |
| `GET`    | `/api/projects/summary` | Get a summary of all projects    |

## 📂 Project Structure

```text
.
├── src/
│   ├── database/    # DB connection and repository logic
│   ├── handlers/    # HTTP controllers (Fiber)
│   ├── models/      # Data structures and DTOs
│   ├── services/    # Business logic and background monitor
│   └── main.go      # Entry point and route definitions
├── .env             # Environment variables
├── go.mod           # Go module definition
└── README.md        # You are here!
```

## 📄 License

This project is licensed under the [MIT License](LICENSE).
