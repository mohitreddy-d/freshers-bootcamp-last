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

<img width="1512" alt="Screenshot 2025-03-24 at 10 52 58 AM" src="https://github.com/user-attachments/assets/5f4879c0-ae3e-4425-bf73-cf1091a35e55" />
<img width="1512" alt="Screenshot 2025-03-24 at 10 56 22 AM" src="https://github.com/user-attachments/assets/bfca8d19-f6d1-40b8-8c48-4f91435c6df7" />
<img width="1512" alt="Screenshot 2025-03-24 at 11 43 12 AM" src="https://github.com/user-attachments/assets/910bee31-0a02-4568-9830-a253ab1c961c" />
<img width="1512" alt="Screenshot 2025-03-24 at 11 43 42 AM" src="https://github.com/user-attachments/assets/e3b1fb1b-2d10-4bfc-8c2c-be091f5e55da" />
<img width="1512" alt="Screenshot 2025-03-24 at 11 44 21 AM" src="https://github.com/user-attachments/assets/c275824d-6193-4ba0-bb74-d3189d506b49" />
<img width="1512" alt="Screenshot 2025-03-24 at 11 49 39 AM" src="https://github.com/user-attachments/assets/9d6fb1a1-bdc7-4d0b-9dfb-e03c5a1e21ea" />
<img width="1512" alt="Screenshot 2025-03-24 at 11 50 17 AM" src="https://github.com/user-attachments/assets/b9c3244a-b329-4523-bdf8-7c1c3bdd9b8d" />
