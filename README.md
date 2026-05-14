# crud-gin
 
REST API for user management built with Go, Gin, GORM and PostgreSQL.
 
## Tech Stack
 
- **Go** — primary language
- **Gin** — HTTP framework
- **GORM** — ORM
- **PostgreSQL** — database
- **Docker** — containerization

## Getting Started
 
### Prerequisites
 
- [Docker](https://www.docker.com/)
### Setup
 
1. Clone the repository:
```bash
git clone <repository-url>
cd crud-gin
```
 
2. Create `.env` file based on `.env.example`:
```bash
cp .env.example .env
```
 
3. Start the application:
```bash
task dev
```
 
The server will be available at `http://localhost:8000`.

## API Endpoints
 
### Users
 
#### Create user
```
POST /users
```
 
Request body:
```json
{
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true
}
```
 
Response `201 Created`:
```json
{
  "id": 1,
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true,
  "created_at": "2026-05-14T20:30:11+07:00",
  "updated_at": "2026-05-14T20:30:11+07:00",
  "deleted_at": null
}
```
 
---
 
#### Get all users
```
GET /users
```
 
Response `200 OK`:
```json
{
  "users": [
    {
      "id": 1,
      "name": "John",
      "username": "@johndoe",
      "email": "johndow123@gmail.com",
      "phone_number": "+71234567890",
      "age": 25,
      "is_smart": true,
      "created_at": "2026-05-14T20:30:11+07:00",
      "updated_at": "2026-05-14T20:30:11+07:00",
      "deleted_at": null
    }
  ]
}
```
 
---
 
#### Get user by ID
```
GET /users/:id
```
 
Response `200 OK`:
```json
{
  "id": 1,
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true,
  "created_at": "2026-05-14T20:30:11+07:00",
  "updated_at": "2026-05-14T20:30:11+07:00",
  "deleted_at": null
}
```
 
Response `404 Not Found`:
```json
{
  "error": "user with id 1 not found"
}
```
 
---
 
#### Update user
```
PUT /users/:id
```
 
Request body (all fields optional):
```json
{
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true,
}
```
 
Response `200 OK`:
```json
{
  "id": 1,
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true,
  "created_at": "2026-05-14T20:30:11+07:00",
  "updated_at": "2026-05-14T20:30:11+07:00",
  "deleted_at": null
}
```
 
---
 
#### Delete user
```
DELETE /users/:id
```
 
Response `200 OK`:
```json
{
  "id": 1,
  "name": "John",
  "username": "@johndoe",
  "email": "johndow123@gmail.com",
  "phone_number": "+71234567890",
  "age": 25,
  "is_smart": true,
  "created_at": "2026-05-14T20:30:11+07:00",
  "updated_at": "2026-05-14T20:30:11+07:00",
  "deleted_at": "2026-05-14T20:33:11+07:00"
}
```
 
Response `404 Not Found`:
```json
{
  "error": "user with id 1 not found"
}
```
 
---