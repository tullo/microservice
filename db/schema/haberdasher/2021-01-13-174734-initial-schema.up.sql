CREATE TABLE `hat` (
 `id` bigint(20) unsigned NOT NULL COMMENT 'Tracking ID',
 `size` int(11) unsigned NOT NULL COMMENT 'The size of a hat in centimeters.',
 `color` varchar(32) COLLATE utf8_general_ci NOT NULL COMMENT 'The name of a hat.',
 `name` varchar(32) COLLATE utf8_general_ci NOT NULL COMMENT 'The color of a hat.',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='Hat is a piece of headwear made by a Haberdasher.';
