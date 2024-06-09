# Snaptalky

Snaptalky is an application built with Go using the Gin framework. It allows users to get the best response for their messages by uploading either a message or a screenshot.

## Features

- User management: Create, update, and retrieve user information.
- Response processing: Upload a message or a screenshot and get a suitable response.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/snaptalky-app.git
   cd snaptalky-app
   ```

2. **Set up the environment variables:**
   Create a `.env` file in the root directory and add your PostgreSQL credentials:
   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_HOST=localhost
   DB_PORT=5432
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## Endpoints

### User Routes

- **GET /user/:id**: Retrieve user information by ID.
- **PUT /user**: Update user information.

### Scan Routes

- **POST /scan**: Process a text or image input and get a suitable response.

## License

This project is licensed under the MIT License.
