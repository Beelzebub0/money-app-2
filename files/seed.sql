SET foreign_key_checks = 0;

INSERT INTO `user`
(
    `id`,
    `name`,
    `job`,
    `notes`
)
VALUES
(1, 'Teddy', 'magician', 'this is just a test');

INSERT INTO `categories`
(
    `id`,
    `name`,
    `description`
)
VALUES
(1, 'Food & Beverages', 'Something You Eat'),
(2, 'Electricity Bills', 'Something You Use Everyday');
SET foreign_key_checks = 1;