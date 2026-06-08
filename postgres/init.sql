CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'USER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert a default ADMIN user (password is secret123)
INSERT INTO users (name, email, password, role)
VALUES ('Super Admin', 'admin@example.com', '$2a$10$p07uGPwQRtmwjqNIbmYUs.Fqy5TnZOANZAOyNWNun.HVHeyloy9U.', 'ADMIN')
ON CONFLICT (email) DO NOTHING;
