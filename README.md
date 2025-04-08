# Full Cycle Payment Gateway

A robust payment gateway implementation built with Go, featuring account management, transaction processing, and webhook support.

## Project Structure

```
.
├── cmd/
│   └── app/              # Application entry point
├── internal/
│   ├── domain/          # Domain entities and business rules
│   ├── repository/      # Database access layer
│   ├── service/         # Business logic layer
│   ├── web/            # HTTP handlers and server setup
│   └── dto/            # Data Transfer Objects
├── migrations/          # Database migrations
└── test.http           # API test requests
```

## Planned Features

- Account management with API key authentication
- Secure payment processing
- Transaction history tracking
- PostgreSQL database integration

## Prerequisites
- Go 1.16 or higher
- PostgreSQL 12 or higher

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
HTTP_PORT=8080
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/AlecDr/Full-Cycle-Payment-Gateway.git
cd Full-Cycle-Payment-Gateway
```

2. Install dependencies:
```bash
go mod download
```

3. Run database migrations:
```bash
migrate -path migrations -database "your_db_connection_string" up
```

4. Start the application:
```bash
go run cmd/app/main.go
```

## API Endpoints

### Accounts

- `POST /accounts` - Create a new account
- `GET /accounts/{id}` - Get account details
- `GET /accounts` - List all accounts

## Development

### Running Tests

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.