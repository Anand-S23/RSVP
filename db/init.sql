CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255),
);

CREATE TABLE people (
    id UUID PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(255),
    answered BOOLEAN DEFAULT FALSE,
    attending BOOLEAN DEFAULT FALSE,
    date_arriving DATE,
    date_departure DATE,
    comment VARCHAR(500),
);

