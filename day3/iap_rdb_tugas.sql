/* Create Table: customers (id, customer_name) */
create table if not exists customers (
	id int not null,
	customer_name char(50) not null,
	primary key (id)
);

/*Create Table: products (id, name)*/
create table if not exists products (
	id int not null,
	name char(50) not null,
	primary key (id)
);

/*Create Table: 
 orders (order_id, customer_id, product_id, order_date, total)*/
create table if not exists Orders (
	order_id int not null,
	customer_id int not null,
	product_id int not null,
	order_date date not null,
	total float not null,
	primary key (order_id),
	foreign key (customer_id)
			references customers(id),
	foreign key (product_id)
			references products(id)
);

/*Insert data to customers table*/
insert into public.customers (id, customer_name)
values (1, 'Iqbal Ardy'),
	   (2, 'Lala Po'),
	   (3, 'John Doe'),
	   (4, 'Rowan Atkinson');
	   
/*Insert data to products table*/
insert into public.products (id, name)
values (1, 'balon'),
	   (2, 'kemeja'),
	   (3, 'anting-anting'),
	   (4, 'tas jinjing');

/*Insert data to Orders table*/
insert into public.Orders (order_id, customer_id, product_id, order_date, total)
values (1, 1, 1, '2022-10-23', 2),
	   (2, 1, 2, '2022-10-23-', 1),
	   (3, 2, 3, '2022-10-24', 3),
	   (4, 3, 4, '2022-10-25', 5),
	   (5, 4, 3, '2022-10-26', 10),
	   (6, 4, 1, '2022-10-26', 5),
	   (7, 2, 3, '2022-10-27', 1),
	   (8, 3, 2, '2022-10-27', -1);

/*Update data in Orders table*/
update Orders
set total = 1
where order_id = 5;

/*Delete data in Orders table*/
delete from Orders
where order_id = 8;

/*View information of someone who bought 'anting-anting'*/
select customers.customer_name, products.name, sum(Orders.total)
from customers
	full outer join Orders on customers.id = Orders.customer_id
	full outer join products on customers.id = products.id
group by customers.customer_name, products.name;
