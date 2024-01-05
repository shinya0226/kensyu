-- MySQL dump 10.13  Distrib 5.7.43, for osx10.19 (x86_64)
--
-- Host: 127.0.0.1    Database: yamamoto
-- ------------------------------------------------------
-- Server version	5.7.43

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
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('shinya.yamamoto6@persol-pt.co.jp','$2a$10$/NQlLMzK4kNEa8kyi/RHEO82UNuvIvwaJtAYW9F2puBUkfo/W5e3C','山本真也',1),('kentaro.suzuki@persol-pt.co.jp','$2a$10$vScJivjRkO4N/E5EYLr6qui1SeZpB58lu8c5nfVeTGwzl8yNGu2g2','鈴木健太郎',0),('rio.matumoto@persol-pt.co.jp','$2a$10$GcVn7sgAXTsilTf9RSX1k.N6AclvYUME5p6aOxYZlVoUCkVb7juAO','松本莉央',1),('tsutom.utida@persol-pt.co.jp','$2a$10$cK0AbZvBCi06Xtge2Dazwe/328QJidLOoXyCSg0bUHvwOuJ9tOV2q','内田勉',0),('kyousuke.mizuno@persol-pt.co.jp','$2a$10$d4lKmw55MwQyj8ejxtwzpuW5FVgAuICWSwI62KHsXRx5Laa/x5CRq','水野京介',1),('makoto.ueda@persol-pt.co.jp','$2a$10$zWPyycsmMI5DESxS/m/SpehLvR92uYvAKpC/DH.ozBruEobFCzzna','上田誠',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-05 12:57:51
