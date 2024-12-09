# Golang Todo Server  

A simple Todo API server built using Go.  

---

## **Table of Contents**  
- [Getting Started](#getting-started)  
- [Running Locally](#running-locally)  
- [Running with Docker](#running-with-docker)  
- [Environment Variables Example](#environment-variables-example)  
- [Troubleshooting](#troubleshooting)  
- [Contributing](#contributing)  

---

## **Getting Started**  

### **Prerequisites**  
- Go v1.23.3
- MySQL Database  
- Docker & Docker Compose (if using Docker)  

---

## **Running Locally**

1. **Setup Environment Variables**  
   - Create a `.env` file in the project root directory.  
   - Define the necessary environment variables (see [Environment Variables Example](#environment-variables-example)).  

2. **Setup MySQL Database**  
   - Ensure a MySQL database is running locally.  
   - Use the database name set in the `.env` file.  

3. **Install Dependencies**  
   ```bash
   go mod download
   ```

4. **Run the Server with Hot Reload**  
   ```bash
   air
   ```

---

## **Running with Docker**  

1. **Setup Environment Variables**  
   - Create a `.env` file in the project root directory.  
   - Define the necessary environment variables (see [Environment Variables Example](#environment-variables-example)).  

2. **Run Docker Compose**  
   ```bash
   docker compose up
   ```

### **Stopping the Application**  
```bash
docker compose down
```

---

## **Environment Variables Example (.env)**  
```env
APP_NAME=golang_todo_server

SERVER_PORT=8080

DATABASE_USER=root
DATABASE_PASS=password
DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_NAME=todo_db

JWT_SECRET=your_jwt_secret_key
```

---

## **Troubleshooting**  

- Ensure Go, Docker, and MySQL are properly installed and configured.  
- Verify correct database credentials in the `.env` file.  
- Check Docker logs with:  
  ```bash
  docker compose logs -f
  ```
- If `air` is not found, install it using:  
  ```bash
  go install github.com/cosmtrek/air@latest
  ```

---

## **Contributing**  

Contributions are welcome! Feel free to submit issues or create pull requests.
