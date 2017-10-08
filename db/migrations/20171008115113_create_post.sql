CREATE TABLE post
(
  id VARCHAR(255) PRIMARY KEY,
  body TEXT,
  blog_id VARCHAR(255) NOT NULL,
  CONSTRAINT post_blog_id_fk FOREIGN KEY (blog_id) REFERENCES blog (id)
);
CREATE UNIQUE INDEX post_id_uindex ON post (id);