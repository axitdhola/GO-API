// // package main

// // import "testing"

// // func Testcalculate(t *testing.T){
// // 	if func1(2,4)!=6 {
// //        t.Error("Expected value is 6")
// // 	}
// // }

// package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAPI(t *testing.T) {
// 	// Set up a test gin engine and a test MySQL database
// 	gin.SetMode(gin.TestMode)
// 	router := gin.New()
// 	db, _ := gorm.Open("mysql", "root:root@tcp(localhost:3306)/Config?charset=utf8&parseTime=True&loc=Local")

// 	// Define a test route
// 	router.GET("/test", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Route is Created!"})
// 	})

// 	// Create a request and response recorder
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/test", nil)

// 	// Serve the request and check the response
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, `{"message":"Route is Created!"}`, w.Body.String())

// 	// Close the test MySQL database
// 	db.Close()
// }

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPerson(t *testing.T) {
	r := gin.Default()
	r.GET("/Person", GetPerson)
	req, _ := http.NewRequest("GET", "/Person", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var persons []Person
	json.Unmarshal(w.Body.Bytes(), &persons)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, persons)
}

func TestUpdatePerson(t *testing.T) {
	r := gin.Default()
	r.PUT("/Person/:id", UpdatePerson)
	newPerson := Person{
		ID:   144,
		Name: "Bhaskar",
	}
	jsonValue, _ := json.Marshal(newPerson)
	reqFound, _ := http.NewRequest("PUT", "/Person/"+strconv.Itoa(newPerson.ID), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/Person/12", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
