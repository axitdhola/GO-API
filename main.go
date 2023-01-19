package main

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct{
	ID    int `json:"id"`
	Name  string `json:"name"`
} 

func resultQuery(query string) ([]Person , bool){

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/Config")
	if err != nil {
		panic(err.Error())
		return nil , false
	}
	defer db.Close()

	queryResult, err := db.Query(query)
    if err != nil {
		panic(err.Error())
		return nil , false
	}
	defer queryResult.Close()
    
	var result []Person
	for queryResult.Next() {
        var newPerson Person
        err = queryResult.Scan(&newPerson.ID,&newPerson.Name)
        if err != nil {
			panic(err.Error())
			return nil , false
		}
        result = append(result,newPerson)
    }
	return result , true 
}

func getPerson(context *gin.Context){  // Context contains data
	   
	result , flag := resultQuery("SELECT * FROM demo1")
	if(!flag) {
		return
	}
    context.IndentedJSON(http.StatusOK , result)
}

func addPerson(context *gin.Context){
	var newPerson Person
	err := context.BindJSON(&newPerson)
	if(err != nil) {
		return
	}
	
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/Config")
	if err != nil {
		return
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO demo1 VALUES (?,?)" , newPerson.ID , newPerson.Name)
	defer insert.Close()
	context.IndentedJSON(http.StatusCreated , newPerson)  // (status , JSON)
}

func deletePerson(context *gin.Context){
	id := context.Param("id")
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/Config")
	if err != nil {
		panic(err.Error())
	    return
	}
	defer db.Close()

	res, err:= db.Query("DELETE FROM demo1 WHERE pid = ?" , id)
	defer res.Close()
	if err != nil {
		return
	}
	
	context.IndentedJSON(http.StatusOK, "completed")
}

func updatePerson(context *gin.Context){
	var newPerson Person
	err := context.BindJSON(&newPerson)
	if(err != nil) {
		return
	}

	id := context.Param("id")
    db, err := sql.Open("mysql" , "root:root@tcp(localhost:3306)/Config")
	if(err!=nil){
		panic(err.Error())
		return
	}

	defer db.Close()

	res, err := db.Query("UPDATE demo1 SET Name = ? WHERE pid = ?" , newPerson.Name , id)
	defer res.Close()
    if(err!=nil){
		panic(err.Error())
		return
	}
    
	context.IndentedJSON(http.StatusOK, res)
}

func main(){
	router := gin.Default() // create server 
	router.GET("/Person" , getPerson)
	router.POST("/Person" , addPerson)
	router.PUT("/Person/:id" , updatePerson)
	router.DELETE("/Person/:id" , deletePerson) // : represent dynamic
	router.Run("localhost:9090") // path endpoint
}

