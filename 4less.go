package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// JSON
type User struct {
	Name  string `json:"name" validate:"required"`
	Login string `json:"login" validate:"email,required"`
}

// http
// стандартная библиотека
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello World") //записать
//}
//func main() {
//	http.HandleFunc("/hello", handler)
//	http.ListenAndServe(":8080", nil)
//}

// фреймворк GIN
// go get -u github.com/gin-gonic/gin
func handler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello GET")
}
func handlerPost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello POST")
}
func handlerPut(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello PUT")
}
func handlerName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusAccepted, "hello %s", name)
}
func handlerFLName(ctx *gin.Context) {
	firstName := ctx.DefaultQuery("first", "LOL")
	lastName := ctx.DefaultQuery("last", "lastLOL")
	ctx.String(http.StatusAccepted, "hello %s %s", firstName, lastName)
	//http://localhost:8080/welcome?first=Patric&last=Swag
}

func loginHandler(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //gin.H - как бы создаем мапу
		return
	}
	val := validator.New()
	if err := val.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ctx.String(http.StatusAccepted, "hello %s", user.Name)
	ctx.JSON(http.StatusAccepted, user) //ctx.JSON - в каком формате вернуть, есьт так же ctx.YAML
}
func main() {
	user := User{
		Name:  "huawei",
		Login: "huawei@mail.ru",
	}
	fmt.Println("user:", user)
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonStr))

	router := gin.Default()
	userGroup := router.Group("/hello")
	{
		userGroup.GET("/get", handler) // http://localhost:8080/hello/get
		userGroup.POST("/post", handlerPost)
		userGroup.PUT("/put", handlerPut)
	}
	router.GET("/hi/:name", handlerName)
	router.GET("/welcome", handlerFLName)
	router.POST("/login", loginHandler)
	router.Run(":8080")
}
