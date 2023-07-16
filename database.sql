-- Adminer 4.8.1 MySQL 5.5.5-10.7.3-MariaDB-1:10.7.3+maria~focal dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE TABLE `categories` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(100) DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `categories_name_IDX` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

INSERT INTO `categories` (`id`, `name`) VALUES
                                            (1,	'Serba Ayam'),
                                            (5,	'Serba Ikan'),
                                            (6,	'Serba Udang');

CREATE TABLE `ingredients` (
                               `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                               `recipe_id` bigint(20) unsigned DEFAULT NULL,
                               `type` longtext DEFAULT NULL,
                               `name` longtext DEFAULT NULL,
                               `sort` tinyint(3) unsigned DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `fk_recipes_ingredients` (`recipe_id`),
                               CONSTRAINT `fk_recipes_ingredients` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4;

INSERT INTO `ingredients` (`id`, `recipe_id`, `type`, `name`, `sort`) VALUES
                                                                          (1,	1,	'content',	'4 buah roti burger',	0),
                                                                          (2,	1,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (3,	1,	'header',	'Chicken Patty',	2),
                                                                          (4,	1,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (5,	1,	'content',	'150 g lemak ayam, potong-potong',	4),
                                                                          (6,	2,	'content',	'4 buah roti burger',	0),
                                                                          (7,	2,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (8,	2,	'header',	'Chicken Patty',	2),
                                                                          (9,	2,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (10,	2,	'content',	'150 g lemak ayam, potong-potong',	4),
                                                                          (11,	NULL,	'content',	'4 buah roti burger',	0),
                                                                          (12,	NULL,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (13,	NULL,	'header',	'Chicken Patty',	2),
                                                                          (14,	NULL,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (15,	NULL,	'content',	'150 g lemak ayam, potong-potong',	4),
                                                                          (16,	NULL,	'content',	'4 buah roti burger',	0),
                                                                          (17,	NULL,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (18,	NULL,	'header',	'Chicken Patty',	2),
                                                                          (19,	NULL,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (20,	NULL,	'content',	'150 g lemak ayam, potong-potong',	4),
                                                                          (21,	NULL,	'content',	'4 buah roti burger',	0),
                                                                          (22,	NULL,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (23,	NULL,	'header',	'Chicken Patty',	2),
                                                                          (24,	NULL,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (25,	NULL,	'content',	'150 g lemak ayam, potong-potong',	4),
                                                                          (26,	NULL,	'content',	'4 buah roti burger',	0),
                                                                          (27,	NULL,	'content',	'100 g lettuce, iris kasar',	1),
                                                                          (28,	NULL,	'header',	'Chicken Patty',	2),
                                                                          (29,	NULL,	'content',	'600 g fillet dada ayam, potong-potong',	3),
                                                                          (30,	NULL,	'content',	'150 g lemak ayam, potong-potong',	4);

CREATE TABLE `recipes` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `category_id` bigint(20) unsigned DEFAULT NULL,
                           `title` longtext DEFAULT NULL,
                           `description` longtext DEFAULT NULL,
                           `cooking_time` tinyint(3) unsigned DEFAULT NULL,
                           `thumbnail_url` longtext DEFAULT NULL,
                           `video_url` longtext DEFAULT NULL,
                           `created_at` datetime(3) DEFAULT NULL,
                           `updated_at` datetime(3) DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `fk_recipes_category` (`category_id`),
                           CONSTRAINT `fk_recipes_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

INSERT INTO `recipes` (`id`, `category_id`, `title`, `description`, `cooking_time`, `thumbnail_url`, `video_url`, `created_at`, `updated_at`) VALUES
                                                                                                                                                  (1,	1,	'Chicken Burger Ala Mcd',	'Buat sendiri chicken burger ala restoran cepat saji di rumah. Dengan bahan dan bumbu yang dibuat semirip mungkin dengan aslinya. Nikmati burger yang lebih lezat tanpa mengantri!',	40,	'https://dummyimage.com/300x300/000/fff&text=Chicken+Burger+Ala+Mcd',	'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',	'2023-07-16 21:03:50.747',	'2023-07-16 21:03:50.747'),
                                                                                                                                                  (2,	1,	'Chicken Burger Ala Mcd',	'Buat sendiri chicken burger ala restoran cepat saji di rumah. Dengan bahan dan bumbu yang dibuat semirip mungkin dengan aslinya. Nikmati burger yang lebih lezat tanpa mengantri!',	40,	'https://dummyimage.com/300x300/000/fff&text=Chicken+Burger+Ala+Mcd',	'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',	'2023-07-16 21:20:45.359',	'2023-07-16 21:20:45.359');

CREATE TABLE `steps` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `recipe_id` bigint(20) unsigned DEFAULT NULL,
                         `type` longtext DEFAULT NULL,
                         `text` text DEFAULT NULL,
                         `sort` tinyint(3) unsigned DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `fk_recipes_steps` (`recipe_id`),
                         CONSTRAINT `fk_recipes_steps` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

INSERT INTO `steps` (`id`, `recipe_id`, `type`, `text`, `sort`) VALUES
                                                                    (1,	1,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (2,	1,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (3,	1,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2),
                                                                    (4,	2,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (5,	2,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (6,	2,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2),
                                                                    (7,	NULL,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (8,	NULL,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (9,	NULL,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2),
                                                                    (10,	NULL,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (11,	NULL,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (12,	NULL,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2),
                                                                    (13,	NULL,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (14,	NULL,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (15,	NULL,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2),
                                                                    (16,	NULL,	'content',	'Chicken patty: Dalam food processor, masukkan potongan ayam dan lemak ayam. Proses hingga halus',	0),
                                                                    (17,	NULL,	'content',	'Tambahkan garam, gula, kaldu ayam, merica, seledri, sage, bawang putih, dan bawang bombai bubuk. Proses hingga rata.',	1),
                                                                    (18,	NULL,	'content',	'Bagi adonan patty jadi 4 bagian, bentuk bulat pipih. Simpan dalam kulkas selama 30-45 menit.',	2);

-- 2023-07-16 14:45:17