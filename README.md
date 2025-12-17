# Go-go

A full-stack Go application template with PostgreSQL database, featuring a RESTful API backend and a modern web frontend.

## 🚀 Features

- **Backend (Go)**
  - RESTful API with CRUD operations
  - PostgreSQL database integration
  - Environment-based configuration
  - Health check endpoint
  - CORS support
  - Clean architecture

- **Frontend**
  - Modern, responsive UI
  - Vanilla JavaScript (no framework dependencies)
  - Real-time user management
  - Clean and intuitive design
  - Mobile-friendly

- **Database**
  - PostgreSQL 15
  - Automated migrations
  - Sample data included

- **DevOps**
  - Docker support
  - Docker Compose orchestration
  - pgAdmin for database management
  - Easy deployment

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose (optional, for containerized deployment)

## 🛠️ Installation

### Option 1: Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/donhatkha/Go-go.git
   cd Go-go
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Install PostgreSQL** (if not already installed)
   - macOS: `brew install postgresql`
   - Ubuntu: `sudo apt-get install postgresql`
   - Windows: Download from [postgresql.org](https://www.postgresql.org/download/)

4. **Create the database**
   ```bash
   psql -U postgres
   CREATE DATABASE gogo_db;
   \q
   ```

5. **Install Go dependencies**
   ```bash
   go mod download
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

7. **Access the application**
   - Frontend: http://localhost:8080
   - API: http://localhost:8080/api
   - Health Check: http://localhost:8080/api/health

### Option 2: Docker Deployment (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/donhatkha/Go-go.git
   cd Go-go
   ```

2. **Start the application**
   ```bash
   docker-compose up -d
   ```

3. **Access the services**
   - Application: http://localhost:8080
   - pgAdmin: http://localhost:5050 (admin@admin.com / admin)

4. **Stop the application**
   ```bash
   docker-compose down
   ```

## 📚 API Endpoints

### Health Check
- `GET /api/health` - Check server status

### User Management
- `GET /api/users` - Get all users
- `GET /api/users/{id}` - Get a specific user
- `POST /api/users` - Create a new user
- `PUT /api/users/{id}` - Update a user
- `DELETE /api/users/{id}` - Delete a user

### Request/Response Examples

**Create User**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

**Response**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

**Get All Users**
```bash
curl http://localhost:8080/api/users
```

**Update User**
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

**Delete User**
```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## 🏗️ Project Structure

```
Go-go/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── .env.example            # Environment variables template
├── Dockerfile              # Docker configuration
├── docker-compose.yml      # Docker Compose orchestration
├── frontend/               # Frontend files
│   ├── index.html         # Main HTML file
│   ├── style.css          # Styles
│   └── app.js             # JavaScript application
├── migrations/             # Database migrations
│   └── 001_create_users_table.sql
└── README.md              # This file
```

## 🔧 Configuration

Edit the `.env` file to configure the application:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gogo_db
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Environment
ENV=development
```

## 🐳 Docker Configuration

The project includes Docker support for easy deployment:

- **Dockerfile**: Multi-stage build for optimized image size
- **docker-compose.yml**: Orchestrates backend, database, and pgAdmin

Services included:
- `app`: Go application (port 8080)
- `db`: PostgreSQL database (port 5432)
- `pgadmin`: Database management UI (port 5050)

## 🧪 Testing

Run the application and test the endpoints:

```bash
# Start the server
go run main.go

# In another terminal, test the API
curl http://localhost:8080/api/health
curl http://localhost:8080/api/users
```

## 📦 Building for Production

**Build binary**
```bash
go build -o gogo-app main.go
./gogo-app
```

**Build Docker image**
```bash
docker build -t gogo-app .
docker run -p 8080:8080 gogo-app
```

## 🔒 Security Considerations

- Update database credentials in production
- Use HTTPS in production
- Implement authentication/authorization as needed
- Add rate limiting for API endpoints
- Validate and sanitize all user inputs
- Use prepared statements (already implemented)

## 🛣️ Roadmap

- [ ] Add user authentication (JWT)
- [ ] Implement role-based access control
- [ ] Add pagination for user lists
- [ ] Implement search and filtering
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add API documentation (Swagger)
- [ ] Add logging middleware
- [ ] Implement caching (Redis)
- [ ] Add CI/CD pipeline

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

Created by [donhatkha](https://github.com/donhatkha)

## 🙏 Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router
- [pq](https://github.com/lib/pq) - PostgreSQL driver
- [godotenv](https://github.com/joho/godotenv) - Environment configuration
- [PostgreSQL](https://www.postgresql.org/) - Database

## 📧 Support

For support, email donhatkha@example.com or open an issue on GitHub.

---

Made with ❤️ using Go and PostgreSQL