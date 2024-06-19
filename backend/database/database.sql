CREATE DATABASE IF NOT EXISTS easychat_db DEFAULT CHARACTER SET utf8;
USE easychat_db;

CREATE TABLE IF NOT EXISTS list (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	message TEXT NOT NULL,
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);