package models

type Student struct {
	SID       string `bson:"_student_id,omitempty" json:student_id,omitempty"`
	FirstName string `bson:"f_name" json:"f_name"`
	LastName  string `bson:"l_name" json:"l_name"`
	Grade     string `bson:"grade" json:"grade"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
}
