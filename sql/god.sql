/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : god

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001

Date: 2019-12-06 01:21:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for f_admins
-- ----------------------------
DROP TABLE IF EXISTS `f_admins`;
CREATE TABLE `f_admins` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL COMMENT '用户名',
  `email` varchar(25) NOT NULL COMMENT '邮箱',
  `mobile` char(11) NOT NULL COMMENT '手机',
  `password` char(60) NOT NULL COMMENT '密码',
  `avatar` varchar(120) NOT NULL DEFAULT '' COMMENT '头像',
  `state` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '0开启，1关闭，99删除',
  `group_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '组ID',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `login_at` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `login_ip` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后登陆IP',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_name` (`name`) USING BTREE,
  KEY `idx_email` (`email`) USING BTREE,
  KEY `idx_mobile` (`mobile`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_admins
-- ----------------------------
INSERT INTO `f_admins` VALUES ('1', 'dylan', 'email@qq.com', '', '$2a$10$gvNqEVrUF1A6VOExc4s.Vezc3Am8oxknNyfIhEOttytFZhkhysnOy', 'resource/upload/images/avatar/1202633773586321408_{size}.jpg', '0', '1', '2019-09-11 00:46:04', '2019-12-06 01:00:30', null, '2019-12-06 01:00:30', '2130706433');
INSERT INTO `f_admins` VALUES ('2', '产品经理88', '97072365@qq.com', '13750581124', '$2a$10$bMxRBbZA5JIxlpWpsJjDbes1XJPujpvmW653vfI30tCga655EE9qO', 'resource/upload/images/avatar/1185557874609229824_{size}.jpg', '0', '3', '2019-10-19 17:54:30', '2019-10-19 22:23:37', null, null, '0');
INSERT INTO `f_admins` VALUES ('3', '产品经理8855', '', '', '$2a$10$01FCgixysy7NuorQroRhE.9N6vTzf.U7gP715iVb4sVgnY5gArWQK', 'resource/upload/images/avatar/1185578390522957824_{size}.jpg', '0', '2', '2019-10-19 23:28:10', '2019-10-19 23:28:15', '2019-10-20 00:06:05', null, '0');
INSERT INTO `f_admins` VALUES ('4', 'imlab', '', '', '$2a$10$X4T/A.VUUbmTWkHu8wDAJuuuV5jKq8pP98zL/TvsUllOAl4HvnU9O', 'resource/upload/images/avatar/1195275787146629120_{size}.png', '0', '1', '2019-10-20 15:49:08', '2019-11-15 17:42:10', null, null, '0');
INSERT INTO `f_admins` VALUES ('5', 'hilolab', '504534678@qq.com', '', '$2a$10$uUpcAaNLSyUTNACIvgZ2COwFEwfxBIepGGW7msjM1xGx2OdMp5Npy', 'resource/upload/images/avatar/1185931172044083200_{size}.png', '0', '3', '2019-10-20 22:50:00', '2019-10-21 16:52:16', null, null, '0');
INSERT INTO `f_admins` VALUES ('6', '管理员123', 'test@qq.com', '13650857406', '$2a$10$PlZmhXXDX69X15mg1dhHbe.LCXPOz6XxThCoJLP5qYfSKC.KLEQBq', 'resource/upload/images/avatar/1186276839480365056_{size}.jpg', '0', '3', '2019-10-21 21:43:05', '2019-10-21 21:44:00', '2019-10-21 21:53:03', null, '0');

-- ----------------------------
-- Table structure for f_admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `f_admin_roles`;
CREATE TABLE `f_admin_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_admin_role` (`admin_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_admin_roles
-- ----------------------------
INSERT INTO `f_admin_roles` VALUES ('6', '1', '2');
INSERT INTO `f_admin_roles` VALUES ('8', '1', '3');
INSERT INTO `f_admin_roles` VALUES ('3', '2', '2');
INSERT INTO `f_admin_roles` VALUES ('4', '2', '3');

-- ----------------------------
-- Table structure for f_groups
-- ----------------------------
DROP TABLE IF EXISTS `f_groups`;
CREATE TABLE `f_groups` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL COMMENT '角色名称',
  `state` tinyint(2) unsigned NOT NULL COMMENT '状态，0开启，1关闭，99删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_groups
-- ----------------------------
INSERT INTO `f_groups` VALUES ('1', 'CEO', '0', '2019-10-20 15:28:48', '2019-11-03 10:10:46', null);
INSERT INTO `f_groups` VALUES ('2', 'CFO', '1', '2019-10-20 15:29:20', '2019-10-20 15:29:29', '2019-10-20 15:29:46');
INSERT INTO `f_groups` VALUES ('3', '产品经理', '0', '2019-10-20 15:49:55', '2019-10-20 15:50:21', null);

-- ----------------------------
-- Table structure for f_privileges
-- ----------------------------
DROP TABLE IF EXISTS `f_privileges`;
CREATE TABLE `f_privileges` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `root` varchar(255) NOT NULL DEFAULT '' COMMENT '父ID集合，以“/”分割',
  `icon` varchar(25) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `uri_rule` varchar(55) NOT NULL DEFAULT '' COMMENT '访问节点规则',
  `is_menu` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否是菜单：0否，1是',
  `sort_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_root` (`root`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_privileges
-- ----------------------------
INSERT INTO `f_privileges` VALUES ('1', '0', '后台配置', '', 'fa-user', '', '1', '0', '2019-11-02 22:26:04', '2019-11-04 11:26:06', null);
INSERT INTO `f_privileges` VALUES ('2', '1', '管理员', '1/', '', '/admin/accounts', '1', '0', '2019-11-03 09:56:18', '2019-12-01 22:14:37', null);
INSERT INTO `f_privileges` VALUES ('3', '2', '添加管理员', '1/2', '', '/admin/account/add', '0', '0', '2019-11-03 10:23:15', '2019-11-04 16:28:06', null);
INSERT INTO `f_privileges` VALUES ('4', '2', '更新管理员', '1/2', '', '/admin/account/update', '0', '0', '2019-11-03 10:23:39', '2019-11-04 16:28:13', null);
INSERT INTO `f_privileges` VALUES ('5', '2', '删除管理员', '1/2', '', '/admin/account/delete', '0', '0', '2019-11-03 10:23:51', '2019-11-03 22:28:52', null);
INSERT INTO `f_privileges` VALUES ('6', '1', '分组管理', '1/', '', '/admin/groups', '1', '0', '2019-11-03 10:24:16', '2019-11-04 16:27:57', null);
INSERT INTO `f_privileges` VALUES ('7', '6', '添加分组', '1/6', '', '/admin/group/add', '0', '0', '2019-11-03 10:24:29', '2019-11-03 22:31:14', null);
INSERT INTO `f_privileges` VALUES ('8', '6', '更新分组', '1/6', '', '/admin/group/update', '0', '0', '2019-11-03 10:24:47', '2019-11-03 10:24:47', null);
INSERT INTO `f_privileges` VALUES ('9', '6', '删除分组', '1/6', '', '/admin/group/delete', '0', '0', '2019-11-03 10:24:58', '2019-11-03 10:24:58', null);
INSERT INTO `f_privileges` VALUES ('15', '0', '系统日志', '', '', '', '0', '0', '2019-11-04 11:08:31', '2019-11-04 11:08:31', null);
INSERT INTO `f_privileges` VALUES ('16', '15', '管理员登录日志', '15/', '', '', '0', '0', '2019-11-04 11:09:12', '2019-11-04 11:21:37', null);
INSERT INTO `f_privileges` VALUES ('17', '15', '管理员操作日志', '15/', '', '', '0', '0', '2019-11-04 11:09:40', '2019-11-04 11:22:40', null);
INSERT INTO `f_privileges` VALUES ('18', '16', '删除日志', '15/16', '', '', '0', '0', '2019-11-04 11:10:46', '2019-11-04 11:22:49', null);
INSERT INTO `f_privileges` VALUES ('19', '1', '角色管理', '1/', '', '/admin/roles', '1', '0', '2019-11-04 16:30:36', '2019-11-04 16:30:36', null);
INSERT INTO `f_privileges` VALUES ('20', '19', '添加角色', '1/19', '', '/admin/role/add', '0', '0', '2019-11-04 16:31:31', '2019-11-04 16:31:31', null);
INSERT INTO `f_privileges` VALUES ('21', '19', '更新角色', '1/19', '', '/admin/role/update', '0', '0', '2019-11-04 16:31:57', '2019-11-04 16:31:57', null);
INSERT INTO `f_privileges` VALUES ('22', '19', '删除角色', '1/19', '', '/admin/role/delete', '0', '0', '2019-11-04 16:32:20', '2019-11-04 16:32:20', null);
INSERT INTO `f_privileges` VALUES ('23', '1', '权限管理', '1/', '', '/admin/privileges', '1', '0', '2019-11-04 16:33:29', '2019-11-04 16:33:29', null);
INSERT INTO `f_privileges` VALUES ('24', '23', '添加权限', '1/23', '', '/admin/privilege/add', '0', '0', '2019-11-04 16:33:54', '2019-11-04 16:33:54', null);
INSERT INTO `f_privileges` VALUES ('25', '23', '更新权限', '1/23', '', '/admin/privilege/update', '0', '0', '2019-11-04 16:34:17', '2019-11-04 16:34:17', null);
INSERT INTO `f_privileges` VALUES ('26', '23', '删除权限', '1/23', '', '/admin/privilege/delete', '0', '0', '2019-11-04 16:34:39', '2019-11-04 16:34:39', null);
INSERT INTO `f_privileges` VALUES ('27', '1', '个人信息', '1/', '', '/admin/profile', '1', '0', '2019-11-04 16:35:28', '2019-11-04 16:35:40', null);
INSERT INTO `f_privileges` VALUES ('28', '27', '更新个人信息', '1/27', '', '/admin/profile/update', '0', '0', '2019-11-04 16:36:05', '2019-11-04 16:36:05', null);
INSERT INTO `f_privileges` VALUES ('29', '19', '关联权限', '1/19', '', '/admin/role/associate_privilege', '0', '0', '2019-11-05 18:40:02', '2019-11-05 18:40:02', null);
INSERT INTO `f_privileges` VALUES ('30', '23', '查看权限管理的角色', '1/23', '', '/admin/privilege/role_ids', '0', '0', '2019-11-05 18:40:39', '2019-12-01 20:32:29', null);
INSERT INTO `f_privileges` VALUES ('31', '19', '查看角色管理的权限', '1/19', '', '/admin/role/privileges', '0', '0', '2019-12-01 20:33:29', '2019-12-01 20:33:29', null);
INSERT INTO `f_privileges` VALUES ('35', '2', '查看管理员角色', '1/2', '', '/admin/account/role_ids', '0', '0', '2019-12-03 17:30:10', '2019-12-03 17:30:10', null);
INSERT INTO `f_privileges` VALUES ('36', '2', '关联管理员角色', '1/2', '', '/admin/account/associate_role', '0', '0', '2019-12-04 16:20:42', '2019-12-04 16:20:42', null);

-- ----------------------------
-- Table structure for f_roles
-- ----------------------------
DROP TABLE IF EXISTS `f_roles`;
CREATE TABLE `f_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL COMMENT '角色名称',
  `state` tinyint(2) unsigned NOT NULL COMMENT '状态，0开启，1关闭，99删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_roles
-- ----------------------------
INSERT INTO `f_roles` VALUES ('2', '管理员', '0', '2019-10-13 22:20:13', '2019-10-13 22:20:13', null);
INSERT INTO `f_roles` VALUES ('3', '产品经理', '0', '2019-10-13 23:10:06', '2019-10-14 22:19:01', null);
INSERT INTO `f_roles` VALUES ('4', '开发者', '1', '2019-10-16 16:45:09', '2019-10-16 16:45:09', '2019-12-01 14:42:35');

-- ----------------------------
-- Table structure for f_role_privileges
-- ----------------------------
DROP TABLE IF EXISTS `f_role_privileges`;
CREATE TABLE `f_role_privileges` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `privilege_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_role_privilege` (`role_id`,`privilege_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of f_role_privileges
-- ----------------------------
INSERT INTO `f_role_privileges` VALUES ('75', '2', '1');
INSERT INTO `f_role_privileges` VALUES ('15', '2', '2');
INSERT INTO `f_role_privileges` VALUES ('60', '2', '3');
INSERT INTO `f_role_privileges` VALUES ('59', '2', '4');
INSERT INTO `f_role_privileges` VALUES ('58', '2', '5');
INSERT INTO `f_role_privileges` VALUES ('61', '2', '6');
INSERT INTO `f_role_privileges` VALUES ('62', '2', '7');
INSERT INTO `f_role_privileges` VALUES ('63', '2', '8');
INSERT INTO `f_role_privileges` VALUES ('65', '2', '9');
INSERT INTO `f_role_privileges` VALUES ('42', '2', '16');
INSERT INTO `f_role_privileges` VALUES ('43', '2', '18');
INSERT INTO `f_role_privileges` VALUES ('50', '2', '19');
INSERT INTO `f_role_privileges` VALUES ('73', '2', '20');
INSERT INTO `f_role_privileges` VALUES ('70', '2', '21');
INSERT INTO `f_role_privileges` VALUES ('71', '2', '22');
INSERT INTO `f_role_privileges` VALUES ('28', '2', '23');
INSERT INTO `f_role_privileges` VALUES ('66', '2', '24');
INSERT INTO `f_role_privileges` VALUES ('67', '2', '25');
INSERT INTO `f_role_privileges` VALUES ('68', '2', '26');
INSERT INTO `f_role_privileges` VALUES ('48', '2', '27');
INSERT INTO `f_role_privileges` VALUES ('76', '2', '28');
INSERT INTO `f_role_privileges` VALUES ('74', '2', '29');
INSERT INTO `f_role_privileges` VALUES ('49', '2', '30');
INSERT INTO `f_role_privileges` VALUES ('72', '2', '31');
INSERT INTO `f_role_privileges` VALUES ('57', '2', '35');
INSERT INTO `f_role_privileges` VALUES ('83', '2', '36');
INSERT INTO `f_role_privileges` VALUES ('11', '3', '1');
INSERT INTO `f_role_privileges` VALUES ('3', '3', '2');
INSERT INTO `f_role_privileges` VALUES ('4', '3', '3');
INSERT INTO `f_role_privileges` VALUES ('5', '3', '4');
INSERT INTO `f_role_privileges` VALUES ('6', '3', '5');
INSERT INTO `f_role_privileges` VALUES ('10', '3', '6');
INSERT INTO `f_role_privileges` VALUES ('64', '3', '9');
