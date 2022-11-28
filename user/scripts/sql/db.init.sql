create database user_center;
use user_center;

CREATE TABLE `user`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT COMMENT 'user id',
    `nickname`   varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'User Nickname',
    `avatar`     varchar(255) COLLATE utf8mb4_general_ci          DEFAULT NULL COMMENT 'User Avatar',
    `sex`        tinyint unsigned                                 DEFAULT '0' COMMENT 'User sex: 0-unknown; 1-male; 2-female',
    `password`   varchar(255) COLLATE utf8mb4_general_ci          DEFAULT NULL COMMENT 'password digest',
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `created_by` bigint unsigned                         NOT NULL COMMENT 'creator',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


CREATE TABLE `user_extend`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned                                  DEFAULT NULL COMMENT 'user id',
    `key`        varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Extend config',
    `value`      json                                             DEFAULT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `created_by` bigint unsigned                         NOT NULL COMMENT 'creator',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


CREATE TABLE `user_history`
(
    `id`         bigint unsigned  NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned  NOT NULL COMMENT 'User id',
    `created_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create time',
    `created_by` bigint unsigned  NOT NULL COMMENT 'Creator',
    `content`    json             NOT NULL COMMENT 'User History',
    `action`     tinyint unsigned NOT NULL COMMENT '0 : create\n1 : login\n2 : logout\n3 : deleted\n4 : edit user\n',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;