--
-- Character Sets and Collations: utf8_general_ci (default collation for utf8 charset)
-- String comparisons and sorting _ci (case insensitive) _cs, _bin, ...
--
CREATE TABLE `incoming` (
 `id` bigint(20) unsigned NOT NULL COMMENT 'Tracking ID',
 `property` varchar(32) COLLATE utf8_general_ci NOT NULL COMMENT 'Property name (human readable, a-z).',
 `property_section` int(11) unsigned NOT NULL COMMENT 'Property Section ID.',
 `property_id` int(11) unsigned NOT NULL COMMENT 'Property Item ID.',
 `remote_ip` varchar(255) COLLATE utf8_general_ci NOT NULL COMMENT 'Remote IP from user making request.',
 `stamp` datetime NOT NULL COMMENT 'Timestamp of request.',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='Incoming stats log, writes only.';

CREATE TABLE `incoming_proc` LIKE `incoming`;
--
-- RENAME TABLE incoming TO incoming_old, incoming_proc TO incoming, incoming_old TO incoming_proc;
--
