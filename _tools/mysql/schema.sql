CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     varchar(255) NOT NULL COMMENT 'ユーザー名',
    `email`     varchar(255) NOT NULL COMMENT 'メールアドレス',
    `twitter_id`     varchar(255) NOT NULL COMMENT 'twitterのid',
    `password`     varchar(255) NOT NULL COMMENT 'パスワード',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `event_participation`
(
    `user_id`       BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーの識別子',
    `event_id`       BIGINT UNSIGNED NOT NULL COMMENT 'イベントの識別子'
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
CREATE TABLE `event`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name`     varchar(255) NOT NULL COMMENT 'イベント名',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
