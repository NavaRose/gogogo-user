package users

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primarykey"`
	Name         string `binding:"required"`
	Password     string `binding:"required"`
	Email        string `binding:"required"`
	AccessToken  string
	RefreshToken string
	RoleId       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func GetUserSchema() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Email": &graphql.Field{
					Type: graphql.String,
				},
				"AccessToken": &graphql.Field{
					Type: graphql.String,
				},
				"RefreshToken": &graphql.Field{
					Type: graphql.String,
				},
				"RoleId": &graphql.Field{
					Type: graphql.Int,
				},
				"CreatedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"UpdatedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"DeletedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
