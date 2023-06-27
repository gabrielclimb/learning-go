CREATE TABLE IF NOT EXISTS products
(
    ID          serial PRIMARY KEY,
    NAME        varchar(255) NOT NULL,
    DESCRIPTION varchar(255) NOT NULL,
    PRICE       decimal      NOT NULL,
    QUANTITY    integer      NOT NULL);