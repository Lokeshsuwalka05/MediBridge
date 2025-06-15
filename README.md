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
- `GET /auth/validate` - Validate JWT token and get user details

### Receptionist Endpoints
- `POST /receptionist/patients` - Create a new patient record. Requires `firstName`, `lastName`, `email`, `phone`, `dateOfBirth` (YYYY-MM-DD), `gender` (male/female/other), `address`, `emergencyContact`, `emergencyPhone`. Optional: `bloodGroup`, `allergies`.
- `GET /receptionist/patients` - Get paginated list of all patients. Supports `page`, `limit`, and `search` query parameters.
- `PUT /receptionist/patients/:id` - Update patient information. Allows partial updates for `firstName`, `lastName`, `email`, `phone`, `dateOfBirth`, `gender`, `address`, `emergencyContact`, `emergencyPhone`, `bloodGroup`, `allergies`.
- `DELETE /receptionist/patients/:id` - Delete a patient record.

### Doctor Endpoints
- `GET /doctor/patients` - View paginated list of all patients. Supports `page`, `limit`, and `search` query parameters.
- `PATCH /doctor/patients/:id` - Update patient medical record (diagnosis and notes). Only `diagnosis` and `notes` fields can be updated by doctors.

## API Documentation with Postman

This project includes a Postman collection and environment to help you easily test and interact with the API endpoints.

### Online API Documentation

You can also view and interact with the MediBridge API documentation online via Postman:

[https://documenter.getpostman.com/view/45825870/2sB2x6nsrU](https://documenter.getpostman.com/view/45825870/2sB2x6nsrU)

### 1. Import Postman Files

1.  Open Postman (ensure you have it installed).
2.  Click on the "Import" button in the sidebar (or File > Import).
3.  Select the following files from your `backend/` directory:
    *   `MediBridge.postman_collection.json`
    *   `MediBridge.postman_environment.json`

### 2. Configure Environment

1.  In Postman, select the "Environments" tab in the sidebar.
2.  Choose "MediBridge Environment" from the dropdown.
3.  Ensure the `base_url` variable is set correctly. For local development, it should be `http://localhost:8080`.
    *   The `token` variable will be automatically populated after a successful login.

### 3. Using the Collection

The collection is organized into folders:

*   **Auth**: Contains endpoints for user authentication (Login, Validate Token).
*   **Receptionist**: Contains endpoints for managing patient records (Create, Get, Update, Delete Patients).
*   **Doctor**: Contains endpoints for viewing and updating patient medical records (Get, Update Patient Medical Record).

**Workflow Example:**

1.  **Login**: Send a `POST` request to `/login` with doctor or receptionist credentials. The response will include a JWT. Postman's test script (if configured) will automatically store this token in your environment.
    *   Example Doctor Credentials: `{"email": "doctor@medibridge.com", "password": "doctor@#123"}`
2.  **Validate Token (Optional but Recommended)**: Use the `GET /auth/validate` endpoint to confirm your token is valid and retrieve your user details.
3.  **Access Protected Routes**: For subsequent requests to protected endpoints (e.g., `/receptionist/patients`, `/doctor/patients`), ensure the `Authorization` header is set to `Bearer {{token}}`. Postman should automatically handle this if the token was saved to the environment.

Each request in the collection includes descriptions, example request bodies, and details on query parameters or path variables where applicable.

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
