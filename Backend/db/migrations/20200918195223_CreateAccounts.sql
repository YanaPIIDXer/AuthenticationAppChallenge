-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `accounts` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `email` VARCHAR(64) NOT NULL,
  `password` VARCHAR(32) NOT NULL
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `accounts`;