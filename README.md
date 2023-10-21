# RESTful_Blogging_API

This project details my journey in creating a RESTful Blogging API in Golang.
Tech stack used will be 
* GIN web framework
* Postgres database

In the database a simple table to store blog data will be created
```CREATE TABLE BLOG( ID serial PRIMARY KEY, TITLE VARCHAR( 60 ) NOT NULL, CONTENT TEXT)```
We limit title to 60 characters as that is the ideal title length for blogs to be visible on search engines like google.

Possible Operations
1. View all blog posts
2. View Specific blog post by id
3. Create a new blog post
4. Update an existing blog post
5. Delete a blog post
