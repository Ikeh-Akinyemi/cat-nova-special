-- Create the Customers table
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Insert sample customer data
INSERT INTO customers (name, email, created_at)
VALUES
    ('John Doe', 'john@example.com', NOW()),
    ('Jane Smith', 'jane@example.com', NOW());
