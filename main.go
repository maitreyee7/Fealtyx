package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Define a struct for Student data
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Profile string `json:"profile,omitempty"`
}

// Global variables for student data and synchronization
var (
	studentData = make(map[int]Student)
	mutex       sync.Mutex
	studentID   = 1
)

// Function to handle student creation
func addStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	json.NewDecoder(r.Body).Decode(&newStudent)

	mutex.Lock()
	newStudent.ID = studentID
	studentData[studentID] = newStudent
	studentID++
	mutex.Unlock()

	json.NewEncoder(w).Encode(newStudent)
}

// Function to retrieve all students
func retrieveStudents(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var allStudents []Student
	for _, student := range studentData {
		allStudents = append(allStudents, student)
	}
	json.NewEncoder(w).Encode(allStudents)
}

// Function to retrieve student by ID
func getStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	student, found := studentData[id]
	mutex.Unlock()

	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

// Function to generate a summary for a student by ID
func fetchStudentProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/students/profile/"):])
	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	student, found := studentData[id]
	mutex.Unlock()

	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	if student.Profile == "" {
		student.Profile, _ = createProfileDescription(student)
		mutex.Lock()
		studentData[id] = student
		mutex.Unlock()
	}

	json.NewEncoder(w).Encode(map[string]string{"profile": student.Profile})
}

// Helper function to create a profile description
func createProfileDescription(student Student) (string, error) {
	return fmt.Sprintf("Name: %s, Age: %d, Contact: %s", student.Name, student.Age, student.Email), nil
}

// Function to modify student data by ID
func modifyStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	var updatedStudent Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)

	mutex.Lock()
	student, found := studentData[id]
	if found {
		updatedStudent.ID = student.ID
		studentData[id] = updatedStudent
	}
	mutex.Unlock()

	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedStudent)
}

// Function to delete a student by ID
func removeStudentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, found := studentData[id]
	if found {
		delete(studentData, id)
	}
	mutex.Unlock()

	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			addStudent(w, r)
		case "GET":
			retrieveStudents(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getStudentByID(w, r)
		case "PUT":
			modifyStudentByID(w, r)
		case "DELETE":
			removeStudentByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/profile/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fetchStudentProfile(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
