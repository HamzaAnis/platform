CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL, -- This acts as a logical reference to a User in the User Service.
    type VARCHAR(50) NOT NULL, -- e.g., 'credit', 'debit', 'transfer'
    amount DECIMAL(10, 2) NOT NULL,
    reference_user_id INT, -- For transfers, this acts as a logical reference to another User.
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO public.transactions (user_id, "type", amount, reference_user_id, created_at) VALUES(1, 'C', 400.0, 0, CURRENT_TIMESTAMP);
INSERT INTO public.transactions (user_id, "type", amount, reference_user_id, created_at) VALUES(2, 'C', 250.0, 0, CURRENT_TIMESTAMP);
INSERT INTO public.transactions (user_id, "type", amount, reference_user_id, created_at) VALUES(1, 'T', 150.0, 2, CURRENT_TIMESTAMP);
INSERT INTO public.transactions (user_id, "type", amount, reference_user_id, created_at) VALUES(1, 'D', 50.0, 0, CURRENT_TIMESTAMP);
