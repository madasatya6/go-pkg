package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/madasatya6/go-pkg/pagination"
	"github.com/madasatya6/go-pkg/response"
)

const port string = "9090"

type UserList struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	Email string `json:"email"`
}

func Pagination(w http.ResponseWriter, r *http.Request){
	
	if r.Method == "GET" {
		
		/**
		* untuk menangkap page keberapa gunakan query param p
		* misal: http://localhost:9090/example/pagination?p=2
		*/
		var Data = []interface{}{
			UserList{ID:1, Nama:"Mada", Email: "madasatya6@gmail.com"},
			UserList{ID:2, Nama:"Seno", Email: "seno@gmail.com"},
			UserList{ID:3, Nama:"Pandu", Email: "pandu@gmail.com"},
			UserList{ID:4, Nama:"Angga", Email: "angga@gmail.com"},
			UserList{ID:5, Nama:"Putra", Email: "putra@gmail.com"},
			UserList{ID:6, Nama:"Agus", Email: "agus@gmail.com"},
			UserList{ID:7, Nama:"Joko", Email: "joko@gmail.com"},
			UserList{ID:8, Nama:"Edi", Email: "edi@gmail.com"},
			UserList{ID:9, Nama:"Sutris", Email: "sutris@gmail.com"},
			UserList{ID:10, Nama:"Deni", Email: "deni@gmail.com"},
			UserList{ID:11, Nama:"Capunk", Email: "capunk@gmail.com"},
			UserList{ID:12, Nama:"Ervin", Email: "ervin@gmail.com"},
		}

		perpage := 4
		page := pagination.Paginate(r, Data, perpage)

		var data = map[string]interface{}{
			"status": true,
			"message": "Latihan pagination!",
			"data": map[string]interface{}{
				"example": page["posts"],
			},
		}
		response.JSON(w, data, http.StatusOK)
		return
	} 
	
	http.Error(w, "Tidak diizinkan.", http.StatusNotFound)
	return
}

func main(){
	
	http.HandleFunc("/mahasiswa", Pagination)
	
	err := http.ListenAndServe(":"+port, nil)
	
	fmt.Println("Berjalan pada port :", port)
	
	if err != nil {
		log.Fatal(err.Error())
	}

}