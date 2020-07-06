CREATE TABLE `user_similarity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT 'userid i',
  `s_user_id` bigint(20) DEFAULT NULL COMMENT 'userid v',
  `similarity` double DEFAULT NULL COMMENT 'similarity',
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;