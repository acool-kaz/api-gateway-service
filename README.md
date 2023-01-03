# README

As part of this repository, the API Gateway service has been implemented

Links to another services:
- [CRUD Service](https://www.github.com/acool-kaz/post-crud-service-server)
- [Web-scapper-service](https://github.com/acool-kaz/parser-service-server)

# Complete task

This project contains three services:

- Service 1: Collects 50 pages of posts from the open API at https://gorest.co.in/public/v1/posts and stores the collected data in a PostgreSQL database. Optional: Data collection can be performed in multiple threads.

- Service 2: Implements CRUD logic for the collected posts, including the ability to retrieve multiple posts, retrieve a specific post, delete a post, modify a post, and create a post.

- Service 3: Acts as an API gateway and provides methods to perform the operations of Service 1 and Service 2. Optional: Service 3 can test coverage and interact with other services via gRPC.

# Tools

- Golang
- gRPC
- PostgreSQL
- Protobuf

# Project structure

```code

```