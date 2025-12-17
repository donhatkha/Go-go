# Go-go Quick Reference

## Project Structure

```
Go-go/
├── .github/
│   └── workflows/
│       └── ci.yml              # GitHub Actions CI/CD workflow
├── frontend/                    # Frontend files
│   ├── index.html              # Main HTML page
│   ├── style.css               # Styles
│   └── app.js                  # JavaScript application
├── migrations/                  # Database migrations
│   └── 001_create_users_table.sql
├── main.go                      # Go application entry point
├── main_test.go                 # Unit tests
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Dockerfile                   # Docker container definition
├── docker-compose.yml           # Multi-container orchestration
├── Makefile                     # Build automation
├── start.sh                     # Quick start script
├── .env.example                 # Environment variables template
├── .gitignore                   # Git ignore rules
├── README.md                    # Main documentation
├── CONTRIBUTING.md              # Contribution guidelines
└── LICENSE                      # MIT License

```

## Quick Commands

### Local Development
```bash
# Install dependencies
make install

# Run the application
make run

# Build the application
make build

# Run tests
make test

# Format code
make format

# Lint code
make lint
```

### Docker
```bash
# Start all services
make docker-up
# or
docker-compose up -d

# View logs
make docker-logs

# Stop all services
make docker-down

# Restart services
make docker-restart
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/users` | Get all users |
| GET | `/api/users/{id}` | Get user by ID |
| POST | `/api/users` | Create user |
| PUT | `/api/users/{id}` | Update user |
| DELETE | `/api/users/{id}` | Delete user |

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | Database host |
| `DB_PORT` | 5432 | Database port |
| `DB_USER` | postgres | Database user |
| `DB_PASSWORD` | postgres | Database password |
| `DB_NAME` | gogo_db | Database name |
| `DB_SSLMODE` | disable | Database SSL mode |
| `SERVER_PORT` | 8080 | Server port |
| `SERVER_HOST` | 0.0.0.0 | Server host |
| `CORS_ALLOWED_ORIGIN` | * | CORS allowed origin |
| `ENV` | development | Environment |

## Technology Stack

### Backend
- **Language:** Go 1.21
- **Router:** Gorilla Mux
- **Database:** PostgreSQL 15
- **Database Driver:** lib/pq
- **Config:** godotenv

### Frontend
- **HTML5**
- **CSS3** (with modern features)
- **Vanilla JavaScript** (ES6+)

### DevOps
- **Docker** & **Docker Compose**
- **GitHub Actions** (CI/CD)
- **pgAdmin** (Database Management)

## Security Features

✅ Parameterized SQL queries (SQL injection prevention)
✅ CORS configuration
✅ XSS prevention in frontend
✅ Environment-based configuration
✅ Secure GitHub Actions workflows
✅ CodeQL security scanning

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detection
go test -race ./...

# Run specific test
go test -run TestFunctionName
```

## Deployment

### Docker Deployment
```bash
# Quick start
./start.sh

# Manual start
docker-compose up -d
```

### Production Build
```bash
# Build optimized binary
go build -ldflags="-w -s" -o gogo-app main.go

# Run
./gogo-app
```

## Troubleshooting

### Database Connection Issues
1. Check PostgreSQL is running: `pg_isready`
2. Verify credentials in `.env`
3. Check database exists: `psql -U postgres -c "\l"`

### Port Already in Use
```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>
```

### Docker Issues
```bash
# Clean up containers
docker-compose down -v

# Rebuild images
docker-compose build --no-cache

# View logs
docker-compose logs -f app
```

## Useful Links

- **Repository:** https://github.com/donhatkha/Go-go
- **Go Docs:** https://golang.org/doc/
- **PostgreSQL Docs:** https://www.postgresql.org/docs/
- **Gorilla Mux:** https://github.com/gorilla/mux
- **Docker Docs:** https://docs.docker.com/

## Support

- 📧 Email: donhatkha@example.com
- 🐛 Issues: https://github.com/donhatkha/Go-go/issues
- 💬 Discussions: https://github.com/donhatkha/Go-go/discussions

---

**Last Updated:** 2024-12-17
