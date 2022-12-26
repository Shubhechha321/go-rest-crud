How the code is designed:
The backend code is written in Go language and the database used is MongoDB.
The connection with the database is made using MongoDB Atlas.

How to run the code:
Install and configure Go
Clone the project from GitHub: Shubhechha321/go-rest-crud (github.com)
Run the server using the command: go run main.go
Server starts running on port 8080

API endpoints
•	Create a user:
POST/create

•	Get a user with id in the request url:
GET/user/{id}

•	Get all users:
•	GET/users

•	Delete a user with id in the request url:
DELETE /{id}

•	Update a user with id in the request url:
PATCH /{id}

Output of the rest calls:
Create a user:
 

Get all users:![image](https://user-images.githubusercontent.com/69155599/209503958-44acbbc4-d92c-4e30-959a-31d1ce5380f3.png)

Get a single user:![image](https://user-images.githubusercontent.com/69155599/209503939-dfe7f135-704d-402b-9f41-2751f0af08db.png) 

Update a user:![image](https://user-images.githubusercontent.com/69155599/209503916-88b5c08b-7f41-4532-bdb6-9d761ba62096.png)

Delete a user:![image](https://user-images.githubusercontent.com/69155599/209503898-9baa1e04-7a2d-4614-883a-7205b35cb9d3.png)
![image](https://user-images.githubusercontent.com/69155599/209503981-dea374a8-d730-45f8-a336-c9982b802ff1.png)

