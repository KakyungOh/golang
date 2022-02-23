package main

import (
	"github.com/gin-gonic/gin"

	"sampleserver/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

// set gin router
func setupRouter() *gin.Engine {

	r := gin.Default()

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/welcome/:name", welcomePathParam)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

// Welcome godoc
// @Summary 테스트입니다
// @Description 아 왤케 어려워
// @name get Name
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} welcomeModel
func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	welcomeMessage := welcomeModel{1, name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}

func main() {

	// programatically set swagger info
	docs.SwaggerInfo_swagger.Title = "Axgate Swagger Example API"
	docs.SwaggerInfo_swagger.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo_swagger.Version = "1.0"
	// test server info
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.BasePath = ""
	// test schemes
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}

	r := setupRouter()

	r.Run()
}
