USE todo-mis;

CREATE TABLE IF NOT EXISTS `todo_items` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `title` varchar(100) NOT NULL,
    `description` varchar(255) NOT NULL,
    `status` tinyint(5) DEFAULT 0,
    `user_id` int(11) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user_created_at` (`user_id`, `created_at`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='todo items table';



CREATE TABLE IF NOT EXISTS `users` (
      `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
      `username` varchar(50) NOT NULL,
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='user table';
