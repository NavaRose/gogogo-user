package main

import (
	"github.com/NavaRose/gogogo/users/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"net/http"
)

func List(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var users []models.User
	result := db.Find(&users)
	message := "Oke, this is user list"
	if result.RowsAffected == 0 {
		message = "No users found"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    users,
	})
}

func Detail(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	result := db.First(&user, ctx.Param("id"))

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is user detail",
		"data":    result,
	})
}

func Create(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	newUser := models.User{}
	err := ctx.ShouldBindWith(&newUser, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to create user: Validation failed",
			"data":    err.Error(),
		})
		return
	}

	result := db.Create(&newUser)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Create user failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is user create",
		"data":    newUser,
	})
}

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is user update",
	})
}

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This is user delete",
	})
}
