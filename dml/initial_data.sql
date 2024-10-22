--poblate event table

INSERT INTO events(start_date, end_date, address, name) 
VALUES('2024-09-01 10:00:00', '2024-12-01 10:00:00', 'calle 7a', 'comic con');

INSERT INTO events(start_date, end_date, address, name) 
VALUES('2024-08-01 10:00:00', '2024-08-29 10:00:00', 'calle 8a', 'Feria del carro');

INSERT INTO events(start_date, end_date, address, name) 
VALUES('2024-09-01 10:00:00', '2024-11-01 10:00:00', 'calle 6a', 'Filbo');

--poblate accomodations
INSERT INTO accomodations(is_house, address) VALUES(true, 'carrera 8');

--poblate bookings
INSERT INTO bookings (entry_date, leaving_date, acc_id) 
VALUES (
    '2024-05-01 10:00:00', 
    '2025-11-01 10:00:00', 
    (SELECT acc_id FROM accomodations WHERE address = 'carrera 8' LIMIT 1)
);

--poblate users
INSERT INTO users(email, number) VALUES ('juan@gmail.com', '3008522437');

--poblate occasions
INSERT INTO occasions(user_id, booking_id, state_id) VALUES(
    (select user_id from users where email = 'juan@gmail.com' LIMIT 1),
    (select booking_id from bookings LIMIT 1),
    (select state_id from states where name = 'REGISTRADO' LIMIT 1)
);
INSERT INTO occasions(user_id, event_id, state_id) VALUES(
    (select user_id from users where email = 'juan@gmail.com' LIMIT 1),
    (select event_id from events where name = 'comic con' LIMIT 1),
    (select state_id from states where name = 'REGISTRADO' LIMIT 1)
);
INSERT INTO occasions(user_id, event_id, state_id) VALUES(
    (select user_id from users where email = 'juan@gmail.com' LIMIT 1),
    (select event_id from events where name = 'Feria del carro' LIMIT 1),
    (select state_id from states where name = 'REGISTRADO' LIMIT 1)
);
INSERT INTO occasions(user_id, event_id, state_id) VALUES(
    (select user_id from users where email = 'juan@gmail.com' LIMIT 1),
    (select event_id from events where name = 'Filbo' LIMIT 1),
    (select state_id from states where name = 'REGISTRADO' LIMIT 1)
);