-- Check if 'cms' database exists; if not, create it
CREATE DATABASE IF NOT EXISTS demo
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_general_ci;

-- Use the 'demo' database
USE demo;

-- ACL
CREATE TABLE IF NOT EXISTS user
(
    `id`           bigint(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `email`        VARCHAR(255)                                   NOT NULL,
    `name`         VARCHAR(255)                                   NOT NULL,
    `password`     TEXT                                           NOT NULL COMMENT 'bcrypt(plain_password)',
    `created_at`   bigint(20) UNSIGNED                            NOT NULL,
    `created_by`   bigint(20) UNSIGNED                            NOT NULL,
    `updated_at`   bigint(20) UNSIGNED                            NOT NULL,
    `updated_by`   bigint(20) UNSIGNED                            NOT NULL,
    UNIQUE KEY `idx_unique_email` (`email`) USING HASH,
    FULLTEXT KEY `idx_fulltext_email` (`email`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_general_ci;

insert into user (`id`, `email`, `name`, `password`, `created_at`, `created_by`, `updated_at`, `updated_by`)
values (1, 'john.doe@example.com', 'John Doe', 'secret', 1758118390, 0, 1758118390, 0);