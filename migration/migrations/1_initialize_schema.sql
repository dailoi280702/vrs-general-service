-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,

    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS videos
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    likes BIGINT NOT NULL DEFAULT 0,
    comments BIGINT NOT NULL DEFAULT 0,
    shares BIGINT NOT NULL DEFAULT 0,
    length BIGINT NOT NULL DEFAULT 0,
    watch_time BIGINT NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE users;
DROP TABLE videos;
