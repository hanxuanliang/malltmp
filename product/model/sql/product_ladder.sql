-- add 2021-03-11

-- ----------------------------
-- Table structure for pms_product_ladder
-- ----------------------------
DROP TABLE IF EXISTS `pms_product_ladder`;
CREATE TABLE `pms_product_ladder` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) DEFAULT NULL,
  `count` int(11) DEFAULT NULL COMMENT '满足的商品数量',
  `discount` decimal(10,2) DEFAULT NULL COMMENT '折扣',
  `price` decimal(10,2) DEFAULT NULL COMMENT '折后价格',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8 COMMENT='产品阶梯价格表(只针对同商品)';


-- ----------------------------
-- Records of pms_product_ladder
-- ----------------------------
INSERT INTO `pms_product_ladder` VALUES ('1', '7', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('2', '8', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('3', '9', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('4', '10', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('5', '11', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('6', '12', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('7', '13', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('8', '14', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('12', '18', '3', '0.70', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('14', '7', '4', '0.60', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('15', '7', '5', '0.50', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('18', '22', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('20', '24', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('38', '23', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('48', '31', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('50', '32', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('52', '33', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('59', '34', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('60', '30', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('64', '27', '2', '0.80', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('65', '27', '3', '0.75', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('66', '28', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('67', '29', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('68', '26', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('83', '36', '0', '0.00', '0.00');
INSERT INTO `pms_product_ladder` VALUES ('84', '35', '0', '0.00', '0.00');
