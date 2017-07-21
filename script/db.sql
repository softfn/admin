-- ----------------------------
--  Table structure for `sys_function`
-- ----------------------------
DROP TABLE IF EXISTS `sys_function`;
CREATE TABLE `sys_function` (
  `id` varchar(27) CHARACTER SET utf8 NOT NULL COMMENT '主键标识',
  `menu_id` varchar(27) CHARACTER SET utf8 NOT NULL COMMENT '菜单标识',
  `name` varchar(32) CHARACTER SET utf8 DEFAULT NULL COMMENT '功能名称',
  `code` varchar(64) CHARACTER SET utf8 DEFAULT NULL COMMENT '功能编码',
  `seq` tinyint(1) DEFAULT '1' COMMENT '排列次序',
  `created_by` varchar(27) CHARACTER SET utf8 NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) CHARACTER SET utf8 DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_code` (`code`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='菜单功能表';

-- ----------------------------
--  Table structure for `sys_grant`
-- ----------------------------
DROP TABLE IF EXISTS `sys_grant`;
CREATE TABLE `sys_grant` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `principal_type` tinyint(1) NOT NULL COMMENT '授权主体类型(1:角色 2:用户)',
  `principal_id` varchar(27) NOT NULL COMMENT '授权主体标识',
  `resource_type` tinyint(1) NOT NULL COMMENT '资源类型(1：菜单 2：功能)',
  `resource_id` varchar(27) NOT NULL COMMENT '资源标识',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_principalid_resourceid` (`principal_id`,`resource_id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='授权表';

-- ----------------------------
--  Table structure for `sys_menu`
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `parent_id` varchar(27) DEFAULT NULL COMMENT '上级菜单标识',
  `name` varchar(32) NOT NULL COMMENT '菜单名称',
  `url` varchar(500) DEFAULT NULL COMMENT '访问地址',
  `seq` tinyint(1) DEFAULT '1' COMMENT '排列次序',
  `visible` tinyint(1) DEFAULT '1' COMMENT '是否可见(1:可见，2：隐藏)',
  `tooltip` varchar(64) DEFAULT NULL COMMENT '提示信息',
  `style` varchar(500) DEFAULT NULL COMMENT '样式（JSON）',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_parentid_name` (`parent_id`,`name`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='菜单表';

-- ----------------------------
--  Table structure for `sys_org`
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `parent_id` varchar(27) DEFAULT NULL COMMENT '父结点标识',
  `virtual_code` varchar(64) DEFAULT NULL COMMENT '虚拟编码',
  `code` varchar(32) DEFAULT NULL COMMENT '组织标签编码',
  `name` varchar(128) NOT NULL COMMENT '机构名称',
  `full_name` varchar(512) DEFAULT NULL COMMENT '组织单元名称',
  `type` tinyint(1) DEFAULT NULL COMMENT '类型（1：集团 2：分部 3:权威机构专用）',
  `dep_chief` varchar(64) DEFAULT NULL COMMENT '部门主管',
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `tel` varchar(12) DEFAULT NULL COMMENT '办公电话',
  `address` varchar(512) DEFAULT NULL COMMENT '地址',
  `seq` smallint(6) DEFAULT '1' COMMENT '显示顺序',
  `del_flag` tinyint(1) DEFAULT '0' COMMENT '删除标志（0：否 1：是）',
  `del_time` datetime DEFAULT NULL COMMENT '删除时间',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `origin_id` varchar(64) DEFAULT NULL COMMENT '来源标识',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_sec_org_pk` (`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='机构表';

-- ----------------------------
--  Table structure for `sys_org_staff`
-- ----------------------------
DROP TABLE IF EXISTS `sys_org_staff`;
CREATE TABLE `sys_org_staff` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `staff_id` varchar(27) NOT NULL COMMENT '员工标识',
  `org_id` varchar(64) DEFAULT NULL COMMENT '部门标识',
  `blong_type` tinyint(1) DEFAULT '1' COMMENT '归属类型(1主 2从)',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_sec_org_staff_unique` (`staff_id`,`org_id`,`blong_type`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='组织和员工关系表';

-- ----------------------------
--  Table structure for `sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `type` tinyint(1) DEFAULT '2' COMMENT '角色类别（1：管理员角色  2：普通角色）',
  `del_flag` tinyint(1) DEFAULT '0' COMMENT '删除标志（1：是, 2：否 ）',
  `del_time` datetime DEFAULT NULL COMMENT '删除时间',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_name` (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
--  Table structure for `sys_role_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_user`;
CREATE TABLE `sys_role_user` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `role_id` varchar(27) NOT NULL COMMENT '角色标识',
  `user_id` varchar(27) NOT NULL COMMENT '用户标识',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_roleid_userid` (`role_id`,`user_id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

-- ----------------------------
--  Table structure for `sys_staff`
-- ----------------------------
DROP TABLE IF EXISTS `sys_staff`;
CREATE TABLE `sys_staff` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `ch_name` varchar(16) DEFAULT NULL COMMENT '中文名',
  `ch_py` varchar(32) DEFAULT NULL COMMENT '中文拼音',
  `en_name` varchar(32) DEFAULT NULL COMMENT '英文名',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `emp_no` varchar(128) DEFAULT NULL COMMENT '员工编号',
  `emp_status` tinyint(1) DEFAULT NULL COMMENT '员工状态',
  `position` varchar(32) DEFAULT NULL COMMENT '职位',
  `rank` varchar(32) DEFAULT NULL COMMENT '职等',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `id_card_no` varchar(32) DEFAULT NULL COMMENT '身份证号',
  `entry_time` datetime DEFAULT NULL COMMENT '入职日期',
  `email` varchar(200) DEFAULT NULL COMMENT '电子邮箱',
  `postal_address` varchar(256) DEFAULT NULL COMMENT '邮寄地址',
  `postal_code` varchar(32) DEFAULT NULL COMMENT '邮政编码',
  `home_address` varchar(500) DEFAULT NULL COMMENT '家庭住址',
  `mobile_no` varchar(16) DEFAULT NULL COMMENT '移动电话',
  `telephone` varchar(16) DEFAULT NULL COMMENT '办公电话',
  `sex` tinyint(1) DEFAULT NULL COMMENT '性别（1 男  2 女）',
  `is_enabled` tinyint(1) DEFAULT NULL COMMENT '是否有效（1：有效，2：无效 ）',
  `del_flag` tinyint(1) DEFAULT '2' COMMENT '删除标志（ 1：是，2：否）',
  `del_time` datetime DEFAULT NULL COMMENT '删除时间',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `origin_id` varchar(64) DEFAULT NULL COMMENT '来源标识',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_username` (`username`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='员工表';

-- ----------------------------
--  Table structure for `sys_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` varchar(27) NOT NULL COMMENT '主键标识',
  `name` varchar(32) DEFAULT NULL COMMENT '姓名',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(32) DEFAULT NULL COMMENT '用户密码',
  `login_ip` varchar(50) DEFAULT NULL COMMENT '当前登录的客户端ip地址',
  `login_time` datetime DEFAULT NULL COMMENT '当前登录时间',
  `last_login_ip` varchar(50) DEFAULT NULL COMMENT '上一次登录的客户端ip地址',
  `last_login_time` datetime DEFAULT NULL COMMENT '上一次登录的时间',
  `failed_count` int(11) DEFAULT NULL COMMENT '登录失败次数',
  `del_flag` tinyint(1) DEFAULT '0' COMMENT '删除标志（1：是, 2：否 ）',
  `del_time` datetime DEFAULT NULL COMMENT '删除时间',
  `is_enabled` tinyint(1) DEFAULT NULL COMMENT '是否有效(1: 有效，2：无效)',
  `login_type` tinyint(1) DEFAULT '1' COMMENT '登录方式（1：本地认证）',
  `description` varchar(500) DEFAULT NULL COMMENT '描述',
  `created_by` varchar(27) NOT NULL COMMENT '创建人标识',
  `creation_date` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(27) DEFAULT NULL COMMENT '修改人标识',
  `update_date` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_table_username` (`username`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
--  Records 
-- ----------------------------
