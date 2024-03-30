package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ielts/component"
	"ielts/component/uploadprovider"
	gincourse2 "ielts/module/course/coursetranspot/gincourse"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}

}
func runService(db *gorm.DB, upProvider uploadprovider.UpLoadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// CRUD
	course := r.Group("/courses")
	{
		course.GET("", gincourse2.ListCourse(appCtx))
		course.POST("", gincourse2.CreateCourse(appCtx))
		course.GET("/:id", gincourse2.GetCourse(appCtx))

	}

	return r.Run()
}
