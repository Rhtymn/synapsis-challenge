# synapsis-challenge

This project was created to fulfill the backend challenge test at synapsis.id. This project was created using golang programming language using the gin framework and using postgres and redis as data storage.

## Entity Relationship Diagram (ERD)

The figure below displays the results of the ERD analysis and design. The ERD displayed has represented the relationship between tables and attributes owned by each table. However, it does not display the constraints of each attribute. The ERD that has been defined will be implemented into the database using Postgres.

![ERD](./erd.png)

## Redis

In this project, redis is used to store customer cart data. The structure used is quite simple using hset. The key used in the hset storage uses the `cart:accountID` keyword. The following is an example of data representation stored in redis.

`cart:1 1 "1:The Psychology Money:1:76500:1:andi shop"`

`cart:1 2 "2:Zero To One:1:57500:1:raharja shop"`

## Feature

The successfully developed features can be seen as follows.

- Login and register as a customer
- Email verification
- Update profiles
- View product list with the capability to filter by category, search by name, and sort by name and price.
- Add products to cart
- Remove products from cart
- Checkout (create a transaction) and make a payment

There are limitations in some of the features created. Here are the limitations that need to be considered.

- Customers cannot add products to the cart if they have not verified their email and filled out their profile.
- Currently, each product put into the cart that comes from the same store will be considered a different transaction. So that each transaction made is only for one product only with the number of products as added.
- Customers make payments by uploading proof of payment. Later, the admin can confirm manually (the feature has not been created yet).

## Postman Collection

You can import the postman collection from the root of this folder.

## Docker HUB

https://hub.docker.com/repository/docker/rhtymn/synapsis-challenge/general

## Account

Data seeding has been done in this application. All accounts created have the same password. The following customer accounts can be used:

`email: roihan@gmail.com`
`password: password`
