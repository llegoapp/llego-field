CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    country_code TEXT,
    phone_number TEXT
);

CREATE TABLE fields (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES owners(id),
    street TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    status TEXT NOT NULL,
    open_time TIMESTAMP NOT NULL,
    close_time TIMESTAMP NOT NULL
);


CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    field_id INTEGER NOT NULL REFERENCES fields(id),
    booker_id INTEGER NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    payment_status TEXT,
    payment_id TEXT
);
