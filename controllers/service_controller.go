package controllers

import (
	"apiservice/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//access db instance
var db *mongo.Database

func init() {
	db = database.ConfigDB()
}

//example get endpoint
func ServiceGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Resource": "example get endpoint"})
}

//example post endpoint
func ServicePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Resource": "example post endpoint"})
}

//example put endpoint
func ServicePut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Resource": "example put endpoint"})
}
