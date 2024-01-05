package entities

import "github.com/google/uuid"

type UserEntity struct {
	ID        uuid.UUID `bson:"_id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Age       int       `bson:"age"`
}
