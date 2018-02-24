CREATE TABLE `moniple`.`schedules` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `url` VARCHAR(256) NOT NULL,
  `method` VARCHAR(6) NOT NULL,
  `interval` INT NOT NULL,
  `status` VARCHAR(8) NOT NULL,
  PRIMARY KEY (`id`));
