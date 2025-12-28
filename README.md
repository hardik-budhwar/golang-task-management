
# golang-task-management
A robust containerized task management system built with Golang, Gin, and GORM. This system handles user authentication, authorization, and access management, allowing users to create, update, and manage tasks securely.


# key-features
1. Support for bulk upload of tasks using CSV.
2. Secure user registration and authentication using JWT (JSON Web Tokens).


# pre-requisites
Make sure you have the following installed:
1. Docker and Docker Compose
2. Golang

# getting-started
1. Clone the repository
2. Copy the .env.example file to .env and update the values with your own configuration:
3. Build and run docker containers -> { docker-compose up -d }

# api endpoints
1. POST /user/register: Register a new user.

2. POST /user/login: Log in with registered user credentials and receive a JWT token.

3. PUT /user/logout: Log out and invalidate the JWT token.

4. DELETE /user/delete: Invalidate the JWT token and Delete User.

5. GET /task/:id: Get a single task by ID.

6. GET /task/: Get all tasks.

8. POST /task/create: Create a new task.

9. POST /task/bulkupload: Create a new tasks using bulk upload (explained below).

10. PUT /task/update/:id: Update an existing task by ID.

11. DELETE /task/delete/:id: Delete a task by ID.

12. GET /task/user/:user_id : Get all tasks of a particular user

13. Get /user/ : Get all users 

