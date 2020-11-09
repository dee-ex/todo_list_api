# Todo List API
A RESTful API template for simple todo application with Go.
# User
## POST Register User
```
/user/register
```
Register a new user, save it to database and return as JSON.  
BODY
```
{
    "username": "trung",
    "email": "trung@gmail.com",
    "password": "hcmut"
}
```
## POST Login
```
/user/login
```
Logged in to registered user account then save `access_token` to database and return.  
BODY
```
{
    "username": "trung",
    "password": "hcmut"
}
```
## POST Logout
```
/user/logout
```
Logged out the user and clear token from database, authorization needed.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## GET Logged User Profile
```
/user/me
```
Return user profile that logged in.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## PUT Update User Profile
```
/user/me
```
Update logged in user profile and save it to database.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
BODY
```
{
    "password": "ftu"
}
```
## POST Reset Password
```
/user/reset-password
```
Generate token to reset password.  
BODY
```
{
    "email": "trung@gmail.com"
}
```
## POST Submit Reset Password
```
/user/reset-password/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
```
Reset to new password and save it to database.  
BODY
```
{
    "password": "ueh"
}
```
## DELETE User
Delete and remove user from database.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
# Task
## POST Create Task
```
/task
```
Create a new task, save it to database and return as JSON.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
BODY
```
{
    "name": "import task",
    "detail": "this task should be done first"
}
```
## GET Get All Tasks
```
/task
```
Get all task from user and return as JSON.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## GET Get All Completed Tasks
```
/task?completed=true
```
Get all task from user that are sorted by done status and return as JSON.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## GET Get Detail Task by ID
```
/task/123
```
Get detail a task from user by ID and return as JSON.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## PUT Update Task by ID
```
/task/123
```
Update task by ID.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
BODY
```
    "name": "normal",
    "detail": "this task isn't important"
```
## PUT Archive Task by ID
```
/task/123/archive
```
Archive task by ID.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## Delete Restore Task by ID
```
/task/123/archive
```
Restore task by ID.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk
## DELETE Delete Task by ID
```
/task/123
```
Delete task from database by ID.  
HEADERS
KEY | VALUE
--- | ---
access_token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ4NTg5NTUsInVzZXJuYW1lIjoibG9uZyJ9.ARsDLE8lqSlNTfhJmfKxcC6wwqikqRWwplKtx53CRVk