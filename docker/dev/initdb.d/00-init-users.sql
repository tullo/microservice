CREATE DATABASE IF NOT EXISTS `haberdasher`;
CREATE USER 'haberdasher'@'%' IDENTIFIED BY 'haberdasher';
GRANT ALL ON `haberdasher`.* TO 'haberdasher'@'%';

CREATE DATABASE IF NOT EXISTS `stats`;
CREATE USER 'stats'@'%' IDENTIFIED BY 'stats';
GRANT ALL ON `stats`.* TO 'stats'@'%';

FLUSH PRIVILEGES;
