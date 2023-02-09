package routes

import (
	"machine-test-1/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(newMux *mux.Router) {
	newMux.HandleFunc("/insert", handlers.InsertStudentHandler)
	newMux.HandleFunc("/get/{id}", handlers.GetStudentHandler)
	newMux.HandleFunc("/get-all-students", handlers.GetAllStudentsHandler)
	newMux.HandleFunc("/update/{id}", handlers.UpdateStudentHandler)
	newMux.HandleFunc("/delete/{id}", handlers.DeleteStudentHandler)
	newMux.HandleFunc("/drop-collection", handlers.DropCollectionHandler)
}
