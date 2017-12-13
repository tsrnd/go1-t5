drop table order_products;
drop table products;
drop table category;
drop table images;
drop table orders;
drop table payments;
drop table users;
create table users
(
    id serial primary key,
    username varchar(255) not null,
    password varchar(255) not null,
    email varchar(255) not null unique,
    gender varchar(20) not null,
    role varchar(50) not null,
    avatar varchar(255),
    phone varchar(15) not null,
    address varchar(255),
    created_at timestamp not null,
    updated_at timestamp,
    delete_at timestamp
);
create table payments
(
    id serial primary key,
    method varchar(255) not null,
    created_at timestamp not null  ,
    updated_at timestamp,
    delete_at timestamp
);
create table categories
(
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp not null  ,
    updated_at timestamp,
    delete_at timestamp
);
create table products
(
    id serial primary key,
    name varchar(255) not null,
    quantity integer not null,
    price money not null,
    category_id integer not null references categories(id),
    description text,
    created_at timestamp not null,
    updated_at timestamp,
    delete_at timestamp
);
create table images
(
    id serial primary key,
    name varchar(255) not null,
    url varchar(255) not null,
    product_id integer not null references products(id),
    created_at timestamp not null  ,
    updated_at timestamp,
    delete_at timestamp
);
create table orders
(
    id serial primary key,
    user_id integer not null REFERENCES users(id),
    total_money money not null,
    status varchar(20) not null,
    payment_id integer REFERENCES payments(id) ,
    created_at timestamp not null  ,
    updated_at timestamp,
    delete_at timestamp
);
create table order_products
(
    id serial primary key,
    order_id integer not null REFERENCES orders(id),
    product_id integer not null REFERENCES products(id),
    quantity integer not null,
    created_at timestamp not null  ,
    updated_at timestamp,
    delete_at timestamp
);
