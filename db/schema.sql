-- MySQL dump 10.13  Distrib 5.5.57, for debian-linux-gnu (x86_64)
--
-- Host: mysql    Database: default
-- ------------------------------------------------------
-- Server version	5.7.19

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
-- Table structure for table `blog`
--

DROP TABLE IF EXISTS `blog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `blog_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog`
--

LOCK TABLES `blog` WRITE;
/*!40000 ALTER TABLE `blog` DISABLE KEYS */;
INSERT INTO `blog` VALUES ('092982e7-b73f-4c8c-9687-7f8ba39b41ef','what up'),('13598cfe-a997-4bfb-8048-22882ccd48c1','what up'),('1bbbe7da-e115-4f30-884e-61eb63a095aa','yea'),('1c349165-7bf9-4ed6-82a9-a91ff0e406d4','yea'),('20e2c537-4e48-4bc0-a810-256b4b290d3f','what up'),('2609887a-68f2-4c0f-b30c-cbe88789f71e','what up'),('27a2579b-960c-4ed2-92cb-f18884280c51','what up'),('2c21c073-130c-4542-ad27-741db9b32d49','what up'),('43145eb1-7bbc-4b48-a64f-9422f09ccf20','yea'),('433f9a99-493f-40c1-8612-d1692e6dd478','what up'),('4b59d66e-b44f-44fc-bac8-941d0c53b155','yea'),('5419d02f-4f7f-4d52-bf25-08ef4aeba0c5','yea'),('67c3cda6-18de-4f86-956d-4a133685a5b4','what up'),('6de897b3-650f-4a81-8d96-55a2a7f92fc6','what up'),('7841e44c-5f7f-4bb6-a0ee-9b517b1a9d5b','yea'),('7c21e9e0-1bc0-459f-848a-8d2546f28874','what up'),('875256d7-517a-4f57-9696-fee8a073a3f4','what up'),('97e4813d-a3cf-46f5-a885-68eb404b4f96','what up'),('bba9d57a-7626-4671-8c33-26a57fbfe1a9','what up'),('c920d427-da82-42be-b3ba-ee812aab8da1','yea'),('df5bf3e6-6bd9-4747-8cdd-ce9afcf61f61','what up'),('e0c66835-182a-4248-bf68-fffcd552bd33','what up'),('e33d53b2-6b44-4aad-a218-dffeca0200b8','yea'),('e88b224e-b4c1-4e4e-bc11-1e9d5be3bd45','yea'),('f3acd27a-0a38-4eaa-bdaa-3d9f1bb1cfaf','yea'),('f699d974-5ea7-45be-b132-d478b80d29a9','what up'),('f8d37ae8-8be5-45eb-b83c-33ae9b77f3c3','yea'),('fb68ff02-ca9e-4dbf-b64e-c4225d6c7c8e','what up'),('fea60ee7-8d34-46a4-82f2-2b2c50acf5ad','what up');
/*!40000 ALTER TABLE `blog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post` (
  `id` varchar(255) NOT NULL,
  `body` text,
  `blog_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `post_id_uindex` (`id`),
  KEY `post_blog_id_fk` (`blog_id`),
  CONSTRAINT `post_blog_id_fk` FOREIGN KEY (`blog_id`) REFERENCES `blog` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES ('17394657-0f05-45ff-aedd-8eda30afe00e','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('212e80ba-2ff1-4d8a-be57-7b0033b13a60','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('23135e87-7aaf-4924-9576-162b8c21f7e9','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('23841b78-90b1-4a71-80a4-91519f519f25','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('272df1fd-3289-4e0f-ab6c-8fdf87a68a12','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('28d16246-a4dd-4e47-a318-bf4fcccc0980','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('32a34fe0-f4fe-48fb-9e71-1f0dd4a500a8','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('34e33459-f623-4355-8cc1-11f6747273c7','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('3cb2613a-bed3-465e-b58f-fe81eafa52ea','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('3fc6eff5-ffb3-4842-bdb3-3d940cc4847f','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('4a3fd59b-c93c-4a50-ae6b-490a08703ca2','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('580ac047-5210-482f-bbc2-3597d0daa3d5','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('5e396d13-0df5-4f6a-9bc6-1132044c6e57','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('72c0cdd6-c516-4257-97f6-42f4e06c41a9','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('83224d22-f1fd-4f1e-b99d-cf520d73acee','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('8b9e27ef-dcaa-420b-9dc3-f289acb618c1','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('8f173c00-ad4b-48cd-9285-c7b3c4247b94','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('994e167f-861e-43cf-a7f3-d9a17e465f9b','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('9a1e8377-ff39-4fe5-9ac0-feb6fa7eb475','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('9f0cabcd-20e5-4867-ac64-250366239c5b','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('a5f46052-bbb2-46f4-baff-d79f5be2f0d1','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('a6d31327-7fbb-4ed3-a1d7-f78e2e9f40b2','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('a8b69098-035b-4d38-8e8a-eb242f59e07e','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('b6f5820c-a592-466b-aa41-cbe6191b23ae','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('b996eb8a-ab64-4d35-9ffb-2523b1a984f9','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('c7d41ac5-9610-47b8-8f07-1137e3516d29','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('cc62517f-37c9-4e20-8586-7b2dbeb1cae0','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('dd177ff9-ae5c-4c19-8064-0200a9d0a7ab','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('de2ea835-6914-4e76-ab52-5e5f8264897e','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4'),('fc868c1f-46d4-4974-adb2-6ac0f2cea862','this is my blog post','875256d7-517a-4f57-9696-fee8a073a3f4');
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-10-15 19:57:27
