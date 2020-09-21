
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `token` (
    `id` INT NOT NULL,
    `token` VARCHAR(64) CHARACTER SET BINARY NOT NULL,
    `period` DATETIME NOT NULL,
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `token`;
