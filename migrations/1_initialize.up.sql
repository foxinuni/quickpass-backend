CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    number VARCHAR(13) UNIQUE NOT NULL
);

CREATE TABLE sessions (
    session_id SERIAL PRIMARY KEY,
    jwt_token VARCHAR(255),
    phone_model VARCHAR(255),
    enabled BOOLEAN,
    user_id INTEGER REFERENCES users(user_id)
);

CREATE TABLE events (
    event_id SERIAL PRIMARY KEY,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    address VARCHAR(255),
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE states (
    state_id SERIAL PRIMARY KEY,
    state VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE accomodations (
    acc_id SERIAL PRIMARY KEY,
    is_house BOOLEAN,
    address VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE bookings (
    booking_id SERIAL PRIMARY KEY,
    entry_date TIMESTAMP,
    leaving_date TIMESTAMP,
    acc_id INTEGER REFERENCES accomodations(acc_id)
);

CREATE TABLE occasions (
    occasion_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    event_id INTEGER REFERENCES events(event_id),
    booking_id INTEGER REFERENCES bookings(booking_id),
    state_id INTEGER REFERENCES states(state_id)
);

CREATE TABLE logs (
    log_id SERIAL PRIMARY KEY,
    occasion_id INTEGER REFERENCES occasions(occasion_id),
    time TIMESTAMP,
    is_inside BOOLEAN
);
