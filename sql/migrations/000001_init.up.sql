create table IF NOT EXISTS product_types(
    id varchar(40) not null,
    name varchar(100) NOT NULL,
    description varchar(200) NOT NULL,
    created_at timestamp,
    primary key (id)
);

create table if not EXISTS products(
    id varchar(40) NOT NULL , 
    name varchar(100) NOT NULL,
    description varchar(200) NOT NULL,
    product_type_id varchar(40),
    created_at timestamp,
    primary key (id),
    foreign key (product_type_id) references product_types(id)
);

