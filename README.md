# API Documentation

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

---

## /post/{id}
**GET Method**

### Get Post by ID Response
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

---

## /post/get-list
**POST Method**

### Get List of Posts Request
```json
{
  "user_id": "optional",
  "status": "optional",
  "real_estate_type": "optional",
  "service_type": "optional",
  "search": "optional",
  "limit": "optional"
}
```

### Get List of Posts Response
```json
{
    "count": 1,
    "posts": [
      {
        "id": "c25d9562-9ffb-4f09-a628-abdcfba9323f",
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
        "rent_price": 2000000,
        "special_benefits": ["Swimming Pool", "Gym Access"],
        "images": null,
        "status": "accepted",
        "created_at": "",
        "updated_at": ""
      }
    ]
  }
```

### Notes
- For `Owner`, `user_id` must be included in the request.
- For `Sys_admin`, pass status to request status: "pending"
- Filters:
  - `real_estate_type`
  - `service_type`
  - `title` pass to search field
- If `limit` is not provided, the default value is 10.

---

## /post/status/{id}
**PUT Method**

### Update Post Status Request
```json
{
  "status": "accepted"
}
```

- Empty Response

### Notes
- The `status` field is required.
- If the post is accepted, the `status` should be set to `accepted`.