# API Documentation

## Overview
This document provides details about the available API endpoints and their functionality.

---

## Endpoints

### 1. **/register**

#### Description
This endpoint allows users to register by providing necessary details such as email, fullname, password, address, and type.

#### Request
**Method:** `POST`

**Fields:**
- **email** (string): The user's email address.
- **fullname** (string): The user's full name.
- **password** (string): The user's password.
- **address** (string): The user's address.
- **type** (string): The type of user. Possible values are:
  - `Client`
  - `Owner`
  - `Sys_admin`

#### Response
**Content-Type:** `application/json`

**Response Body:**
```json
{
  "id": "string",
  "email": "string",
  "fullname": "string",
  "password": "string",
  "address": "string",
  "type": "UserType",
  "created_at": "string",
  "updated_at": "string"
}
```

**Fields in Response:**
- **id** (string): A unique identifier for the user.
- **email** (string): The registered email address.
- **fullname** (string): The registered full name.
- **password** (string): The hashed password.
- **address** (string): The registered address.
- **type** (UserType): The user type (Client, Owner, Sys_admin).
- **created_at** (string): The timestamp when the user was created.
- **updated_at** (string): The timestamp when the user information was last updated.

#### Example
**Request:**
```json
{
  "email": "example@example.com",
  "fullname": "John Doe",
  "password": "password123",
  "address": "123 Main Street",
  "type": "Client"
}
```

**Response:**
```json
{
  "id": "abc123",
  "email": "example@example.com",
  "fullname": "John Doe",
  "password": "hashed_password",
  "address": "123 Main Street",
  "type": "Client",
  "created_at": "2024-12-22T12:00:00Z",
  "updated_at": "2024-12-22T12:00:00Z"
}
