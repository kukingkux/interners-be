CREATE TABLE IF NOT EXISTS users (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `firstName` VARCHAR(255) NOT NULL,
    `lastName` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phoneNumber` VARCHAR(15),
    `zipCode` VARCHAR(9),
    `city` VARCHAR(255),
    `address` TEXT,
    `cv` TEXT,
    `profilePicture` TEXT,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`email`)
);
