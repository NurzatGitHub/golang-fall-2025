CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    price INTEGER NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO categories (name) VALUES 
('phones'),
('laptops'),
('tablets');

INSERT INTO products (name, category_id, price) VALUES
('iPhone 14', 1, 400000),
('Samsung Galaxy', 1, 300000),
('MacBook Pro', 2, 600000),
('Dell XPS', 2, 500000),
('iPad Air', 3, 350000),
('Samsung Tablet', 3, 250000);