CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     varchar(255) NOT NULL COMMENT 'ユーザー名',
    `email`     varchar(255) NOT NULL COMMENT 'メールアドレス',
    `twitter_id`     varchar(255) NOT NULL COMMENT 'twitterのid',
    `password`     varchar(255) NOT NULL COMMENT 'パスワード',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
