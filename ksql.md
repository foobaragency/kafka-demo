## Create stream from product events
```sql
CREATE STREAM PRODUCTS_STREAM (
    ID VARCHAR,
    CE_TYPE VARCHAR,
    NAME VARCHAR,
    DESCRIPTION VARCHAR,
    IMAGE_URL VARCHAR,
    STOCK INTEGER
) WITH (
    KAFKA_TOPIC='products', VALUE_FORMAT='JSON'
);
```

## Create stream from order events
```sql
CREATE STREAM ORDERS_STREAM (
    ID VARCHAR,
    CE_TYPE VARCHAR,
    QUANTITY INTEGER
) WITH (
    KAFKA_TOPIC='orders', VALUE_FORMAT='JSON'
);
```

## Sum all the product refills
```sql
CREATE TABLE PRODUCT_REFILLS AS
	SELECT ID, SUM(STOCK) AS stock FROM PRODUCTS_STREAM
    WHERE CE_TYPE = 'kafka_demo.RefillStock' GROUP BY ID;
```

## Sum all the order quantities
```sql
CREATE TABLE ORDER_PRODUCTS AS
	SELECT ID, SUM(QUANTITY) AS quantity FROM  ORDERS_STREAM 
    WHERE CE_TYPE = 'kafka_demo.BindProductToOrder' GROUP BY ID;
```

## Diff products and orders
```sql
CREATE TABLE PRODUCTS_STOCK AS
    SELECT p.id as ID, (p.STOCK - o.QUANTITY) AS current_stock FROM  PRODUCT_REFILLS AS p
    LEFT JOIN ORDER_PRODUCTS AS o
    ON p.id = o.id;
```