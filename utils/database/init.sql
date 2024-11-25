create database bioskop_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100) NOT NULL,
    seats INTEGER NOT NULL
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    cinema_id INT REFERENCES cinemas(id),
    payment_id VARCHAR(100),
    payment_detail JSONB,
    date DATE NOT NULL,
    time TIME NOT NULL
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE seats (
    id SERIAL PRIMARY KEY,
    cinema_id INT REFERENCES cinemas(id),
    seat_id VARCHAR(10) UNIQUE NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'available'
);
