# API Documentation

## Table of Contents
- [Register](#register)
  - [Request](#register-request)
  - [Response](#register-response)
- [Login](#login)
  - [Request](#login-request)
  - [Response](#login-response)

---

## /register
**POST Method**

### Register Request
```json
{
  "email": "example@example.com",
  "fullname": "John Doe",
  "password": "password123",
  "address": "123 Main Street",
  "type": "Client"
}
```

### Register Response
```json
{
  "id": "abc123",
  "email": "example@example.com",
  "fullname": "John Doe",
  "password": "password123",
  "address": "123 Main Street",
  "type": "Client",
  "created_at": "2024-12-22T12:00:00Z",
  "updated_at": "2024-12-22T12:00:00Z"
}
```

---

## /login
**POST Method**

### Login Request
```json
{
  "email": "example@example.com",
  "password": "password123"
}
```

### Login Response
```json
{
  "id": "abc123",
  "email": "example@example.com",
  "fullname": "John Doe",
  "password": "password123",
  "address": "123 Main Street",
  "type": "Client",
  "created_at": "2024-12-22T12:00:00Z",
  "updated_at": "2024-12-22T12:00:00Z"
}
