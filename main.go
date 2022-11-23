package main

import "github.com/gin-gonic/gin"
import "net/http"

// Binding from JSON with validation
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "sunil", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if json.User != "sunil" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 
		
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})


	// localhost:8080/user/sunil
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/ping", func(c *gin.Context) { 

		var p Person // Person Object

		p.Name = "Sunil Kumar"; // Note that p.Name becomes "name" in the JSON
		p.Age = 24;

		c.Bind(&p)
         
		// c.JSON() return output in json format
		c.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
			"name": p.Name,
			"Age":p.Age,
		})
	})


	router.Run()
}

// -------------- GET -------------------

// curl --location --request GET 'localhost:8080/user/sunil' \
// --data-raw ''




// ------------------------- POST -----------------------------------

// $ curl -v -X POST \
//   http://localhost:8080/loginJSON \
//   -H 'content-type: application/json' \
//   -d '{ "user": "manu" }'
// > POST /loginJSON HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.51.0
// > Accept: */*
// > content-type: application/json
// > Content-Length: 18
// >
// * upload completely sent off: 18 out of 18 bytes
// < HTTP/1.1 400 Bad Request
// < Content-Type: application/json; charset=utf-8
// < Date: Fri, 04 Aug 2017 03:51:31 GMT
// < Content-Length: 100
// <
// {"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
// Skip validate