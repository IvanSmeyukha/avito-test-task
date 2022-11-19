USE avito_db;

CREATE TABLE users (
  id INT PRIMARY KEY NOT NULL
);

CREATE TABLE accounts (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  value DECIMAL(8,2) NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE reserve (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  value DECIMAL(8,2) NOT NULL,
  user_id INT NOT NULL,
  deleted BOOLEAN NOT NULL DEFAULT '0',
  FOREIGN KEY (user_id) REFERENCES users (id)
);
   
   
CREATE TABLE transactions (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  service_id INT DEFAULT '0',
  order_id INT DEFAULT '0',
  reserve_id INT,
  value DECIMAL(8,2) NOT NULL,
  date DATETIME NOT NULL,
  type VARCHAR(255) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (reserve_id) REFERENCES reserve (id)
);