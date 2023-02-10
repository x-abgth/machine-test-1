package database

import (
	"context"
	"log"
	"machine-test-1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Respository struct {
	Client  *mongo.Client
	Student models.Student
}

func (s Respository) InsertStduent(student models.Student) error {

	collection := s.Client.Database("students").Collection("students")

	_, err := collection.InsertOne(context.TODO(), models.Student{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Grade:     student.Grade,
		Email:     student.Email,
		Phone:     student.Phone,
	})

	if err != nil {
		log.Println("Error inserting", err)
		return err
	}

	return nil
}

func (s Respository) GetAllStudentsDetails() ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := s.Client.Database("students").Collection("students")

	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var student []models.Student

	for cursor.Next(ctx) {
		var item models.Student

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding log into slice", err)
		} else {
			student = append(student, item)
		}
	}

	return student, nil
}

func (s Respository) GetStudent(id string) (models.Student, error) {
	var student models.Student

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := s.Client.Database("students").Collection("students")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return student, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&student)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s Respository) Update(id string, student models.Student) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := s.Client.Database("students").Collection("students")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{
				"$set", bson.D{
					{"f_name", student.FirstName},
					{"l_name", student.LastName},
					{"grade", student.Grade},
					{"email", student.Email},
					{"phone", student.Phone},
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Respository) DeleteOneStudent(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := s.Client.Database("students").Collection("students")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Respository) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := s.Client.Database("students").Collection("students")

	if err := collection.Drop(ctx); err != nil {
		return err
	}

	return nil
}
