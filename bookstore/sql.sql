-- 创建sessions表
CREATE TABLE sessions(
session_id VARCHAR(100) PRIMARY KEY,
username VARCHAR(100) NOT NULL,
user_id INT NOT NULL,
FOREIGN KEY(user_id) REFERENCES users(id)
)

--创建购物车表
CREATE TABLE carts(
id VARCHAR(100) PRIMARY KEY,
total_count INT NOT NULL,
total_amount DOUBLE(11,2) NOT NULL,
user_id INT NOT NULL,
FOREIGN KEY(user_id) REFERENCES users(id)
)

-- 创建购物项表
CREATE TABLE cart_itmes(
id INT PRIMARY KEY AUTO_INCREMENT,
COUNT INT NOT NULL,
amount DOUBLE(11,2) NOT NULL,
book_id INT NOT NULL,
cart_id VARCHAR(100) NOT NULL,
FOREIGN KEY(book_id) REFERENCES books(id),
FOREIGN KEY(cart_id) REFERENCES carts(id)
)