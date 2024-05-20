BEGIN;

INSERT INTO WAREHOUSES
VALUES (1, 'north', true),
       (2, 'east', true),
       (3, 'west', true),
       (4, 'south', true),
       (5, 'north-south', true)
;


INSERT INTO PRODUCTS
VALUES (1, 't-shirt', '1a1', 5),
       (2, 'shorts', '1a2', 3),
       (3, 'hoodie', '1a3', 8),
       (4, 'cap', '1a4', 1),
       (5, 'jacket', '1a5', 15)
;

INSERT INTO WAREHOUSE_PRODUCTS
VALUES (1, 1, 1, 5),
       (2, 2, 1, 8),
       (3, 3, 1, 13),
       (4, 1, 2, 4),
       (5, 1, 3, 1),
       (6, 4, 5, 2),
       (7, 5, 4, 9),
       (8, 5, 3, 12),
       (9, 2, 3, 19),
       (10, 4, 1, 4)
;

COMMIT;