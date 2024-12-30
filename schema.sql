CREATE TABLE ventas (
    id serial NOT NULL PRIMARY KEY,
    nombre varchar(50) NOT NULL,
    cantidad_sandwich int NOT NULL,
    total float NOT NULL
);