create table customers_addresses(
  address_id varchar(40) primary key, 
  customer_id varchar(40), 
  type varchar(20),
  address1 varchar(120),
  address2 varchar(120), 
  city varchar(50), 
  state varchar(2), 
  postal_code varchar(9), 
  country varchar(3), 
  validated BOOLEAN, 
  deleted_at datetime,
  constraint customer_address unique (customer_id, address1)
);
