/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : casbin

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001

Date: 2019-01-22 17:21:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) DEFAULT NULL,
  `sub` varchar(100) DEFAULT NULL,-- v0
  `obj` varchar(100) DEFAULT NULL,-- v1
  `act` varchar(100) DEFAULT NULL,-- v2
  `ext` varchar(100) DEFAULT NULL,-- v3
  `name` varchar(100) DEFAULT NULL,-- v4
  `des` varchar(100) DEFAULT NULL,-- v5
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_act` (`act`),
  KEY `IDX_casbin_rule_ext` (`ext`),
  KEY `IDX_casbin_rule_name` (`name`),
  KEY `IDX_casbin_rule_des` (`des`),
  KEY `IDX_casbin_rule_sub` (`sub`),
  KEY `IDX_casbin_rule_obj` (`obj`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_casbin_rule
-- ----------------------------
INSERT INTO `sys_casbin_rule` VALUES ('65', 'p', '90', '/*', 'ANY', '.*', '超级用户', '', null);
INSERT INTO `sys_casbin_rule` VALUES ('66', 'g', '90', 'admin', '', '', '', '', null);
INSERT INTO `sys_casbin_rule` VALUES ('67', 'g', '90', 'demo', '', '', '', '', null);
INSERT INTO `sys_casbin_rule` VALUES ('68', 'p', 'admin', '/admin*', 'GET|POST|DELETE|PUT', '.*', '角色管理', '', null);
INSERT INTO `sys_casbin_rule` VALUES ('69', 'p', 'demo', '/demo*', 'GET|POST|DELETE|PUT', '.*', 'demo角色12', '', null);
INSERT INTO `sys_casbin_rule` VALUES ('71', 'p', 't1', '/*', 'PUT|DELETE|GET|POST', '.*', '测试1', '', null);

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo` (
  `pid` int(10) NOT NULL AUTO_INCREMENT,
  `product_code` varchar(255) DEFAULT '',
  `product_name` varchar(255) DEFAULT '',
  `number` int(11) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of demo
-- ----------------------------
INSERT INTO `demo` VALUES ('1', 'PD-001', 'asc', '2', '2019-01-08 10:30:33');
INSERT INTO `demo` VALUES ('2', 'PD-002', 'asc', '2', '2019-01-08 10:38:21');
INSERT INTO `demo` VALUES ('3', 'PD-003', 'asc', '2', '2019-01-08 10:38:51');
INSERT INTO `demo` VALUES ('4', 'PD-005', 'asc', '2', '2019-01-08 10:39:33');
INSERT INTO `demo` VALUES ('5', 'PD-006', 'asc', '2', '2019-01-08 10:44:21');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `url` varchar(64) DEFAULT '',
  `modular` varchar(64) DEFAULT '' COMMENT '模块名',
  `component` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '',
  `icon` varchar(64) DEFAULT '',
  `keep_alive` varchar(64) DEFAULT '',
  `require_auth` varchar(64) DEFAULT '',
  `parent_id` int(10) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT '1',
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('1', '/', null, '', null, '所有', null, null, null, null, '1', null, null);
INSERT INTO `sys_menu` VALUES ('2', '/', '/home', '', 'Home', '权限管理', 'fa fa-cog', null, null, '1', '1', null, null);
INSERT INTO `sys_menu` VALUES ('3', '/sys_user', '/a', 'admin', 'User', '用户管理', 'fa fa-sys_user', null, null, '2', '1', null, null);
INSERT INTO `sys_menu` VALUES ('4', '/role', '/a', 'admin', 'Role', '角色管理', 'fa fa-sys_user-secret', '', '', '2', '1', null, null);
INSERT INTO `sys_menu` VALUES ('5', '/sys_menu', '/a', 'admin', 'Menu', '菜单管理', 'fa fa-quora', '', '', '2', '1', null, null);
INSERT INTO `sys_menu` VALUES ('14', '/sys_user', '/', '7', '7', 'as', '7', '', '', '3', '7', '2018-07-25 13:59:04', null);
INSERT INTO `sys_menu` VALUES ('15', '/sys_user', '/', '8', '8', '大萨达撒の21321', '8', '', '', '3', '8', '2018-07-25 13:59:57', null);
INSERT INTO `sys_menu` VALUES ('16', '/sys_menu', '/', '9', '9', 'eqwewqedsads', '9', '', '', '5', '9', '2018-07-25 14:00:27', null);
INSERT INTO `sys_menu` VALUES ('17', '/sys_menu', '/', '1', '1', 'kjhjhgjghjgh炬华科技好看', '1', '', '', '5', '1', '2018-07-25 14:14:05', null);
INSERT INTO `sys_menu` VALUES ('18', '/sys_menu', '/', '计划国际化', '1', 'ss', '1', '', '', '5', '1', '2018-07-25 14:14:13', null);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `rid` int(10) NOT NULL COMMENT '角色id',
  `mid` int(10) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  KEY `m_r_id` (`mid`),
  KEY `r_m_id` (`rid`),
  CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`mid`) REFERENCES `sys_menu` (`id`),
  CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`rid`) REFERENCES `sys_casbin_rule` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('29', '68', '2');
INSERT INTO `sys_role_menu` VALUES ('30', '68', '3');
INSERT INTO `sys_role_menu` VALUES ('31', '68', '4');
INSERT INTO `sys_role_menu` VALUES ('32', '68', '5');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
	`token` varchar(255) DEFAULT '',
  `enabled` tinyint(1) DEFAULT '0' COMMENT '0：启用用户 1：禁用用户',
  `appid` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('76', 'yhm1', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', '', '', '', '', '2019-01-14 11:54:11', null);
INSERT INTO `sys_user` VALUES ('90', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', '超级用户', '', '', '', '2019-01-16 10:42:34', null);
INSERT INTO `sys_user` VALUES ('92', 'yhm2', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '', 'yy', '3213123', 'qq.com', '', '2019-01-18 10:08:12', null);
INSERT INTO `sys_user` VALUES ('104', '1-01', '1234', '0', '', 'mhjmhghf', 'nnvbn432423', '98089089', '', '2019-01-18 11:58:53', null);
INSERT INTO `sys_user` VALUES ('105', '22-201234', '123456', '1', '', 'fsdvcdxcvx', 'ffffffffffffffffffffff', '.,.mn,nm,hgntfrgffffffffffff', '', '2019-01-18 12:54:54', null);
INSERT INTO `sys_user` VALUES ('106', '3-12', '32qewqewqe', '0', '', '3ewqdd343444', '', 'sadsad1fs,..', '', '2019-01-18 12:56:31', null);
INSERT INTO `sys_user` VALUES ('108', '4', '6Xf9dl8Cv7OVFdc45vR7ig==', '1', '', '', '', '', '', '2019-01-18 13:34:56', null);
