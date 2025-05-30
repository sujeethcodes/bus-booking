CREATE TABLE user(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(100),
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100),
    status VARCHAR(100),
    token VARCHAR(100),
    address JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE user
ADD COLUMN phone_number VARCHAR(100);

ALTER TABLE user 
MODIFY COLUMN token VARCHAR(500);