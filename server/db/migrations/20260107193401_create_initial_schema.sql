-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    reg_no VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    no_show_count INTEGER NOT NULL DEFAULT 0,
    password_hash VARCHAR(255),
    role VARCHAR(50) NOT NULL DEFAULT 'student',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stands (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    location TEXT,
    capacity INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bikes (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    display_id VARCHAR(255) UNIQUE NOT NULL,
    current_stand_id INTEGER REFERENCES stands(id) ON DELETE SET NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bike_id INTEGER NOT NULL REFERENCES bikes(id) ON DELETE CASCADE,
    pickup_stand_id INTEGER NOT NULL REFERENCES stands(id) ON DELETE RESTRICT,
    return_stand_id INTEGER REFERENCES stands(id) ON DELETE RESTRICT,
    status VARCHAR(50) NOT NULL DEFAULT 'booked', -- 'booked', 'picked_up', 'returned', 'cancelled'
    picked_up_at TIMESTAMP WITH TIME ZONE,
    returned_at TIMESTAMP WITH TIME ZONE,
    duration_minutes INTEGER,
    pickup_photo_url TEXT,
    return_photo_url TEXT,
    guard_pickup_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    guard_return_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE violations (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL, -- 'late_return', 'fast_return', 'no_show', 'damage'
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS violations;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS bikes;
DROP TABLE IF EXISTS stands;
DROP TABLE IF EXISTS users;