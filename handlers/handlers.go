package handlers

import (
	"log"
	"machine-test-1/usecase"
	"machine-test-1/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func InsertStudentHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	fname, lname, grade := r.FormValue("fname"), r.FormValue("lname"), r.FormValue("grade")
	email, phone := r.FormValue("email"), r.FormValue("phone")

	err = usecase.InsertStudentUsecase(fname, lname, grade, email, phone)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "User has been inserted",
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := usecase.GetStudentUsecase(id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched user data",
		Data:    data,
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}

func GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := usecase.GetAllStudentsUsecase()
	if err != nil {
		log.Println(err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all user data",
		Data:    data,
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fname, lname, grade := r.FormValue("fname"), r.FormValue("lname"), r.FormValue("grade")
	email, phone := r.FormValue("email"), r.FormValue("phone")

	err := usecase.UpdateStudentUsecase(id, fname, lname, grade, email, phone)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "Successfully updated the user data",
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := usecase.DeleteStudentUsecase(id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "Successfully deleted the user data",
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}

func DropCollectionHandler(w http.ResponseWriter, r *http.Request) {
	err := usecase.DropCollectionUsecase()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JsonResponse{
		Error:   false,
		Message: "Successfully dropped the collection",
	}

	utils.WriteJSON(w, http.StatusAccepted, resp)
}
