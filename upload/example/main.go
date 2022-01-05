package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"

	"github.com/madasatya6/go-pkg/upload"
	"github.com/madasatya6/go-pkg/response"
)

const port string = "9090"

func tampilForm(w http.ResponseWriter, r *http.Request){
	
	var tmpl = template.Must(template.New("form").
		ParseFiles("form.html"))
		
	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func upload_file (w http.ResponseWriter, r *http.Request){
	
	if r.Method == "POST" {
		
		var name = r.FormValue("filename")
		
		fileName, err := upload.FileValidate(r, "image", []string{"png","jpg","jpeg","bmp"}, 1, 1000, true)
		if err != nil {
			
			var data = map[string]interface{}{
				"status": false,
				"message": err.Error(),
				"data": nil,
			}
			
			response.JSON(w, data, http.StatusBadRequest)
			return
		}

		var data = map[string]interface{}{
			"status": true,
			"message": "Data berhasil diupload!",
			"data": map[string]interface{}{
				"filename": fileName,
			},
		}
		response.JSON(w, data, http.StatusOK)
		return
	} 
	
	http.Error(w, "Tidak diizinkan.", http.StatusNotFound)
	return
}

func main(){
	
	//load upload file asset
	http.Handle("/upload/", http.StripPrefix("/upload/",
		http.FileServer(http.Dir("upload"))))
	
	http.HandleFunc("/", tampilForm)
	
	http.Handle("/upload", validasi_upload(http.HandlerFunc(upload_file)))
	
	err := http.ListenAndServe(":"+port, nil)
	
	fmt.Println("Berjalan pada port :", port)
	
	if err != nil {
		log.Fatal(err.Error())
	}

}