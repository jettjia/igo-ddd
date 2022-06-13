CREATE TABLE `user`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`   varchar(45)      NOT NULL COMMENT 'User Passport',
    `password`   varchar(45)      NOT NULL COMMENT 'User Password',
    `nickname`   varchar(45)      NOT NULL COMMENT 'User Nickname',
    `create_at`  datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at`  datetime DEFAULT NULL COMMENT 'Updated Time',
    `deleted_at` datetime DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE `user_log`
(
    `id`      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20)          NOT NULL,
    `log`     varchar(64)         NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


