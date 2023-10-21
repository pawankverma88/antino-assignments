-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.37-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             9.5.0.5196
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for blogs
DROP DATABASE IF EXISTS `blogs`;
CREATE DATABASE IF NOT EXISTS `blogs` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `blogs`;

-- Dumping structure for table blogs.tbl_blogs
DROP TABLE IF EXISTS `tbl_blogs`;
CREATE TABLE IF NOT EXISTS `tbl_blogs` (
  `blog_id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'Unique identifier for each blog post',
  `author` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'References the user who created the blog post',
  `title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Title of the blog post',
  `content` text COLLATE utf8_unicode_ci COMMENT 'The main content of the blog post',
  `creation_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the blog post was created',
  `last_modified_date` datetime DEFAULT NULL COMMENT 'Date and time of the last modification',
  `status` varchar(2) COLLATE utf8_unicode_ci DEFAULT 'Y',
  PRIMARY KEY (`blog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='This table stores user generated blog posts';

-- Dumping data for table blogs.tbl_blogs: ~0 rows (approximately)
/*!40000 ALTER TABLE `tbl_blogs` DISABLE KEYS */;
/*!40000 ALTER TABLE `tbl_blogs` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
