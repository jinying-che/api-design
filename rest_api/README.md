This API has four endpoints:

POST /users: Creates a new user with the specified name and email.
GET /users/:id: Retrieves an existing user with the specified ID.
PUT /users/:id: Updates an existing user with the specified ID.
DELETE /users/:id: Deletes an existing user with the specified ID.

```
POST /users
{
  "name": "John Smith",
  "email": "john@example.com"
}

GET /users/:id

PUT /users/:id
{
  "name": "John Smith",
  "email": "john@example.com"
}

DELETE /users/:id

```

