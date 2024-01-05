CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE owners (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    country_code TEXT,
    phone_number TEXT
);

CREATE TABLE fields (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID NOT NULL REFERENCES owners(id),
    street TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    status TEXT NOT NULL,
    open_time TIMESTAMP NOT NULL,
    close_time TIMESTAMP NOT NULL
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    field_id UUID NOT NULL REFERENCES fields(id),
    booker_id UUID NOT NULL,  -- Assuming booker_id references another table that hasn't been converted to UUID
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status TEXT NOT NULL,
    payment_status TEXT,
    payment_id TEXT
);
