
-- name="diary"
drop table diary;
create database diary;
use diary;
set foreign_key_checks=0;

create table `user` (
    `id` int(10) unsigned not null auto_increment,
    `username` varchar(50) default '' comment '账号',
    `password` varchar(50) default '' comment '密码',
    `email` varchar(255) default '' comment '邮箱',
    `role` int(5) default 0 comment "用户角色",
    `created_on` int(10) unsigned not null comment '创建时间',
    `modified_on` int(10) unsigned not null comment '修改时间',
    primary key(`id`)
)engine=InnoDB auto_increment=2 default charset=utf8;

-- create table `diary` (
--     `id` int(10) unsigned not null auto_increment,
--     `created_on` int(10) unsigned not null comment '创建时间',
--     `modified_on` int(10) unsigned comment '修改时间',
--     `deleted_on` int(10) unsigned comment '删除时间',
--     `username` varchar(50) comment "用户名",
--     `title` int(100) comment '标题',
--     `text` varchar(255) comment "文本内容",
--     primary key(`id`)
-- )engine=InnoDB auto_increment=1 DEFAULT charset=utf8;

CREATE TABLE `lrs_audit_rule_package`(
  `id` BIGINT(20) AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
  `package_code` varchar(6)  NOT NULL COMMENT '规则包',
  `package_type` varchar(2)  NOT NULL COMMENT '规则包类型',
  `package_desc` varchar(100)  COMMENT '描述',
  `create_time` datetime DEFAULT NULL COMMENT '创始时间',
  `modified_time` datetime DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='审核规则包';