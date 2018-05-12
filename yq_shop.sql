/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost:3306
 Source Schema         : yq_shop

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 12/05/2018 20:21:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `goods_id` int(11) NULL DEFAULT NULL,
  `goods_nums` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES (1, 1, 1, 3);
INSERT INTO `cart` VALUES (2, 1, 3, 3);
INSERT INTO `cart` VALUES (4, 3, 2, 2);
INSERT INTO `cart` VALUES (5, 3, 3, 3);
INSERT INTO `cart` VALUES (7, 3, 1, 1);

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `goods_classify_id` int(11) NULL DEFAULT NULL,
  `goods_saled_nums` int(11) NULL DEFAULT NULL,
  `goods_prices` double(11, 1) NULL DEFAULT 0.0,
  `goods_inventory` int(11) NULL DEFAULT 0 COMMENT '库存',
  `goods_picture_address` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '校服1', 1, 1, 89.9, 0, '/upload/test.jpg');
INSERT INTO `goods` VALUES (2, '校服2', 1, 1, 99.9, 0, '/upload/test.jpg');
INSERT INTO `goods` VALUES (3, '校服3', 1, 1, 88.8, 0, '/upload/test.jpg');
INSERT INTO `goods` VALUES (4, '校服4', 1, 1, 88.8, 0, '/upload/test.jpg');
INSERT INTO `goods` VALUES (5, '校服5', 1, 1, 88.8, 0, '/upload/test.jpg');

-- ----------------------------
-- Table structure for goodsclassify
-- ----------------------------
DROP TABLE IF EXISTS `goodsclassify`;
CREATE TABLE `goodsclassify`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `classify_name` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goodsclassify
-- ----------------------------
INSERT INTO `goodsclassify` VALUES (1, '校服');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
  `saler_id` int(11) NULL DEFAULT NULL COMMENT '分销的销售者id',
  `payment` decimal(10, 2) NULL DEFAULT NULL COMMENT '实付金额',
  `status` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '未付款' COMMENT '交易状态,待付款/已付款/已发货/未发货',
  `buyer_message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '买家留言',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, 3, NULL, 55.00, '未付款', '尽快');
INSERT INTO `order` VALUES (2, 3, NULL, 88.88, '未付款', '测试');

-- ----------------------------
-- Table structure for orderitem
-- ----------------------------
DROP TABLE IF EXISTS `orderitem`;
CREATE TABLE `orderitem`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '订单详情,主要关联订单中的商品',
  `order_id` int(11) NULL DEFAULT NULL COMMENT '订单id',
  `goods_id` int(11) NULL DEFAULT NULL COMMENT '商品id',
  `goods_nums` int(11) NULL DEFAULT NULL COMMENT '商品数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orderitem
-- ----------------------------
INSERT INTO `orderitem` VALUES (1, 1, 1, 2);
INSERT INTO `orderitem` VALUES (2, 1, 2, 3);
INSERT INTO `orderitem` VALUES (3, 2, 3, 3);
INSERT INTO `orderitem` VALUES (4, 2, 5, 5);

-- ----------------------------
-- Table structure for saler
-- ----------------------------
DROP TABLE IF EXISTS `saler`;
CREATE TABLE `saler`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户',
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '上一级',
  `branch_id` int(11) NULL DEFAULT NULL COMMENT '下一级',
  `pass` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '否' COMMENT '是否审核通过称为分销商',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `contact` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE,
  UNIQUE INDEX `openid_2`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'test', '15869190407');
INSERT INTO `user` VALUES (2, 'test2', '15869190407');
INSERT INTO `user` VALUES (3, 'oDi1V4zzhlmP_NkAyEA2-6oChnYw', '');

SET FOREIGN_KEY_CHECKS = 1;
