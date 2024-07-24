package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"os"
	//model "users/Models"
)

func RouteCreator() func(*gin.Engine) {
	if os.Getenv("REQUEST_METHOD") == "graphql" {
		return applyGraphQL
	}
	return applyRoutes
}

func applyRoutes(engine *gin.Engine) {
	engine.GET("/", List)
	engine.GET("/:id", Detail)
	engine.POST("/", Create)
	engine.PUT("/:id", Update)
	engine.DELETE("/:id", Delete)
}

func applyGraphQL(engine *gin.Engine) {
	engine.POST("/", func(ctx *gin.Context) {
		userType := model.GetUserSchema()

		queryType := graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"list": &graphql.Field{
					Type:        graphql.NewList(userType),
					Description: "Get user list",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						db := core.InitDatabaseWithoutEngine()
						var users []model.User
						result := db.Find(&users)
						if result.RowsAffected > 0 {
							return users, nil
						}
						return nil, nil
					},
				},
				"detail": &graphql.Field{
					Type:        userType,
					Description: "Get user by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(int)
						if ok {
							db := core.InitDatabaseWithoutEngine()
							var user model.User
							result := db.First(&user, id)
							if result.RowsAffected > 0 {
								return user, nil
							}
						}
						return nil, nil
					},
				},
			},
		})

		schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
		type Request struct {
			Query string `json:"query"`
		}
		var request Request
		if err := ctx.BindJSON(&request); err != nil {
			log.Fatalf("failed to execute graphql operation, errors: %+v", err)
			return
		}

		result := graphql.Do(graphql.Params{Schema: schema, RequestString: request.Query})
		ctx.JSON(http.StatusOK, result)
	})
}
