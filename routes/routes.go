package routes

import (
	"machine-test-1/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(newMux *mux.Router) {

	var hanlder handlers.JsonResponse

	newMux.HandleFunc("/insert", hanlder.InsertStudentHandler)
	newMux.HandleFunc("/get/{id}", hanlder.GetStudentHandler)
	newMux.HandleFunc("/get-all-students", hanlder.GetAllStudentsHandler)
	newMux.HandleFunc("/update/{id}", hanlder.UpdateStudentHandler)
	newMux.HandleFunc("/delete/{id}", hanlder.DeleteStudentHandler)
	newMux.HandleFunc("/drop-collection", hanlder.DropCollectionHandler)
}
