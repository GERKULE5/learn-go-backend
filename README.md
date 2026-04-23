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
- [Make](https://www.gnu.org/software/make/)
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
make up
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
  "full_name": "John Doe",
  "age": 25,
  "is_smart": true
}
```
 
Response `201 Created`:
```json
{
  "id": 1,
  "name": "John",
  "full_name": "John Doe",
  "age": 25,
  "is_smart": true
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
      "full_name": "John Doe",
      "age": 25,
      "is_smart": true
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
  "full_name": "John Doe",
  "age": 25,
  "is_smart": true
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
  "name": "Jane",
  "full_name": "Jane Doe",
  "age": 30,
  "is_smart": false
}
```
 
Response `200 OK`:
```json
{
  "id": 1,
  "name": "Jane",
  "full_name": "Jane Doe",
  "age": 30,
  "is_smart": false
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
  "full_name": "John Doe",
  "age": 25,
  "is_smart": true
}
```
 
Response `404 Not Found`:
```json
{
  "error": "user with id 1 not found"
}
```
 
---