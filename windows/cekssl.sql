/*
 Navicat Premium Data Transfer

 Source Server         : MYSQL 8 Lokal
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3308
 Source Schema         : cekssl

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 03/02/2023 05:06:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for domain
-- ----------------------------
DROP TABLE IF EXISTS `domain`;
CREATE TABLE `domain`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `domain` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `domain_UNIQUE`(`domain` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of domain
-- ----------------------------
INSERT INTO `domain` VALUES (2, 'detik.com');
INSERT INTO `domain` VALUES (4, 'github.com');
INSERT INTO `domain` VALUES (1, 'google.com');
INSERT INTO `domain` VALUES (3, 'zaki.my.id');

SET FOREIGN_KEY_CHECKS = 1;
