CREATE DATABASE eventori_tech_test;

CREATE TABLE `model_catalogues` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Model_name` varchar(50) DEFAULT NULL,
  `Image_url` text,
  `Description` text,
  `Created_at` timestamp NULL DEFAULT NULL,
  `Updated_at` timestamp NULL DEFAULT NULL,
  `Deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `model_schedules` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Model_id` int NOT NULL,
  `Schedule_date` timestamp NULL DEFAULT NULL,
  `Created_at` timestamp NULL DEFAULT NULL,
  `Updated_at` timestamp NULL DEFAULT NULL,
  `Deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `fk_model_schedules_model_catalogues` (`Model_id`),
  CONSTRAINT `fk_model_schedules_model_catalogues` FOREIGN KEY (`Model_id`) REFERENCES `model_catalogues` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

