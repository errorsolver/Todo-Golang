# Golang Todo API

This is Golang API to create todo list for user. This app use Gin and Gorm framework.

## Detail API

Base URL = 8.219.130.79
Server = AlibabaCloud
Database Service = Neon

## Table Of Contents

1. [User Endpoint](#user-endpoint)  
    1.1 [Sign Up](#sign-up)  
    1.2 [Login](#login)  
    1.3 [Get All Todos](#get-all-users)  
    1.4 [Get User By ID](#get-all-users)  
    1.5 [Update User](#update-user)  
    1.6 [Delete User](#delete-user)  
2. [Todo Endpoint](#todo-endpoint)  
    2.1 [Create Todo](#create-todo)  
    2.2 [Get All Todos By User ID](#get-all-todos)  
    2.3 [Get All Todos By User ID](#get-all-todos-by-user-id)  
    2.4 [Update Todo](#update-todo)  
    2.5 [Delete Todo](#delete-todo)  

## User Endpoint

### **Sign Up**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/user/register`
- **Method:** `POST`
- **Description:** Create a new user.
- **Request Body:**

    ```json
    {
        "name": "Your name",
        "pass": "Your password"
    }
    ```

</details>

### **Login**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/login`
- **Method:** `GET`
- **Description:** logged in to api.
- **Request Body:**

    ```json
    {
        "name": "Your name",
        "pass": "Your password"
    }
    ```

</details>

### **Logout**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/logout`
- **Method:** `GET`
- **Description:** logged out user from api.
- **Request Body:** -

</details>

### **Get All Users**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/user/all`
- **Method:** `GET`
- **Description:** Get all users data from database.
- **Request Body:** -

</details>

### **Get User By ID**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/user/:id`
- **Method:** `GET`
- **Description:** Get user data by id.
- **Request Body:** -

</details>

### **Update User**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/user`
- **Method:** `PUT`
- **Description:** Update the currently logged in user.
- **Request Body:**

    ```json
    {
        "name": "Your name",
        "pass": "Your password"
    }
    ```

</details>

### **Delete User**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/user`
- **Method:** `DELETE`
- **Description:** Delete currently logged in user.
- **Request Body:** -

</details>

## Todo Endpoint

### **Create Todo**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/todo`
- **Method:** `POST`
- **Description:** Create a new todo to currently logged user.
- **Request Body:**

    ```json
    {
        "todo": "Your todo",
    }
    ```

</details>

### **Get All Todos**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/todo/all`
- **Method:** `GET`
- **Description:** Get all todos from database.
- **Request Body:** -

</details>

### **Get All Todos By User ID**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/todo`
- **Method:** `GET`
- **Description:** Get all todos by by user id.
- **Request Body:** -

</details>

### **Update Todo**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/todo`
- **Method:** `PUT`
- **Description:** Edit your todo parameter, like check and uncheck.
- **Request Body:**

    ```json
    {
        "todo": "Your todo",
        "isDone": true or false
    }
    ```

</details>

### **Delete Todo**

<details>
<summary>
    <strong>See Details</strong>
</summary>

- **URL:** `/todo`
- **Method:** `DELETE`
- **Description:** Delete currently user's todo using todo ID.
- **Request Body:** -

</details>
