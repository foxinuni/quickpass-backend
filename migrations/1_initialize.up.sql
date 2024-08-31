CREATE TABLE USERS (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    number INTEGER(10)
);

CREATE TABLE SESSIONS (
    session_id SERIAL PRIMARY KEY,
    jwt_token VARCHAR(255),
    phone_model VARCHAR(255),
    IMEI INTEGER(10),
    enabled INTEGER(1),
    user_id INTEGER(10) REFERENCES USERS(user_id)
);

CREATE TABLE EVENTS (
    event_id SERIAL PRIMARY KEY,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    address VARCHAR(255),
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE STATES (
    state_id SERIAL PRIMARY KEY,
    state VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE ACCOMODATIONS (
    acc_id SERIAL PRIMARY KEY,
    is_house INTEGER(1),
    address VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE BOOKINGS (
    booking_id SERIAL PRIMARY KEY,
    entry_date TIMESTAMP,
    leaving_date TIMESTAMP,
    acc_id INTEGER(10) REFERENCES ACCOMODATIONS(acc_id)
);

CREATE TABLE OCCASIONS (
    occasion_id SERIAL PRIMARY KEY,
    user_id INTEGER(10) REFERENCES USERS(user_id),
    event_id INTEGER(10) REFERENCES EVENTS(event_id),
    booking_id INTEGER(10) REFERENCES BOOKINGS(booking_id),
    state_id INTEGER(10) REFERENCES STATES(state_id)
);

CREATE TABLE LOGS (
    log_id SERIAL PRIMARY KEY,
    time TIMESTAMP,
    is_inside INTEGER(1),
    occasion_id INTEGER(10) REFERENCES OCCASIONS(occasion_id)
);
