package authenticate

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var V1Users = []TestModel{
	{
		ID:   1,
		Name: "user v1 - one",
	},
	{
		ID:   2,
		Name: "user v1 - two",
	},
}

var V1Products = []TestModel{
	{
		ID:   1,
		Name: "product v1 - one",
	},
	{
		ID:   2,
		Name: "product v1 - two",
	},
}

var V2Users = []TestModel{
	{
		ID:   1,
		Name: "user v2 - one",
	},
	{
		ID:   2,
		Name: "user v2 - two",
	},
}

var V2Products = []TestModel{
	{
		ID:   1,
		Name: "product v2 - one",
	},
	{
		ID:   2,
		Name: "product v2 - two",
	},
}

func authenticateMiddleware(c *gin.Context) {
	authToken := c.Request.Header.Get("auth-token")
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No token",
		})
		return
	}
	if authToken != "secret-token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	if len(c.Keys) == 0 {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["received-token"] = authToken
	log.Println("authenticateMiddleware passing")
	c.Next()
	log.Println("authenticateMiddleware passed already")
}

func ageCheckMiddleware(c *gin.Context) {
	userAge := c.Request.Header.Get("age")
	if userAge == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No age information",
		})
		return
	}
	age, err := strconv.Atoi(userAge)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Age info required",
		})
		return
	}
	if age < 18 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Age under 18 not allowed",
		})
		return
	}

	if len(c.Keys) == 0 {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["user-age"] = age
	log.Println("ageCheckMiddleware passing")
	c.Next()
	log.Println("ageCheckMiddleware passed already")
}

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	v2 := r.Group("v2")
	v1.Use(authenticateMiddleware)
	{
		user := v1.Group("user")
		user.Use(ageCheckMiddleware)
		{
			user.GET("", func(c *gin.Context) {
				log.Printf("received token: %v, user age: %v\n",
					c.Keys["received-token"], c.Keys["user-age"])
				c.JSON(http.StatusOK, gin.H{
					"data": V1Users,
				})
			})

		}
		product := v1.Group("product")
		{
			product.GET("", func(c *gin.Context) {
				log.Printf("received token: %v\n", c.Keys["received-token"])
				c.JSON(http.StatusOK, gin.H{
					"data": V1Products,
				})
			})
		}
	}

	v2.Use(authenticateMiddleware)
	{
		user := v2.Group("user")
		user.Use(ageCheckMiddleware)
		{
			user.GET("", func(c *gin.Context) {
				log.Printf("received token: %v, user age: %v\n",
					c.Keys["received-token"], c.Keys["user-age"])
				c.JSON(http.StatusOK, gin.H{
					"data": V2Users,
				})
			})

		}
		product := v2.Group("product")
		{
			product.GET("", func(c *gin.Context) {
				log.Printf("received token: %v\n", c.Keys["received-token"])
				c.JSON(http.StatusOK, gin.H{
					"data": V2Products,
				})
			})
		}
	}
	return r
}
