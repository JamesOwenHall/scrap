# Multi-source Querying

## Connections

You can create database connections to later create sources.

    sales_db = mysql("user", "pass", "host", 3306, "sales_db")
    sessions = cassandra("user", "pass", "host", 9042, "sessions_keyspace")

## Sources

A source represents a slice of rows.

    sales_2016 = sales_db.query("SELECT * FROM my_thing WHERE time >= '2016-01-01' AND time < '2017-01-01'")
    sesssions_2016 = sessions.query("SELECT * FROM sessions WHERE time >= '2016-01-01' AND time < '2017-01-01'")
    orders = csv("filename.csv")

## Queries

    SELECT SUM(sales.total), order.id FROM sales_2016 AS sales JOIN orders ON orders.id = sales.id GROUP BY order.id
