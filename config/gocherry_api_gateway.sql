/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost:3306
 Source Schema         : gocherry_api_gateway

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 29/08/2020 12:54:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_account
-- ----------------------------
DROP TABLE IF EXISTS `admin_account`;
CREATE TABLE `admin_account` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `phone` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号',
  `pwd` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `level` int(10) NOT NULL DEFAULT '1' COMMENT '等级1：普通 2：研发 3：超级管理员',
  `state` int(2) NOT NULL DEFAULT '1' COMMENT '状态1:可用',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`phone`,`pwd`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of admin_account
-- ----------------------------
BEGIN;
INSERT INTO `admin_account` VALUES (1, 'sgg', '110', '202cb962ac59075b964b07152d234b70', '2000-01-01 00:00:00', '2000-01-01 00:00:00', 1, 3, NULL);
INSERT INTO `admin_account` VALUES (2, 'jjj', '333333333', 'b857eed5c9405c1f2b98048aae506792', '2020-08-28 17:35:34', '2020-08-28 17:35:34', 1, 0, NULL);
INSERT INTO `admin_account` VALUES (3, '3333333', '44444444444', '0b5de470bdace90bd6cfb2541eb79f99', '2020-08-28 17:37:29', '2020-08-28 17:37:29', 1, 0, NULL);
INSERT INTO `admin_account` VALUES (4, 'admin', '111', 'e10adc3949ba59abbe56e057f20f883e', '2020-08-28 17:39:26', '2020-08-28 17:39:26', 2, 0, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
