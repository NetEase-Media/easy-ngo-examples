ALTER USER 'root' IDENTIFIED WITH mysql_native_password BY 'root';

CREATE DATABASE `gorm_demo` charset=utf8mb4;

USE `gorm_demo`;

CREATE TABLE `blacklist` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `blackword` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4; 


INSERT INTO `blacklist`(`blackword`) VALUES('initial');
