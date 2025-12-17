#!/bin/bash

# Quick start script for Go-go application

echo "🚀 Go-go Quick Start"
echo "===================="
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker and try again."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose and try again."
    exit 1
fi

echo "✅ Docker and Docker Compose are installed"
echo ""

# Start the application
echo "🐳 Starting containers..."
docker-compose up -d

echo ""
echo "⏳ Waiting for services to be ready..."
sleep 10

echo ""
echo "✅ Application is running!"
echo ""
echo "🌐 Access points:"
echo "   - Application:  http://localhost:8080"
echo "   - API Health:   http://localhost:8080/api/health"
echo "   - API Users:    http://localhost:8080/api/users"
echo "   - pgAdmin:      http://localhost:5050"
echo ""
echo "📝 To view logs:    docker-compose logs -f"
echo "🛑 To stop:         docker-compose down"
echo ""
