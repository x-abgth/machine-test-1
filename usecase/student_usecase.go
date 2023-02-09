package usecase

import (
	"log"
	"machine-test-1/database"
	"machine-test-1/models"
)

var dbInstance database.Respository

func NewDbInstance() {
	dbInstance = database.Respository{
		Client:  database.ConnectMongo(),
		Student: models.Student{},
	}
}

func InsertStudentUsecase(fName, lName, grade, email, phone string) error {
	data := models.Student{
		FirstName: fName,
		LastName:  lName,
		Grade:     grade,
		Email:     email,
		Phone:     phone,
	}

	err := dbInstance.InsertStduent(data)
	if err != nil {
		return err
	}

	return nil
}

func GetStudentUsecase(id string) (models.Student, error) {
	student, err := dbInstance.GetStudent(id)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}
func GetAllStudentsUsecase() ([]models.Student, error) {
	students, err := dbInstance.GetAllStudentsDetails()
	if err != nil {
		return nil, err
	}

	return students, nil
}
func UpdateStudentUsecase(id, fName, lName, grade, email, phone string) error {
	data := models.Student{
		FirstName: fName,
		LastName:  lName,
		Grade:     grade,
		Email:     email,
		Phone:     phone,
	}

	result, err := dbInstance.Update(id, data)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}
func DeleteStudentUsecase(id string) error {
	result, err := dbInstance.DeleteOneStudent(id)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}

func DropCollectionUsecase() error {
	err := dbInstance.DropCollection()
	if err != nil {
		return err
	}

	return nil
}
