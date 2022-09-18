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
    `admin_user_id`       BIGINT UNSIGNED NOT NULL COMMENT '管理者ユーザーの識別子',
    `name`     varchar(255) NOT NULL COMMENT 'イベント名',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `event_group`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `event_id`       BIGINT UNSIGNED NOT NULL COMMENT 'イベントの識別子',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='イベントグループ';
CREATE TABLE `event_group_participation`
(
  `group_id`       BIGINT UNSIGNED NOT NULL COMMENT 'イベントの識別子',
  `user_id`       BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーの識別子'
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='イベントグループ';

CREATE TABLE `question`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `content`     varchar(255) NOT NULL COMMENT '質問内容',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `answer`
(
    `question_id`       BIGINT UNSIGNED NOT NULL COMMENT 'イベントの識別子',
    `content`     varchar(255) NOT NULL COMMENT '回答内容',
    `number`     varchar(255) NOT NULL COMMENT '回答番号'
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `user_answer`
(
    `question_id`       BIGINT UNSIGNED NOT NULL COMMENT 'questionの識別子',
    `user_id`       BIGINT UNSIGNED NOT NULL COMMENT 'ユーザーの識別子',
    `number`     varchar(255) NOT NULL COMMENT '回答番号'
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
