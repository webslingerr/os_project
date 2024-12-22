# API Documentation

## Table of Contents
- [Register](#register)
  - [Request](#register-request)
  - [Response](#register-response)
- [Login](#login)
  - [Request](#login-request)
  - [Response](#login-response)
- [Create Post](#create-post)
  - [Request](#create-post-request)
  - [Response](#create-post-response)

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
```

---

## /post
**POST Method**

### Create Post Request
```json
{
  "real_estate_type": "House",
  "service_type": "Rent",
  "user_id": "086520ea-b9ab-4200-b38a-40e60c5c0dfb",
  "title": "Spacious 3 Bedroom House for Rent",
  "description": "A beautiful and spacious house with a large garden and modern amenities.",
  "region": "Tashkent",
  "address": "5678 Oak Avenue, Metropolis, NY",
  "contact_details": "contact@realestate.com",
  "area": 2500,
  "number_of_rooms": 3,
  "floor_number": 1,
  "rent_price": 2000000.00,
  "special_benefits": ["Swimming Pool", "Gym Access"]
}
```

### Create Post Response
```json
{
  "id": "8ee2ff28-0c24-40cf-befa-6543aca2018b",
  "real_estate_type": "House",
  "service_type": "Rent",
  "user_id": "086520ea-b9ab-4200-b38a-40e60c5c0dfb",
  "title": "Spacious 3 Bedroom House for Rent",
  "description": "A beautiful and spacious house with a large garden and modern amenities.",
  "region": "Tashkent",
  "address": "5678 Oak Avenue, Metropolis, NY",
  "contact_details": "contact@realestate.com",
  "area": 2500,
  "number_of_rooms": 3,
  "floor_number": 1,
  "price": 0,
  "rent_price": 2000000.00,
  "special_benefits": ["Swimming Pool", "Gym Access"],
  "images": [],
  "status": "pending",
  "created_at": "2024-12-22T12:00:00Z",
  "updated_at": "2024-12-22T12:00:00Z"
}
```

### Notes
- `real_estate_type` can be one of the following:
  - `Land`
  - `House`
  - `Apartment`
- `service_type` can be one of the following:
  - `Rent`
  - `Sell`
