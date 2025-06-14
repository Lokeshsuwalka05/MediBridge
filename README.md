# MediBridge - Hospital Management System

A Golang-based hospital management system with role-based access control for doctors and receptionists.

## Features

- JWT-based authentication
- Role-based access control (Doctor and Receptionist roles)
- Patient management system
- RESTful API endpoints
- PostgreSQL database with GORM ORM
- Unit tests

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/medibridge.git
cd medibridge
```

2. Create a `.env` file in the root directory with the following content:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=medibridge
DB_PORT=5432
JWT_SECRET=your-secret-key-here
SERVER_PORT=8080
```

3. Create the PostgreSQL database:
```bash
createdb medibridge
```

4. Install dependencies:
```bash
go mod download
```

5. Run the application:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` (or the port specified in your .env file).

## API Endpoints

### Authentication
- `POST /login` - Login and get JWT token

### Receptionist Endpoints
- `POST /patients` - Create new patient
- `GET /patients` - List all patients
- `PUT /patients/:id` - Update patient
- `DELETE /patients/:id` - Delete patient

### Doctor Endpoints
- `GET /patients` - View patient list
- `PATCH /patients/:id` - Update patient record (diagnosis/notes)

## Default Users

The system comes with two default users:

1. Doctor:
   - Email: doctor@medibridge.com
   - Password: doctor@#123

2. Receptionist:
   - Email: receptionist@medibridge.com
   - Password: reception@#123

## Testing

Run the test suite:
```bash
go test ./...
```

## Live Demo
Frontend:https://medi-bridge.netlify.app/
Backend:https://medibridge-backend-znn8.onrender.com/

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── config/
│   └── database.go       # Database configuration
├── controllers/
│   ├── auth.go          # Authentication controller
│   ├── auth_test.go     # Auth tests
│   └── patient.go       # Patient management controller
├── middleware/
│   └── auth.go          # JWT and role middleware
├── models/
│   ├── user.go          # User model
│   └── patient.go       # Patient model
├── routes/
│   └── routes.go        # Route definitions
├── utils/
│   └── jwt.go           # JWT utilities
├── .env.example         # Example environment variables
├── go.mod              # Go module file
└── README.md           # This file
```

## Security Considerations

- All passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- Role-based access control for all endpoints
- Environment variables for sensitive data

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Request Flow:
Client Request → Routes → Middleware → Controllers → Models → Database
Response Flow:
Database → Models → Controllers → Response to Client

## License

This project is licensed under the MIT License - see the LICENSE file for details.
