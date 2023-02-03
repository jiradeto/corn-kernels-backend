-- insert into products table
INSERT INTO products (name, description)
VALUES 
	("Product A", "Description of product A"),
	("Product B", "Description of product B"),
	("Product C", "Description of product C");

-- insert into product_stocks table
INSERT INTO product_stocks (product_id, quantity)
VALUES 
	(1, 0),
	(2, 0),
	(3, 0);

WITH dates AS (
	SELECT 
		date('2022-01-01') + (rand() * (date('2022-02-28') - date('2022-01-01'))) AS random_date
	FROM 
		generate_series(1, 10) 
), stock_movements AS (
	SELECT 
		1 AS product_id,
		'Movement' || row_number() OVER () AS description,
		round((rand() * 100)::numeric, 2) AS quantity,
		CASE 
			WHEN rand() > 0.5 THEN 'IN'
			ELSE 'OUT'
		END AS movement_type,
		random_date
	FROM 
		dates
)
INSERT INTO stock_movements (product_id, description, quantity, movement_type, created_at)
SELECT 
	product_id, 
	description, 
	quantity, 
	movement_type, 
	random_date::timestamp
FROM 
	stock_movements;
This code generates 10 random dates within the range from January 1st to February 28th, 2022 and 10 random quantities. The CASE statement is used to randomly generate either "IN" or "OUT" for the "movement_type". The generated data is then inserted into the "stock_movements" table.




