CREATE TABLE blog
(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255)
);
CREATE UNIQUE INDEX blog_id_uindex ON blog (id);