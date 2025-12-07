package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Firstname    *string            `json:"first_name" validate:"required, min=2, max=10"`
	Lastname     *string            `json:"last_name" validate:"required, min=2, max=10"`
	Password     *string            `json:"password" validate:"required, min=8, max=100"`
	Email        *string            `json:"email" validate:"email, required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token" validate:"required"`
	Role         *string            `json:"role" validate:"required"`
	RefreshToken *string            `json:"refreshtoken" validate:"required, eq=ADMIN|eq=USER"`
	Createdat    *time.Time         `json:"createdat" validate:"required"`
	Updatedat    *time.Time         `json:"updatedat" validate:"required"`
	UserID       *string            `json:"userid"`
}
