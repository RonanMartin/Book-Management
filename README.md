# Bookstore Management System

This project has the folowing characteristcs:

1. It's composed by Database - Mysql;
2. The GORM package to interact with the database;
3. We also are using JSON marshall and unmarshall;
4. All project is separate in a structure of folders;
5. We also are using the package Gorilla Mux.

## **Project Struture**

![Project Structure](projectStructure.jpg)

We have two main folders, **cmd** and **pkg**, in the **cmd** folder we only have the **main.go** file, and in the **pkg** folder we will have the **config**, **controllers**, **models**,**routes** and **utils** folders.

**config**: We have **app.go** file, wich help us to conect with our database.
**controllers**: We have **book-controller.go** file, wich have the functions that help us process the data that we'll get for as response, and also the request that we'll get from the user and the response that we need send.
**models**: We have **book.go** file, wich help us to create those structs and models that will be used by our database.
**routes**: We just have **bookstore-routes.go** file, with our routes that will see later.
**utils**: we just have the code for marshalling and unmarshalling JSON from our respose and our request.

![Routes](routes.jpg)
