CREATE TABLE IF NOT EXISTS products
(
    ID          serial PRIMARY KEY,
    NAME        varchar(255) NOT NULL,
    DESCRIPTION varchar(255) NOT NULL,
    PRICE       decimal      NOT NULL,
    AMOUNT    integer      NOT NULL);

INSERT INTO PRODUCTS (NAME, DESCRIPTION, PRICE, AMOUNT)
    VALUES ('Camiseta', 'Azul clara', 39, 3),
           ('tenis', 'Comfortable', 150, 1),
           ('Fone Bluetooth', 'Fone sem fio', 109, 10),
           ('Notebook', 'Fast PC', 1900, 2);