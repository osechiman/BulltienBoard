-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: BulltienBoard
-- ------------------------------------------------------
-- Server version	5.5.5-10.9.2-MariaDB-1:10.9.2+maria~ubu2204

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `BulltienBoard`
--

DROP TABLE IF EXISTS `BulltienBoard`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `BulltienBoard` (
  `id` varchar(36) NOT NULL COMMENT 'BulltienBoardのID\r\nUUIDを期待しています',
  `title` varchar(50) NOT NULL COMMENT 'BulltienBoardのタイトル',
  `thread_id` varchar(36) NOT NULL COMMENT 'ThreadのID\r\nBullitienBoardにぶら下がるThreadのIDが格納されます',
  PRIMARY KEY (`id`),
  KEY `BulltienBoard_FK` (`thread_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='BulletinBoard情報を保存するためのテーブル';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Comment`
--

DROP TABLE IF EXISTS `Comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Comment` (
  `id` varchar(36) NOT NULL COMMENT 'CommentのID\r\nUUIDを期待しています',
  `thread_id` varchar(36) NOT NULL COMMENT 'ThreadのID\r\nUUIDを期待しています',
  `text` varchar(2048) NOT NULL COMMENT 'Commentの内容',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'Comment作成日時',
  PRIMARY KEY (`id`),
  KEY `Comment_FK` (`thread_id`),
  KEY `Comment_id_created_at_asc` (`id`,`created_at`) USING BTREE,
  KEY `Comment_id_created_at_desc` (`id`,`created_at` DESC) USING BTREE,
  CONSTRAINT `Comment_FK` FOREIGN KEY (`thread_id`) REFERENCES `Thread` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Comment情報を保存するためのテーブル';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Thread`
--

DROP TABLE IF EXISTS `Thread`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Thread` (
  `id` varchar(36) NOT NULL COMMENT 'ThreadのID\r\nUUIDを期待しています',
  `bulltien_board_id` varchar(36) NOT NULL COMMENT 'BulltienBoardのID\r\nUUIDを期待しています',
  `title` varchar(50) NOT NULL COMMENT 'Threadのタイトル',
  `comment_id` varchar(36) NOT NULL COMMENT 'CommentのID\r\nUUIDを期待しています',
  PRIMARY KEY (`id`),
  KEY `Thread_FK_1` (`comment_id`),
  KEY `Thread_FK` (`bulltien_board_id`),
  CONSTRAINT `Thread_FK` FOREIGN KEY (`bulltien_board_id`) REFERENCES `BulltienBoard` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Thread情報を保存するためのテーブル';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'BulltienBoard'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-19 16:33:06
