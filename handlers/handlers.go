package handlers

import (
	"log"
	"machine-test-1/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *JsonResponse) InsertStudentHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	fname, lname, grade := r.PostFormValue("fname"), r.PostFormValue("lname"), r.PostFormValue("grade")
	email, phone := r.PostFormValue("email"), r.PostFormValue("phone")

	err = usecase.InsertStudentUsecase(fname, lname, grade, email, phone)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "User has been inserted",
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}

func (app *JsonResponse) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := usecase.GetStudentUsecase(id)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully fetched user data",
		Data:    data,
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}

func (app *JsonResponse) GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := usecase.GetAllStudentsUsecase()
	if err != nil {
		log.Println(err)
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully fetched all user data",
		Data:    data,
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}

func (app *JsonResponse) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fname, lname, grade := r.PostFormValue("fname"), r.PostFormValue("lname"), r.PostFormValue("grade")
	email, phone := r.PostFormValue("email"), r.PostFormValue("phone")

	err := usecase.UpdateStudentUsecase(id, fname, lname, grade, email, phone)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully updated the user data",
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}

func (app *JsonResponse) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := usecase.DeleteStudentUsecase(id)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully deleted the user data",
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}

func (app *JsonResponse) DropCollectionHandler(w http.ResponseWriter, r *http.Request) {
	err := usecase.DropCollectionUsecase()
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "Successfully dropped the collection",
	}

	app.WriteJSON(w, http.StatusAccepted, resp)
}
