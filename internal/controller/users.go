package controller

import (
	"fmt"
	"net/http"

	"github.com/TransportMall/payment-systems/internal/model"
	"github.com/TransportMall/payment-systems/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserSignup function is used for customer and dairy signup
func UserSignup(c *gin.Context) {

	var user model.Users

	if err := c.ShouldBind(&user); err != nil {
		utils.ValidationError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		utils.ValidationError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.ResponseHandler(c, http.StatusOK, gin.H{"success": true, "users": user})
}

// UserLogin function is used for customer and dairy login
func UserLogin(c *gin.Context) {

	var user model.UserData

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	if mobile == "" || password == "" {
		utils.ErrorHandler(c, http.StatusUnprocessableEntity, fmt.Errorf("Invalid Request"))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Table("users").Where("mobile = ? AND password = ? AND status = ? AND deleted_at IS NULL", mobile, password, 1).First(&user).Error; err != nil {
		utils.ResponseHandler(c, http.StatusOK, gin.H{"success": false, "users": nil})
		return
	}

	token, err := utils.CreateJWTToken(user.ID)
	if err != nil {
		utils.ErrorHandler(c, http.StatusUnprocessableEntity, fmt.Errorf("Unable to create token"))
		return
	}
	utils.ResponseHandler(c, http.StatusOK, gin.H{"success": true, "users": user, "token": token.AccessToken})
}

// MyProfile function is used for getting account information
func MyProfile(c *gin.Context) {

	var user model.UserData
	db := c.MustGet("db").(*gorm.DB)
	authUser := c.MustGet("auth").(jwt.MapClaims)

	if err := db.Table("users").Where("id = ?", authUser["user_id"]).First(&user).Error; err != nil {
		utils.ResponseHandler(c, http.StatusOK, gin.H{"success": false, "users": nil})
		return
	}

	utils.ResponseHandler(c, http.StatusOK, gin.H{"success": true, "users": user})
}
