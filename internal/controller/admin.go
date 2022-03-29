package controller

import (
	"fmt"
	"net/http"

	"github.com/TransportMall/payment-systems/internal/model"
	"github.com/TransportMall/payment-systems/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetUsers controller function
// @Description Get list of all users
// @Tags Users
// @Produce  json
// @Success 200
// @Router admin/users [get]
func GetUsers(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var users []model.Users
	if err := db.Find(&users).Error; err != nil {
		utils.ErrorHandler(c, http.StatusBadRequest, fmt.Errorf("db cquery error, %s", err))
		return
	}

	utils.ResponseHandler(c, http.StatusOK, gin.H{"users": users})
}
