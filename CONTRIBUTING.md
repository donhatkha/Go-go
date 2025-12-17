# Contributing to Go-go

Thank you for your interest in contributing to Go-go! This document provides guidelines and instructions for contributing.

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose (for containerized development)
- Git

### Setting Up Development Environment

1. **Fork and Clone**
   ```bash
   git clone https://github.com/YOUR_USERNAME/Go-go.git
   cd Go-go
   ```

2. **Install Dependencies**
   ```bash
   make install
   # or
   go mod download
   ```

3. **Set Up Environment**
   ```bash
   cp .env.example .env
   # Edit .env with your local configuration
   ```

4. **Run the Application**
   ```bash
   # Using Docker (recommended)
   make docker-up
   
   # Or locally
   make run
   ```

## 💻 Development Workflow

### Code Style
- Follow standard Go conventions
- Run `go fmt` before committing
- Run `go vet` to catch common issues
- Use meaningful variable and function names

### Making Changes

1. **Create a Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Your Changes**
   - Write clean, readable code
   - Add comments for complex logic
   - Follow existing code structure

3. **Write Tests**
   ```bash
   # Run tests
   make test
   # or
   go test -v ./...
   ```

4. **Format Code**
   ```bash
   make format
   # or
   go fmt ./...
   ```

5. **Check for Issues**
   ```bash
   make lint
   # or
   go vet ./...
   ```

### Commit Guidelines

- Use clear, descriptive commit messages
- Follow the format: `type: description`
- Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

Examples:
```
feat: add user authentication endpoint
fix: resolve database connection timeout
docs: update API documentation
```

### Pull Request Process

1. **Update Documentation**
   - Update README.md if needed
   - Add comments to complex code
   - Update API documentation

2. **Test Your Changes**
   ```bash
   # Run all tests
   go test -v ./...
   
   # Test locally
   go run main.go
   
   # Test with Docker
   docker-compose up
   ```

3. **Create Pull Request**
   - Provide a clear description of changes
   - Reference any related issues
   - Ensure all tests pass
   - Wait for review

## 🧪 Testing

### Writing Tests
- Place tests in `*_test.go` files
- Use table-driven tests when appropriate
- Test both success and error cases
- Aim for high code coverage

Example:
```go
func TestYourFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"test case 1", "input1", "output1"},
        {"test case 2", "input2", "output2"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := YourFunction(tt.input)
            if result != tt.expected {
                t.Errorf("got %s, want %s", result, tt.expected)
            }
        })
    }
}
```

### Running Tests
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestFunctionName
```

## 📝 Documentation

- Update README.md for user-facing changes
- Add comments to exported functions
- Document API endpoints
- Include examples where helpful

## 🐛 Reporting Bugs

### Before Reporting
- Check if the issue already exists
- Verify it's reproducible
- Test on the latest version

### Bug Report Template
```markdown
**Description**
Clear description of the bug

**Steps to Reproduce**
1. Step 1
2. Step 2
3. Step 3

**Expected Behavior**
What should happen

**Actual Behavior**
What actually happens

**Environment**
- OS: 
- Go Version: 
- PostgreSQL Version: 

**Additional Context**
Any other relevant information
```

## 💡 Suggesting Features

### Feature Request Template
```markdown
**Feature Description**
Clear description of the feature

**Use Case**
Why is this feature needed?

**Proposed Solution**
How should it work?

**Alternatives Considered**
Other approaches you've thought about

**Additional Context**
Any other relevant information
```

## 🔍 Code Review Process

### What We Look For
- Code quality and readability
- Test coverage
- Documentation
- Performance implications
- Security considerations
- Backward compatibility

### Review Timeline
- Initial review: 1-3 days
- Follow-up reviews: 1-2 days
- Approval: When all requirements are met

## 🎯 Areas for Contribution

### Good First Issues
- Documentation improvements
- Test coverage
- Bug fixes
- UI enhancements

### Advanced Contributions
- New features
- Performance optimizations
- Architecture improvements
- Security enhancements

## 📚 Resources

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Gorilla Mux Documentation](https://github.com/gorilla/mux)
- [Docker Documentation](https://docs.docker.com/)

## 🤝 Code of Conduct

### Our Standards
- Be respectful and inclusive
- Accept constructive criticism
- Focus on what's best for the community
- Show empathy towards others

### Unacceptable Behavior
- Harassment or discrimination
- Trolling or insulting comments
- Personal or political attacks
- Publishing others' private information

## ❓ Questions?

- Open an issue for discussion
- Reach out to maintainers
- Check existing documentation

## 🙏 Thank You!

Your contributions make Go-go better for everyone. We appreciate your time and effort!

---

Happy Contributing! 🚀
