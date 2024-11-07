Student Management API


A simple REST API for managing student records, built with Go. This API allows users to create, view, update, delete, and retrieve profile summaries for students.

Features
Create Student: Add a new student to the database.
Retrieve All Students: View all students.
Retrieve Student by ID: Get details of a specific student by ID.
Update Student: Modify details of an existing student.
Delete Student: Remove a student by ID.
Profile Summary: Generate and fetch a summary for a student.
Installation
Clone this repository:

sh
Copy code
git clone https://github.com/yourusername/student-api.git
cd student-api
Run the API:

sh
Copy code
go run main.go
The server will run on http://localhost:8080.

API Endpoints
1. Add a New Student
Endpoint: /students
Method: POST
Payload:
json
Copy code
{
  "name": "John Doe",
  "age": 20,
  "email": "johndoe@example.com"
}
2. Retrieve All Students
Endpoint: /students
Method: GET
3. Get Student by ID
Endpoint: /students/{id}
Method: GET
4. Update Student
Endpoint: /students/{id}
Method: PUT
Payload:
json
Copy code
{
  "name": "Jane Doe",
  "age": 21,
  "email": "janedoe@example.com"
}
5. Delete Student
Endpoint: /students/{id}
Method: DELETE
6. Get Profile Summary
Endpoint: /students/profile/{id}
Method: GET
Dependencies
Go: Ensure Go is installed and properly set up on your machine.
License
This project is licensed under the MIT License.

