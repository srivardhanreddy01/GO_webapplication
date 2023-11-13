package handlers

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/srivardhanreddy01/webapplication_go/api/models"
	"gorm.io/gorm"
)

type AllAssignmentsHandlerDependencies struct {
	DB *gorm.DB
}

func Assignments(deps AllAssignmentsHandlerDependencies) ([]models.Assignment, error) {
	var assignments []models.Assignment

	if err := deps.DB.Find(&assignments).Error; err != nil {
		return nil, err
	}

	return assignments, nil
}

func AssignmentById(deps AllAssignmentsHandlerDependencies, id int) (*models.Assignment, error) {
	var assignment models.Assignment

	result := deps.DB.Where("id = ?", id).First(&assignment)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}

	return &assignment, nil
}

func GetAllAssignments(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	user, ok := r.Context().Value("user").(*models.User)

	if !ok || user == nil {
		sendJSONResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dependencies := AllAssignmentsHandlerDependencies{
		DB: db,
	}

	assignments, err := Assignments(dependencies)
	if err != nil {
		sendJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, assignments)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var assignment models.Assignment
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		sendJSONResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := db.Create(&assignment).Error; err != nil {
		sendJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, assignment, http.StatusCreated)
}

func GetAssignment(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id := vars["id"]
	// fmt.Println(id)
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		sendJSONResponse(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)

	if !ok || user == nil {
		sendJSONResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dependencies := AllAssignmentsHandlerDependencies{
		DB: db,
	}

	assignments, err := AssignmentById(dependencies, assignmentID)
	if err != nil {
		sendJSONResponse(w, "Assignment not found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, assignments)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id := vars["id"]
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		sendJSONResponse(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		sendJSONResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dependencies := AllAssignmentsHandlerDependencies{
		DB: db,
	}

	assignment, err := AssignmentById(dependencies, assignmentID)
	if err != nil {
		sendJSONResponse(w, "Assignment not found", http.StatusNotFound)
		return
	}

	if err := db.Delete(&assignment).Error; err != nil {
		sendJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id := vars["id"]
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		sendJSONResponse(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		sendJSONResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dependencies := AllAssignmentsHandlerDependencies{
		DB: db,
	}

	assignment, err := AssignmentById(dependencies, assignmentID)
	if err != nil {
		sendJSONResponse(w, "Assignment not found", http.StatusNotFound)
		return
	}

	var updatedAssignment models.Assignment
	if err := json.NewDecoder(r.Body).Decode(&updatedAssignment); err != nil {
		sendJSONResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	updatedAssignment.ID = uint(assignmentID)
	assignment.Name = updatedAssignment.Name
	assignment.Points = updatedAssignment.Points
	assignment.NumOfAttempts = updatedAssignment.NumOfAttempts
	assignment.Deadline = updatedAssignment.Deadline

	if err := db.Save(&assignment).Error; err != nil {
		sendJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//sendJSONResponse(w, "Assignment updated", http.StatusNoContent)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode ...int) {
	response := struct {
		Data interface{} `json:"data"`
	}{
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")

	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		sendJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
