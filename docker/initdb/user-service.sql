CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255), -- For future use if implementing authentication
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00, -- Assuming currency format
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO public.users (username, password_hash, balance, created_at, updated_at) VALUES('alice', '$2y$10$8mqYbPkQKq8oCDvpKgbtdOvRwcy8G5CO4/S0Nid0AzcQ84tP5gTzW', 200.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.users (username, password_hash, balance, created_at, updated_at) VALUES('bob', '$2y$10$2qW24rlmi9RPk.kS2fCmQO8IS1mr96.RH9bokm69SGSXOh6X57Xou', 400.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
