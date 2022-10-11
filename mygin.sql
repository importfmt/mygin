/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.30.100
 Source Server Type    : MariaDB
 Source Server Version : 50565
 Source Host           : 192.168.30.100:3306
 Source Schema         : mygin

 Target Server Type    : MariaDB
 Target Server Version : 50565
 File Encoding         : 65001

 Date: 21/12/2020 12:44:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cars
-- ----------------------------
DROP TABLE IF EXISTS `cars`;
CREATE TABLE `cars`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `brand` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `license` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `dead_weight` int(10) UNSIGNED NOT NULL,
  `city` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `license`(`license`) USING BTREE,
  INDEX `idx_cars_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of cars
-- ----------------------------
INSERT INTO `cars` VALUES (1, '2020-06-24 06:01:30', '2020-12-06 03:21:07', NULL, 'Benz', '京A66111', 1, 210, '上海市');
INSERT INTO `cars` VALUES (2, '2020-06-24 06:02:03', '2020-12-05 09:15:14', NULL, 'bmw', '闽A82888', 0, 200, '上海市');
INSERT INTO `cars` VALUES (4, '2020-06-24 07:29:16', '2020-12-06 03:19:47', NULL, 'bmw', '闽A99999', 0, 500, '福建省福州市');
INSERT INTO `cars` VALUES (5, '2020-06-24 08:05:30', '2020-07-08 11:28:02', '2020-07-08 11:59:48', 'bmw', '闽A82488', 0, 0, '福建省福州市');
INSERT INTO `cars` VALUES (6, '2020-06-24 08:05:36', '2020-12-05 11:46:31', NULL, 'Audi', '京A88888', 0, 0, '浙江省杭州市');
INSERT INTO `cars` VALUES (7, '2020-06-24 08:05:43', '2020-12-05 14:25:46', NULL, 'bmw', '闽A82048', 1, 1000, '福建省福州市');
INSERT INTO `cars` VALUES (8, '2020-06-24 08:33:59', '2020-12-05 08:14:04', NULL, 'Benz', '京A66611', 0, 3000, '北京市海淀区');
INSERT INTO `cars` VALUES (9, '2020-06-24 08:36:32', '2020-08-09 05:52:20', NULL, 'bmw', '闽A84848', 0, 0, '福建省福州市');
INSERT INTO `cars` VALUES (10, '2020-06-24 08:40:13', '2020-08-09 05:52:20', NULL, 'bmw', '闽A84998', 0, 0, '浙江省杭州市');
INSERT INTO `cars` VALUES (11, '2020-06-24 08:41:29', '2020-08-05 16:11:29', NULL, 'bmw', '闽B84998', 0, 0, '浙江省杭州市');
INSERT INTO `cars` VALUES (12, '2020-06-24 08:45:12', '2020-07-08 08:28:20', NULL, 'bmw', '闽C84998', 0, 0, '浙江省杭州市');
INSERT INTO `cars` VALUES (13, '2020-06-24 08:45:17', '2020-07-08 11:59:08', '2020-08-05 07:45:53', 'bmw', '闽C82933', 0, 200, '福建省福州市');
INSERT INTO `cars` VALUES (14, '2020-06-25 08:07:02', '2020-12-06 02:07:25', NULL, 'bmw', '闽C82198', 0, 500, '福建省福州市');
INSERT INTO `cars` VALUES (15, '2020-07-08 08:47:16', '2020-12-05 08:14:21', NULL, 'bmw', '闽C13198', 0, 100, '北京市');
INSERT INTO `cars` VALUES (16, '2020-07-08 09:05:02', '2020-07-08 09:05:02', '2020-08-09 07:28:40', '一汽丰田', '京A20233', 0, 200, '北京市');
INSERT INTO `cars` VALUES (17, '2020-08-09 07:28:31', '2020-12-06 00:43:06', NULL, 'beijing', '京A111111', 1, 200, '北京市');
INSERT INTO `cars` VALUES (18, '2020-09-05 11:32:33', '2020-12-05 13:55:05', NULL, 'beaz', '京A12222', 0, 200, '北京市');
INSERT INTO `cars` VALUES (19, '2020-08-05 16:09:27', '2020-12-06 02:18:03', NULL, 'benz', '闽A6666', 1, 1000, '福建省福州市');
INSERT INTO `cars` VALUES (20, '2020-12-06 02:18:59', '2020-12-06 02:32:25', NULL, 'benz', '川A11111', 0, 10000, '福建省福州市');

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `price` int(10) UNSIGNED NOT NULL,
  `weight` int(10) UNSIGNED NOT NULL,
  `courier_status` tinyint(1) NOT NULL,
  `courier_number` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `from_city` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `to_city` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_goods_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '2020-06-25 12:55:45', '2020-12-05 09:15:14', NULL, '汽车配件', 500, 2, 1, '74054753996895518159', '福建省福州市', '上海市', 1);
INSERT INTO `goods` VALUES (2, '2020-06-25 12:56:14', '2020-12-06 03:21:07', NULL, 'Apple iphone12', 4999, 120, 0, '24945914427282667781', '上海市', '北京市', 0);
INSERT INTO `goods` VALUES (3, '2020-07-14 07:22:43', '2020-12-06 00:43:06', NULL, '小米笔记本Pro', 2000, 200, 0, '06519356175501339396', '北京市', '北京市', 0);
INSERT INTO `goods` VALUES (4, '2020-07-14 07:23:24', '2020-12-05 13:01:26', NULL, '测试货物', 1111, 1111, 1, '', '北京市', '福建省福州市', 0);
INSERT INTO `goods` VALUES (5, '2020-07-12 22:04:15', '2020-12-06 03:19:47', NULL, 'aaaa', 200, 99, 0, '', '福建省福州市', '北京市', 0);
INSERT INTO `goods` VALUES (6, '2020-08-03 21:30:20', '2020-12-05 14:25:46', NULL, 'Apple iphone 11', 5000, 4, 0, '73045301944677546412', '福建省福州市', '上海市', 0);
INSERT INTO `goods` VALUES (7, '2020-12-06 02:06:04', '2020-12-06 02:18:03', NULL, 'fhdsakhf', 111, 111, 0, '77547825919357525792', '福建省福州市', '北京市', 0);
INSERT INTO `goods` VALUES (8, '2020-12-06 02:19:32', '2020-12-06 02:32:25', NULL, 'd11111', 1888, 9999, 1, '22561755186844487564', '重庆市', '福建省福州市', 1);

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `number` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `goods_id` int(10) UNSIGNED NOT NULL,
  `carrier_car_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `carrier_user_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `price` int(10) UNSIGNED NOT NULL,
  `weight` int(10) UNSIGNED NOT NULL,
  `courier_number` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `from_city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `to_city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `courier_status` tinyint(1) NOT NULL,
  `username` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `license` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `goodsname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `number`(`number`) USING BTREE,
  UNIQUE INDEX `courier_number`(`courier_number`) USING BTREE,
  INDEX `idx_orders_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (4, '2020-06-26 07:36:21', '2020-08-03 21:46:43', '2020-08-09 05:52:06', 'raKUGDKiCnWihwKOfdsn', 1, 1, 1, 500, 2, '93005505693985268989', '福建省福州市', '上海市', 0, 1, '', '', '');
INSERT INTO `orders` VALUES (5, '2020-06-26 09:21:35', '2020-07-14 06:55:12', '2020-08-09 05:43:05', 'FJwUkykgKVANimduMlah', 2, 2, 2, 5000, 4, '67562321611756740996', '福建省福州市', '上海市', 0, 0, '', '', '');
INSERT INTO `orders` VALUES (6, '2020-07-14 08:04:38', '2020-08-03 21:49:41', '2020-08-09 05:52:08', 'IVjwnPXXUwLrYUConXxP', 1, 2, 27, 500, 2, '93997979039332609511', '福建省福州市', '上海市', 0, 0, '', '', '');
INSERT INTO `orders` VALUES (7, '2020-07-12 22:04:32', '2020-08-05 16:12:09', '2020-12-06 03:19:47', 'oovMwAgqXuMzbMdenACh', 5, 4, 36, 200, 99, '08941150968244709162', '福建省福州市', '北京市', 0, 0, '', '', '');
INSERT INTO `orders` VALUES (8, '2020-08-09 05:49:50', '2020-08-09 05:49:50', '2020-08-09 05:50:11', 'jVLkNFlaCOomJtRLSbEX', 5, 2, 2, 200, 99, '99381073422652114423', '福建省福州市', '北京市', 0, 0, '', '', '');
INSERT INTO `orders` VALUES (9, '2020-08-09 05:51:40', '2020-08-09 05:51:40', '2020-08-09 05:51:54', 'TFTkBICofHLjNhJxoAVc', 5, 2, 2, 200, 99, '02224142467936192750', '福建省福州市', '北京市', 0, 0, '', '', '');
INSERT INTO `orders` VALUES (10, '2020-08-09 07:47:56', '2020-08-09 07:48:16', NULL, 'VsYFyErRBEgYANWzSLps', 5, 2, 2, 200, 99, '10187055481737960114', '福建省福州市', '北京市', 1, 1, 'vWLIuhiajV', '闽A82888', 'aaaa');
INSERT INTO `orders` VALUES (11, '2020-08-05 16:06:40', '2020-08-05 16:06:40', '2020-08-05 16:46:21', 'ikiCvEOlJLUJPkZTmoDy', 1, 2, 2, 500, 2, '54347403916318046857', '福建省福州市', '上海市', 1, 1, 'vWLIuhiajV', '闽A82888', '汽车配件');
INSERT INTO `orders` VALUES (12, '2020-08-05 16:53:06', '2020-12-05 05:04:03', NULL, 'JTuxKIpGNyCUivWXMIjr', 1, 2, 2, 500, 2, '74054753996895518159', '福建省福州市', '上海市', 1, 1, 'vWLIuhiajV', '闽A82888', '汽车配件');
INSERT INTO `orders` VALUES (13, '2020-08-05 16:53:06', '2020-12-05 09:15:14', NULL, 'JTufjfijNyCUivWXMIjr', 1, 2, 43, 300, 200, '74054753988885518159', '福建省福州市', '上海市', 1, 1, '测试账号', '闽A82888', '汽车配件');
INSERT INTO `orders` VALUES (14, '2020-12-05 14:15:59', '2020-12-05 14:15:59', NULL, 'CvMAOqmPHEKDepwlLNjp', 5, 4, 36, 200, 99, '69892867444965933537', '福建省福州市', '北京市', 0, 0, 'aaaaa', '闽A99999', 'aaaa');
INSERT INTO `orders` VALUES (15, '2020-12-05 14:25:46', '2020-12-05 14:25:46', NULL, 'vYUkJdoLpTSeYNDZqSNu', 6, 7, 42, 5000, 4, '73045301944677546412', '福建省福州市', '上海市', 0, 0, 'fdgkdsjg', '闽A82048', 'Apple iphone 11');
INSERT INTO `orders` VALUES (16, '2020-12-05 14:57:21', '2020-12-05 14:57:21', '2020-12-06 03:21:01', 'vvBNfbnqxwktBCpFsXlz', 2, 1, 2, 4999, 120, '35454056590248702536', '上海市', '北京市', 0, 0, 'vWLIuhiajV', '京A66111', 'Apple iphone12');
INSERT INTO `orders` VALUES (17, '2020-12-06 00:43:06', '2020-12-06 00:43:06', NULL, 'FHjlGxRQVniwORiJxidM', 3, 17, 37, 2000, 200, '06519356175501339396', '北京市', '北京市', 0, 0, 'cYfPTGbhnV', '京A111111', '小米笔记本Pro');
INSERT INTO `orders` VALUES (18, '2020-12-06 02:18:03', '2020-12-06 02:18:03', NULL, 'dngHjChDAdiwpRPqwRYu', 7, 19, 39, 111, 111, '77547825919357525792', '福建省福州市', '北京市', 0, 0, 'fjhsdkfh', '闽A6666', 'fhdsakhf');
INSERT INTO `orders` VALUES (19, '2020-12-06 02:31:55', '2020-12-06 02:32:25', NULL, 'cqYUiIPLgsqhoqeTWDOZ', 8, 20, 44, 1888, 9999, '22561755186844487564', '重庆市', '福建省福州市', 1, 1, 'ddfsfsdf', '川A11111', 'd11111');
INSERT INTO `orders` VALUES (20, '2020-12-06 03:21:07', '2020-12-06 03:21:07', NULL, 'EbjFNDJBkEdeEBcUigmV', 2, 1, 43, 4999, 120, '24945914427282667781', '上海市', '北京市', 0, 0, '测试账号', '京A66111', 'Apple iphone12');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `username` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `role` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `city` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 46 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2020-06-23 12:00:57', '2020-12-06 02:05:33', NULL, 'gUlnndFLjK', '$2a$10$p8E1qp5duz4yMO2Ece62x.P4leXTTcoTjt0C2v8zgP0.kjRiMv2w.', '1123226535@qq.com', '13331233323', '货物管理员', 1, '上海市');
INSERT INTO `users` VALUES (2, '2020-06-23 12:18:27', '2020-12-06 03:21:58', NULL, 'vWLIuhiajV', '$2a$10$ReT7tl0LuRmVA2fS9yw2cO79dmcNiIQkmejU9IoFWli9Y/c35mlLe', 'fhfsdafsaf@qq.com', '15000000000', '超级管理员', 1, '上海市');
INSERT INTO `users` VALUES (21, '2020-06-24 07:13:10', '2020-12-06 01:08:28', '2020-12-06 01:08:33', 'xcADqcVqEZ', '$2a$10$4tizB.w.7fVIwBQVwjBCf.X07qgdSeFmpS9UYXv03ufykRn2qc282', 'f1111d6535@qq.com', '13333333323', '订单管理员', 0, '福建省福州市');
INSERT INTO `users` VALUES (26, '2020-06-24 07:42:27', '2020-08-09 05:48:58', NULL, 'pQfOhUqibM', '$2a$10$l9bq3jIxNQxx5qWat19NZOA4TfRx2eI6oM5jS.qphwRd4dwDjudyS', '2ads22fas@qq.com', '15000002001', '普通用户', 1, '福建省漳州市');
INSERT INTO `users` VALUES (27, '2020-06-24 07:42:58', '2020-08-05 16:53:06', NULL, 'lQVzwkMsMv', '$2a$10$jq6Dco26y8wT3pMxq/BAje2ziIVUXHUpCXqZDMU.4cus/fs0KEPq2', '2ads2f5as@qq.com', '15000332000', '普通用户', 1, '福建省福州市');
INSERT INTO `users` VALUES (28, '2020-06-27 07:11:37', '2020-12-06 02:05:35', NULL, 'admin', '$2a$10$KeRk4Bpa0ZqzMyrtDQgIYu56TAOBwlce6jjHL741xOaxDvph/Gso2', 'admin@qq.com', '15011111111', '超级管理员', 1, '北京市');
INSERT INTO `users` VALUES (29, '2020-06-27 08:17:00', '2020-12-06 02:05:36', NULL, 'fasklhfk', '$2a$10$sGPAD4/4gyUUfbnUv1ROeOGz/JgvX/GjeFvZQORpHn4Zif7lofxJG', '2352412423@qq.com', '13500019111', '用户管理员', 1, '山东省威海市');
INSERT INTO `users` VALUES (30, '2020-06-27 08:18:42', '2020-07-08 06:48:30', '2020-07-08 11:59:30', 'dsajf', '$2a$10$Ms7W6XF0P/Ap79zdDYRbiOvv9hMxLzR3gzWqUzhCCd308.e5nInxu', '787429@qq.com', '15032333333', '普通用户', 0, '黑龙江省哈尔滨市');
INSERT INTO `users` VALUES (31, '2020-06-27 08:19:43', '2020-12-05 10:41:23', NULL, '王五六', '$2a$10$CfOetFkfX7K8IhqjFLEMMu2e79DkuHcFqxu87MzMVQU2S.y.2IJrS', 'f483f3h@163.com', '18836474442', '车辆管理员', 1, '湖南省长沙市');
INSERT INTO `users` VALUES (32, '2020-06-27 08:20:45', '2020-12-06 02:05:37', NULL, 'hredhg', '$2a$10$QaTLu89v066XuDLB5jcJYO9iIj5pGXYed6tp07aoVatoE1T1vKRN2', '11112@qq.com', '13666666666', '车辆管理员', 1, '上海市');
INSERT INTO `users` VALUES (33, '2020-06-27 08:21:12', '2020-07-14 04:58:15', '2020-12-05 09:51:35', 'fisdfb', '$2a$10$25ndG5M4JTzYJ0f7m9bVGe/3gUYiKAnlB/FfBk2yPGH145YwtrpGu', '7346286482@qq.com', '18988883733', '测试用户', 0, '广东省深圳市');
INSERT INTO `users` VALUES (34, '2020-06-27 08:21:45', '2020-07-08 06:48:35', '2020-12-05 11:12:45', 'fjskdf', '$2a$10$L9qykYZhJpzWxdh4OEd5merM7hXLLplPEUk2ypt4Qr96OzOC91S3.', '34792@qq.com', '13333337777', '普通用户', 0, '新疆维吾尔自治区齐齐哈尔市');
INSERT INTO `users` VALUES (35, '2020-06-27 08:22:10', '2020-07-08 06:48:36', NULL, 'fdshkf', '$2a$2a$10$ReT7tl0LuRmVA2fS9yw2cO79dmcNiIQkmejU9IoFWli9Y/c35mlLe', 'hdskfh@163.com', '15566663821', '普通用户', 0, '重庆市');
INSERT INTO `users` VALUES (36, '2020-07-12 21:58:18', '2020-12-06 03:19:47', NULL, 'aaaaa', '$2a$10$hPZSSGlTMCTKsmIvzopxG.VlSN.P1zJmgtafcR0iOedj5BSS59VGG', '22222222@qq.com', '15047392731', '普通用户', 0, '福建省福州市');
INSERT INTO `users` VALUES (37, '2020-08-03 20:50:23', '2020-12-06 00:43:06', NULL, 'cYfPTGbhnV', '$2a$10$zx9KVdH1uGlUhOZa6P6NZOLjyAUtHrxlFejwIekloJ6VplmIbInn6', 'affdn@qq.com', '15000003240', '普通用户', 1, '北京市');
INSERT INTO `users` VALUES (38, '2020-08-03 22:18:19', '2020-08-05 07:29:48', NULL, 'wengweng', '$2a$10$Rq8V58HQZUf9jueuwE9V.uvAKpnhS23GbGXuvApQnAwrOggWFSW0O', '15111111111@189.com', '15111111111', '普通用户', 0, '北京市');
INSERT INTO `users` VALUES (39, '2020-08-03 22:57:06', '2020-12-06 02:18:03', NULL, 'fjhsdkfh', '$2a$10$.ssF0yK/TwxwDCjoSOo5HeDXwO9P7Yqv4I0oZXj5NyXWnpOKaAtFK', '16474637773@qq.com', '15023772924', '普通用户', 1, '福建省福州市');
INSERT INTO `users` VALUES (40, '2020-08-03 23:04:53', '2020-08-03 23:04:53', NULL, 'fsdfscc', '$2a$10$jzrLYTDbxKXme2AWL9Bwxu8dPAMxoCXPsONadUDxFgQX6Feyr2Eyu', 'hfs2js@qq.com', '15002332222', '普通用户', 0, '北京市');
INSERT INTO `users` VALUES (41, '2020-09-05 11:35:34', '2020-12-06 03:21:59', NULL, 'fdksjsa', '$2a$10$0pD3e5xwb7LNSIf6lmzRUuKuKhVeAHlmke1QXuXPlXnLf6HObjz.6', '2222222228@qq.com', '18000000011', '车辆管理员', 1, '福建省福州市');
INSERT INTO `users` VALUES (42, '2020-12-04 08:10:13', '2020-12-05 14:25:46', NULL, 'fdgkdsjg', '$2a$10$yjGQXIuViDb2kaXvzBGnpeItrLt/PzKraP7YqdvSsW4q92gjyV156', '72389749@qq.com', '13321321333', '普通用户', 1, '福建省福州市');
INSERT INTO `users` VALUES (43, '2020-12-04 08:11:44', '2020-12-06 03:21:07', NULL, '测试账号', '$2a$10$pm3YdnsskbSc7Qpq6hGwtuvwsRquiKAResBo3hQLWU66bagWPtdvq', '423337493@qq.com', '13655079483', '普通用户', 1, '上海市');
INSERT INTO `users` VALUES (44, '2020-12-06 02:31:33', '2020-12-06 02:32:25', NULL, 'ddfsfsdf', '$2a$10$.SMR5vElO9pFeMpahh1eAOq1JFlgHKDEccSz5lBC7fEEMjPNRMRye', 'dsfafs@qq.com', '13333333333', '普通用户', 0, '福建省福州市');
INSERT INTO `users` VALUES (45, '2020-12-06 02:54:56', '2020-12-06 02:54:56', NULL, 'sdfsfdsf', '$2a$10$9KuVRMCvuZg3Rcc8GnJ.iumKz5mE5H17c125AwarnXqu/IjhUNYlq', '2323@qq.com', '18888888888', '普通用户', 0, '福建省福州市');

-- ----------------------------
-- Table structure for wips
-- ----------------------------
DROP TABLE IF EXISTS `wips`;
CREATE TABLE `wips`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `reply` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `user_id` int(10) NOT NULL,
  `username` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_wips_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of wips
-- ----------------------------
INSERT INTO `wips` VALUES (1, '2020-08-09 14:47:56', '2020-08-09 09:24:43', NULL, '测试', '测试！', 'fdsahfksjdhfkjsahfkjsdhfjkshdjakfhsaklfhlskadfsdfa', 1, 1, '');
INSERT INTO `wips` VALUES (2, '2020-08-09 17:25:47', '2020-08-05 06:00:48', NULL, '111', '1111', 'hdskahdlashdjla', 1, 21, '');
INSERT INTO `wips` VALUES (3, '2020-08-09 09:36:02', '2020-08-09 09:39:00', NULL, '222', '2222', '11asasas1', 1, 3, '');
INSERT INTO `wips` VALUES (4, '2020-08-09 09:37:26', '2020-08-05 07:46:44', NULL, '222', '2222', 'this is test.', 1, 11, '');
INSERT INTO `wips` VALUES (5, '2020-12-03 11:10:26', '2020-12-05 08:51:24', NULL, 'wipInfo', 'wipInfo', '收到。', 1, 1, '');
INSERT INTO `wips` VALUES (6, '2020-12-03 11:28:37', '2020-12-03 11:28:37', NULL, 'dsfsafsaf', 'dfsafsadfa', 'fsdfdskjflksadjfldsajfl;sdjfskadfjd;slfjdsajf', 1, 2, '');
INSERT INTO `wips` VALUES (7, '2020-12-03 11:29:15', '2020-12-03 11:29:15', NULL, '测试', '测试', '', 0, 2, '');
INSERT INTO `wips` VALUES (8, '2020-12-03 11:56:32', '2020-12-03 11:56:32', NULL, 'fjskdahfshadfh', 'cecececece', '', 0, 2, 'vWLIuhiajV');
INSERT INTO `wips` VALUES (9, '2020-12-04 08:43:45', '2020-12-04 08:43:45', NULL, 'fgsdgfdgs', 'fgsdgfdsgdsfg', '', 0, 43, '测试账号');
INSERT INTO `wips` VALUES (10, '2020-12-04 08:44:09', '2020-12-04 08:44:09', NULL, 'dfasfdsaf', 'fdsafsdafsa', '', 0, 43, '测试账号');
INSERT INTO `wips` VALUES (11, '2020-12-04 09:14:42', '2020-12-05 09:27:03', NULL, 'fdsfdsa', 'dfsafdsaf', '已收到。', 1, 43, '测试账号');
INSERT INTO `wips` VALUES (12, '2020-12-05 05:43:18', '2020-12-05 08:09:23', NULL, '测试工单', 'fkdsjflksjfjfslkdjfl', '收到。', 1, 43, '测试账号');

SET FOREIGN_KEY_CHECKS = 1;
