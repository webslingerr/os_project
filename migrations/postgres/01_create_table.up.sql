CREATE TYPE user_type AS ENUM ('Client', 'Owner', 'Sys_admin');

CREATE TABLE users (
    "id" UUID PRIMARY KEY NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "fullname" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "address" VARCHAR(500) NOT NULL,
    "type" user_type NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE real_estate_type AS ENUM ('Land', 'House', 'Apartment');

CREATE TYPE service_type AS ENUM ('Rent', 'Sell');

CREATE TABLE posts (
    "id" UUID PRIMARY KEY NOT NULL,
    "real_estate_type" real_estate_type NOT NULL,
    "service_type" service_type NOT NULL,
    "user_id" UUID NOT NULL,
    "title" VARCHAR NOT NULL,
    "description" TEXT,
    "region" VARCHAR NOT NULL,
    "address" VARCHAR NOT NULL,
    "contact_details" VARCHAR NOT NULL,
    "area" INTEGER,
    "number_of_rooms" INTEGER,
    "floor_number" INTEGER,
    "price" NUMERIC(12, 2),
    "rent_price" NUMERIC(12, 2),
    "special_benefits" TEXT[] DEFAULT '{}',
    "images" TEXT[] DEFAULT '{}',
    "status" VARCHAR DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
            REFERENCES users(id)
            ON DELETE CASCADE
);