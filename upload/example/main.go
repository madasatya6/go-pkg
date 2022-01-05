package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"os" //upload file 
	"io" //upload file
	"path/filepath" //uploadfile
	"bytes" //validasi size image
)

const port string = "9090"
const MAX_UPLOAD_SIZE = 1024 * 1024 //1mb

func tampilForm(w http.ResponseWriter, r *http.Request){
	
	var tmpl = template.Must(template.New("form").
		ParseFiles("views/template-post/template-form.html"))
		
	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func upload_file (w http.ResponseWriter, r *http.Request){
	
	if r.Method == "POST" {
		
		var tmpl = template.Must(template.New("post").ParseFiles("views/template-post/template-post.html"))
		
		err := r.ParseMultipartForm(1024)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		
		//upload file begin
		
		uploadedFile, handler, err := r.FormFile("image")
		
		if err != nil {
			
			switch err {
				case http.ErrMissingFile: //file required
					http.Error(w, "File Dibutuhkan", http.StatusInternalServerError)
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			
			return
		}
		
		defer uploadedFile.Close()
		
		dir, err := os.Getwd()
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		filename := handler.Filename
		
		if nama != "" {
			//filename = fmt.Sprintf("%s%s", nama, filepath.Ext(handler.Filename))
			filename = fmt.Sprintf("%s-%s", nama, handler.Filename)
		}
		
		fileLocation := filepath.Join(dir, "upload", filename)
		targetLocation, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		defer targetLocation.Close()
		
		if _, err = io.Copy(targetLocation, uploadedFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		//end upload file
		
		data := map[string]interface{}{
			"Nama":nama,
			"Alamat":alamat,
			"Image":filename,
		}
		
		err = tmpl.Execute(w, data)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		return
	} 
	
	http.Error(w, "Tidak diizinkan.", http.StatusNotFound)
	return
}

func validasi_upload(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		/* Validasi Size 
		if r.ContentLength > MAX_UPLOAD_SIZE {
			http.Error(w, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
			return
		} */
		
		err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		uploadedFile, handler, err := r.FormFile("image")
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		defer uploadedFile.Close()
		
		//validasi file size
		var buff bytes.Buffer
		fileSize, err := buff.ReadFrom(uploadedFile)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		if fileSize > MAX_UPLOAD_SIZE {
			http.Error(w, "Ukuran gambar harus kurang dari 1MB", http.StatusBadRequest)
			return
		}
		
		//validasi ekstensi gambar
		var ekstensi_gambar = filepath.Ext(handler.Filename)
		fmt.Println("Ekstensi file :", ekstensi_gambar)
		if ekstensi_gambar != ".jpg" && ekstensi_gambar != ".png" {
			//http.Error(w, "Ekstensi yang dibolehkan JPG dan PNG.", http.StatusBadRequest)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		
		next.ServeHTTP(w, r)
		return
	})
}

func main(){

	fmt.Println("https://dasarpemrogramangolang.novalagung.com/B-12-form-value.html")
	
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