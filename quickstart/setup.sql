ALTER USER 'root' IDENTIFIED WITH mysql_native_password BY 'root';

CREATE DATABASE `user_demo` charset=utf8mb4;

USE `user_demo`;

CREATE TABLE `user_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `passport` varchar(50) NOT NULL  DEFAULT '',
  `gender` tinyint NOT NULL DEFAULT 1,
  `address` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4; 


INSERT INTO `user_info`(`name`, `passport`, `gender`, `address`) VALUES('Darcy', '12', 1, "Shandong");
INSERT INTO `user_info`(`name`, `passport`, `gender`, `address`) VALUES('Jane', '13', 2, "Shanxi");
