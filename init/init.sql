CREATE DATABASE `capstone` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

USE `capstone`;

-- capstone.divisions definition

CREATE TABLE `divisions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- capstone.users definition

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `division_id` bigint unsigned DEFAULT NULL,
  `role` varchar(20),
  `name` longtext,
  `email` longtext,
  `phone` longtext,
  `password` longtext,
  `gender` varchar(6),
  `avatar` longtext,
  `address` longtext,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `division_id` (`division_id`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`division_id`) REFERENCES `divisions` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- capstone.categories definition

CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- capstone.assets definition

CREATE TABLE `assets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
<<<<<<< HEAD
  `code` varchar(100),
=======
  `code` varchar(25),
>>>>>>> main
  `short_name` longtext,
  `name` longtext,
  `category_id` bigint unsigned DEFAULT NULL,
  `description` longtext,
  `status` longtext,
  `image` longtext,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `assets_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- capstone.borrow/return_requests definition

CREATE TABLE `borrowORreturn_requests` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL,
  `asset_id` bigint unsigned DEFAULT NULL,
  `activity` longtext,
  `status` longtext,
  `request_time` datetime,
  `return_time` datetime,
  `description` longtext,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`user_id`),
  KEY `asset_id` (`asset_id`),
  CONSTRAINT `borrow/return_requests_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `borrow/return_requests_ibfk_2` FOREIGN KEY (`asset_id`) REFERENCES `assets` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- capstone.procurement_requests definition

CREATE TABLE `procurement_requests` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL,
  `category_id` bigint unsigned DEFAULT NULL,
  `activity` longtext,
  `status` longtext,
  `request_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `description` longtext,
  `image` longtext,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `user_id` (`user_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `procurement_requests_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `procurement_requests_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;