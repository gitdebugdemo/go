CREATE TABLE IF not  exists  `userinfo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '姓名',
  `add_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入学时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


insert into demo.userinfo (id, name, add_time) values (1, '新来的', '2019-12-30 15:24:27');
insert into demo.userinfo (id, name, add_time) values (2, '旧来的', '2019-12-30 15:24:41');
