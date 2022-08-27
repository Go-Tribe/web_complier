CREATE TABLE `share` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `gid` varchar(30) NOT NULL COMMENT '唯一 id',
  `code` varchar(2000) NOT NULL DEFAULT '' COMMENT '代码',
  `type` varchar(10) NOT NULL DEFAULT '' COMMENT '语言',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `gid` (`gid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分享表';