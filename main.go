package main

// importing
import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//global decleration of database variable

var db *gorm.DB

// defining structure

type Hotel struct {
	RNO     uint      `json:"rno" gorm:"primary_key"`
	RTYPE   string    `json:"rtype" binding:"required"`
	RMNAME  string    `json:"rmname" binding:"required"`
	RMEMBER string    `json:"rmember" binding:"required"`
	RMALE   int       `json:"rmale" binding:"required"`
	RFEMALE int       `json:"rfemale"`
	ENTRY   time.Time `json:"entry"  binding:"required"`
	EXIT    time.Time `json:"exit"`
	RBILL   int       `json:"rbill" binding:"required"`
	RSTATUS string    `json:"rstatus" binding:"required"`
}

// database deceleration

func initDB() {
	var err error

	db, err = gorm.Open("sqlite3", "hotel.db")

	if err != nil {
		panic("Connection not Established" + err.Error())
	}
	db.AutoMigrate(&Hotel{})

}

// inserting detail of customer
func insertData(c *gin.Context) {
	var hotel Hotel
	if err := c.BindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error()})
		return
	}
	db.Create(&hotel)
	c.JSON(http.StatusCreated, hotel)
}

// getting detail of user
func getData(c *gin.Context) {
	var hotel Hotel
	var rno = c.Param("rno")
	if err := db.First(&hotel, rno).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "RNO does not exist"})
		return
	}
	c.JSON(http.StatusOK, hotel)
}

// deleting  detail of user
func deleteData(c *gin.Context) {
	var hotel Hotel
	var rno = c.Param("rno")
	if err := db.First(&hotel, rno).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "RNO does not exist"})
		return
	}
	db.Delete(&hotel)
	c.JSON(http.StatusOK, gin.H{
		"message: ": "Deletion is Successful....."})
}

// updating data of customer
func updateData(c *gin.Context) {
	var photel Hotel
	var rno = c.Param("rno")
	if err := db.First(&photel, rno).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "RNO does not exist"})
		return
	}

	var nhotel Hotel
	if err := c.BindJSON(&nhotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error()})
		return
	}
	db.Model(&photel).Updates(nhotel)
	c.JSON(http.StatusOK, photel)
}

// reteriving all data
func allData(c *gin.Context) {
	var hotel []Hotel
	db.Find(&hotel)
	c.JSON(http.StatusOK, hotel)
}

//entry function

func main() {
	initDB()
	r := gin.Default()
	r.POST("/hotel", insertData)
	r.GET("/hotel/:rno", getData)
	r.DELETE("/hotel/:rno", deleteData)
	r.PUT("/hotel/:rno", updateData)
	r.GET("/hotel", allData)
	r.Run(":5555")
}
