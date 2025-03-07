# API Register & Login

This API supports user registration, login, and profile management with JWT-based authentication.

## Features

- **User Registration**: Users can register with a username, password, full name, and avatar.
- **User Login**: Users can log in using their credentials and receive a JWT token for authentication.
- **User Profile**: Authenticated users can view their profile information.
- **Read All Users**: Authenticated users can view a list of all registered users.

## Technologies Used

- **Gin**: A high-performance HTTP web framework for Go.
- **GORM**: An ORM library for Go, used for database interactions.
- **JWT (JSON Web Tokens)**: Used for secure user authentication.
- **MySQL**: The database used to store user information.
- **Bcrypt**: Used for securely hashing user passwords.

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/dekbahnerd/Api-Register-Login.git
cd Api-Register-Login
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up the database

- Create a MySQL database.
- Update the `.env` file with your database credentials:

```env
DB_USERNAME=your_db_username
DB_PASSWORD=your_db_password
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_NAME=your_db_name
JWT_SECRET_KEY=your_jwt_secret_key
```

### 4. Run the program

```bash
go run main.go
```

## API Endpoints

#### **POST /register**: Register a new user.

- **Request Body:**

```json
{
  "username": "testuser",
  "password": "testpassword",
  "full_name": "Test User",
  "avatar": "http://example.com/avatar.jpg"
}
```

- **Response:**

```json
{
  "status": "ok",
  "message": "User created successfully",
  "user_id": 1
}
```

#### **POST /login**: Log in an existing user.

- **Request Body:**

```json
{
  "username": "testuser",
  "password": "testpassword"
}
```

- **Response:**

```json
{
  "status": "ok",
  "message": "Login successful",
  "token": "your_jwt_token"
}
```

#### **GET /users/readall**: Get a list of all users (requires authentication).

- **Response:**

```json
{
  "message": "Read all users",
  "users": [
    {
      "id": 1,
      "username": "testuser",
      "full_name": "Test User",
      "avatar": "http://example.com/avatar.jpg"
    }
  ]
}
```

#### **GET /users/profile**: Get the profile of the authenticated user.

- **Response:**

```json
{
  "message": "Read successfully",
  "user": {
    "id": 1,
    "username": "testuser",
    "full_name": "Test User",
    "avatar": "http://example.com/avatar.jpg"
  }
}
```
