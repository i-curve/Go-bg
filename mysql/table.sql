
-- name="diary"
drop table diary;
create database diary;
use diary;
set foreign_key_checks=0;

create table `user` (
    `id` int(10) unsigned not null auto_increment,
    `username` varchar(50) default '' comment '账号',
    `password` varchar(50) default '' comment '密码',
    `nickname` varchar(50) default '' comment '昵称',
    primary key(`id`)
)engine=InnoDB auto_increment=2 default charset=utf8;
create table `diary` (
    `id` int(10) unsigned not null auto_increment,
    `created_on` int(10) unsigned not null comment '创建时间',
    `modified_on` int(10) unsigned comment '修改时间',
    `deleted_on` int(10) unsigned comment '删除时间',
    `username` varchar(50) comment "用户名",
    `title` int(100) comment '标题',
    `text` varchar(255) comment "文本内容",
    primary key(`id`)
)engine=InnoDB auto_increment=1 DEFAULT charset=utf8;