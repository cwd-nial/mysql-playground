SET NAMES utf8;

CREATE DATABASE `playground`;
GRANT ALL PRIVILEGES ON `playground`.* TO 'neo'@'%';

USE `playground`;

CREATE TABLE `message_history_log`
(
    `id`           int(11)                                       NOT NULL AUTO_INCREMENT,
    `message_type` enum ('INFO','WARNING','CRITICAL')            NOT NULL,
    `space_type`   enum ('REALITY','THE_CONSTRUCT','THE_MATRIX') NOT NULL,
    `receiver_id`  bigint(21)                                    NOT NULL,
    `log_data`     text                                                   DEFAULT NULL,
    `create_time`  datetime                                      NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`),
    KEY `message_type_idx` (`message_type`),
    KEY `receiver_idx` (`receiver_id`),
    KEY `create_time_idx` (`create_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('INFO', 'REALITY', 6641334, 'follow the white rabbit_', '2023-10-11 11:35:01');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('INFO', 'REALITY', 6641334, 'follow the white rabbit_', '2023-10-11 11:35:01');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('INFO', 'THE_CONSTRUCT', 6641334, 'follow the white rabbit_', '2023-10-11 11:35:01');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('WARNING', 'THE_MATRIX', 6641334, 'was that another black cat?', '2023-11-01 14:25:01');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('WARNING', 'THE_MATRIX', 6641334, 'was that another black cat?', '2023-11-01 14:25:02');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('WARNING', 'THE_MATRIX', 6641334, 'was that another black cat?', '2023-11-01 14:25:03');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('WARNING', 'THE_MATRIX', 6641334, 'was that another black cat?', '2023-11-01 14:25:03');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('INFO', 'REALITY', 511199, 'everything tastes like chicken...', '2023-11-04 22:01:56');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('INFO', 'REALITY', 511199, 'everything tastes like chicken...', '2023-11-04 22:01:58');
INSERT INTO playground.message_history_log (message_type, space_type, receiver_id, log_data, create_time)
VALUES ('CRITICAL', 'REALITY', 6641334, 'last transmission... signal lost_', '2023-11-15 09:28:50');
