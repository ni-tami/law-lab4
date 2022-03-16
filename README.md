# Lab 4

CRUD + FileUpload in Golang using echo and MongoDB

---

### Dependencies

- [echo]("https://github.com/labstack/echo/v4")
- [go mongodb driver]("https://github.com/mongodb/mongo-go-driver")

---

### Details:

Main files:

1. main.go
2. fileUpload.go

---

> main.go

CRUD operation on cookie database.

MongoDB
Database: `mongo`  
Collection: `cookies`

Dummy data available in `cookies.json`

- Open MongoDB Compass
- Connect to mongodb://localhost:27017 (default port)
- Import cookies.json

Request from `localhost:8080`:

- POST /cookies  
  payload: Cookies in request body  
  return insert operation info
- GET /cookies  
  return all cookies in db
- GET /cookies/`:tag`  
  `tag` path param of cookie  
  return cookie with that `tag`
- PATCH /cookies/`:tag`  
  `tag` path param of cookie  
  return patch operation info
- DELETE /cookies/`:tag`  
  `tag` path param of cookie
  return delete count

> **fileUpload.go**

Upload file into database using `gridfs`

Database: `mongo`  
Collection: `fs.files` and `fs.chunks`

- run via
  ```
  go run main.go <filename>
  ```
  or compile first
  ```
  go build
  main.exe <filename>
  ```
