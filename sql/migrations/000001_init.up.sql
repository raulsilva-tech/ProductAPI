create table products(
    id varchar(40) NOT NULL PRIMARY KEY, 
    name varchar(100) NOT NULL,
    description varchar(200) NOT NULL,
    product_type_id varchar(40),
    created_at datetime,
    foreign key (product_type_id) references product_types(id)
);

create table product_types(
    id varchar(40) not null primary key,
    name varchar(100) NOT NULL,
    description varchar(200) NOT NULL,
    created_at datetime
);