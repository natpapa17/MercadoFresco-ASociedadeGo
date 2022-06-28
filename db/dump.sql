USE `fresh_market`;

CREATE TABLE warehouses (
  id int not null auto_increment,
  warehouse_code varchar(25) not null unique ,
  address varchar(255) not null ,
  telephone varchar(15) not null ,
  minimum_capacity int not null ,
  minimum_temperature float not null ,
  primary key(id)
);

INSERT INTO `warehouses` (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES ('XPTO', 'Rua Brasil, 123', '(44) 99999-9999', 10, 5.4);