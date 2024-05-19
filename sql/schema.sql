CREATE SCHEMA IF NOT EXISTS `um_help` DEFAULT CHARACTER SET utf8;
USE `um_help`;

DROP TABLE IF EXISTS `um_help`.`tab_transaction`;
DROP TABLE IF EXISTS `um_help`.`tab_wallet`;
DROP TABLE IF EXISTS `um_help`.`tab_currency`;
DROP TABLE IF EXISTS `um_help`.`tab_user`;

CREATE TABLE IF NOT EXISTS `um_help`.`tab_user` (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `public_id` VARCHAR(36) NOT NULL UNIQUE DEFAULT (UUID()), -- TODO: Criar um IDX
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `document_number` VARCHAR(14) NOT NULL UNIQUE, -- TODO: Criar um IDX
  `password` VARCHAR(64) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `um_help`.`tab_currency` (
  `currency_id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(4) NOT NULL UNIQUE, -- TODO: Criar um IDX
  `symbol` VARCHAR(4) NOT NULL UNIQUE, -- TODO: Criar um IDX
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`currency_id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `um_help`.`tab_wallet` (
  `wallet_id` BIGINT NOT NULL AUTO_INCREMENT,
  `owner_id` BIGINT NOT NULL,
  `alias` VARCHAR(45) NOT NULL,
  `currency_id` BIGINT NOT NULL,
  `balance` BIGINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`wallet_id`),
  FOREIGN KEY (`owner_id`) REFERENCES `um_help`.`tab_user`(`user_id`),
  FOREIGN KEY (`currency_id`) REFERENCES `um_help`.`tab_currency`(`currency_id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `um_help`.`tab_transaction` (
  `transaction_id` BIGINT NOT NULL AUTO_INCREMENT,
  `origin_id` BIGINT NOT NULL,
  `destination_id` BIGINT NOT NULL,
  `transaction_value` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  FOREIGN KEY (`origin_id`) REFERENCES `um_help`.`tab_wallet`(`wallet_id`),
  FOREIGN KEY (`destination_id`) REFERENCES `um_help`.`tab_wallet`(`wallet_id`))
ENGINE = InnoDB;

INSERT INTO `um_help`.`tab_currency` (`code`, `symbol`) VALUES ('BRL', 'R$');
INSERT INTO `um_help`.`tab_currency` (`code`, `symbol`) VALUES ('USD', '$');
INSERT INTO `um_help`.`tab_currency` (`code`, `symbol`) VALUES ('EUR', '€');