ALTER USER 'root' IDENTIFIED WITH mysql_native_password BY 'root';

CREATE DATABASE `metrics_demo` charset=utf8mb4;

USE `metrics_demo`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `age` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4; 


INSERT INTO `user`(name, age) VALUES('zhao', 18);
INSERT INTO `user`(name, age) VALUES('qian', 21);
INSERT INTO `user`(name, age) VALUES('sun', 20);
INSERT INTO `user`(name, age) VALUES('li', 19);
