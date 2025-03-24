# freshers-bootcamp-last

# Technologies Used
- Backend: Golang (Gin framework)
- Database: MySQL
- ORM: GORM
- Authentication: JWT (JSON Web Tokens)
- Concurrency Handling: Mutex

# Setup Instructions
## Prerequisites
- Go: Install Go
- MySQL: Install MySQL
- Git: Install Git

### **Environment Variables**
Create a `.env` file in the root directory with the following variables:

```env
DB_USER=your_mysql_username
DB_PASSWORD=your_mysql_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=retailer_db
JWT_SECRET=your-256-bit-secret
```

## Database Setup
- Start your MySQL server. 
- Create a database named retailer_db:

```sql
CREATE DATABASE retailer_db;
```
- The application will automatically create the required tables using GORM's AutoMigrate feature.


## Running the Application

1. Clone the repository:
```
git clone https://github.com/your-username/retailer-service.git
cd retailer-service
```
2. Install dependencies:

```
go mod tidy
```

3. Run the application:
```
go run main.go
```
The application will start on http://localhost:8080.

# API Endpoints

This document lists all the available API endpoints for the Retailer Service. Each endpoint is described with its method, URL, and functionality.

---

## **Authentication**

| **Method** | **Endpoint** | **Description**               |
|------------|--------------|-------------------------------|
| `POST`     | `/register`  | Register a new user           |
| `POST`     | `/login`     | Login and get a JWT token     |

---

## **Products**

| **Method** | **Endpoint**         | **Description**                          |
|------------|----------------------|------------------------------------------|
| `POST`     | `/products`          | Add a new product (Admin only)           |
| `GET`      | `/products`          | Get all available products               |
| `PUT`      | `/products/:id`      | Update a product (Admin only)            |
| `DELETE`   | `/products/:id`      | Delete a product (Admin only)            |

---

## **Orders**

| **Method** | **Endpoint**         | **Description**                          |
|------------|----------------------|------------------------------------------|
| `POST`     | `/orders`            | Place a new order                        |
| `GET`      | `/orders`            | Get order history for the logged-in user |
| `GET`      | `/orders/:id`        | Get details of a specific order          |

---

## **Admin**

| **Method** | **Endpoint**         | **Description**                          |
|------------|----------------------|------------------------------------------|
| `POST`     | `/admin/register`    | Register a new admin user (Admin only)   |

---

# testing screenshots (request response):

![Screenshot 2025-03-24 at 10.52.58 AM.png](..%2FScreenshot%202025-03-24%20at%2010.52.58%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 10.56.22 AM.png](..%2FScreenshot%202025-03-24%20at%2010.56.22%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 11.43.12 AM.png](..%2FScreenshot%202025-03-24%20at%2011.43.12%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 11.43.42 AM.png](..%2FScreenshot%202025-03-24%20at%2011.43.42%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 11.44.21 AM.png](..%2FScreenshot%202025-03-24%20at%2011.44.21%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 11.49.39 AM.png](..%2FScreenshot%202025-03-24%20at%2011.49.39%E2%80%AFAM.png)
![Screenshot 2025-03-24 at 11.50.17 AM.png](..%2FScreenshot%202025-03-24%20at%2011.50.17%E2%80%AFAM.png)