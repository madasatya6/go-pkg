package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"

	"github.com/madasatya6/go-pkg/validation"
	"github.com/madasatya6/go-pkg/response"
)

const port string = "9090"

type User struct{
	Age 	int 	`json:"age" form:"age" validate:"required,numeric"`
	Name 	string 	`json:"name" form:"name" validate:"required"`
	Email 	string 	`json:"email" form:"email" validate:"required,email"`
}

func ValidationErrors(w http.ResponseWriter, r *http.Request){
	
	if r.Method == "POST" {
		
		var user User
		// r.ParseForm()

		user.Age, _ = strconv.Atoi(r.FormValue("age"))
		user.Name = r.FormValue("name")
		user.Email = r.FormValue("email")

		if err := validation.FormErrorEN(user); err != nil {
			var data = map[string]interface{}{
				"status": false,
				"message": "Terkena validasi",
				"data": map[string]interface{}{
					"user": user,
				},
				"validated": err,
			}
			response.JSON(w, data, http.StatusBadRequest)
			return 
		}

		var data = map[string]interface{}{
			"status": true,
			"message": "Selamat tidak terkena validasi!",
			"data": user,
			"validated": nil,
		}
		response.JSON(w, data, http.StatusOK)
		return
	} 
	
	http.Error(w, "Tidak diizinkan.", http.StatusNotFound)
	return
}

func main(){
	
	http.HandleFunc("/user", ValidationErrors)
	
	err := http.ListenAndServe(":"+port, nil)
	
	fmt.Println("Berjalan pada port :", port)
	
	if err != nil {
		log.Fatal(err.Error())
	}

}