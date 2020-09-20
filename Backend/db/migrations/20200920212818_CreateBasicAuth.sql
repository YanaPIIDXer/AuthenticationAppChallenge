
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `basic_auth` (
  `id` INT NOT NULL,
  `email` VARCHAR(64) CHARACTER SET BINARY NOT NULL,
  `password` VARCHAR(32) CHARACTER SET BINARY NOT NULL
) ENGINE = InnoDB;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `basic_auth`;
