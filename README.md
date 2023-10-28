# RESTful_Blogging_API

This project details my journey in creating a RESTful Blogging API in Golang.
Tech stack used will be 
* GIN web framework
* GORM
* Postgres database
## DataBase
In the database a simple table to store blog data will be created.

```CREATE TABLE BLOGS( ID serial PRIMARY KEY, TITLE VARCHAR( 60 ) NOT NULL, CONTENT TEXT)```

We limit title to 60 characters as that is the ideal title length for blogs to be visible on search engines like google.

## Configuration
In order to connect to the database, we maintain a Config.yaml file to externalize configuration details and update them as needed.
A simple example:

```yaml
host: localhost
port: 5432
user: david
password: postgres
dbname: postgres
```

## API
Possible Operations
1. View all blog posts - GET - "http://localhost:8080/blogs"
2. View Specific blog post by id - GET - "http://localhost:8080/blogs/:id"
3. Create a new blog post - POST - "http://localhost:8080/blogs"
4. Update an existing blog post - PUT - "http://localhost:8080/blogs/:id"
5. Delete a blog post - DELETE - "http://localhost:8080/blogs/:id"
 Where id is a path variable.