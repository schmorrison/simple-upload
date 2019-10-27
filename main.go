package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("File upload requested from\t%s\n %s\t%s\t(%d)\n", r.RemoteAddr, r.Method, r.URL.String(), r.ContentLength)

	r.ParseMultipartForm(10 << 20) // 10MB
	file, handler, err := r.FormFile("file")
	if err != nil {
		err = fmt.Errorf("Failed to retrieve the uploaded form file: %s", err)
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Retreived uploaded form file:\nHeader:\t%+v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Failed to read form file [%s]: %s", handler.Filename, err)
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(handler.Filename, fileBytes, 0700)
	if err != nil {
		err = fmt.Errorf("Failed to write the uploaded data to the temp file [%s]: %s", handler.Filename, err)
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"Filename": handler.Filename,
		"Size":     handler.Size,
	})

	fmt.Println("File was successfully written to the present working directory: ")
	fmt.Printf("\tName:\t%s\n\tSize:\t%d\n", handler.Filename, handler.Size)
}

// func downloadFile(w http.ResponseWriter, r *http.Request) {
// 	// Get filename from URL query params
// 	path := r.URL.Query().Get("filepath")

// 	b, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		err = fmt.Errorf("Failed to locate file '%s': %s", path, err)
// 		fmt.Println(err)

// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path))
// 	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
// 	w.Header().Set("Content-Length", fmt.Sprint(binary.Size(b)))
// 	w.Write(b)
// }

func main() {
	fmt.Println("Starting Server.... Port : 8765")
	// Upload file endpoint
	http.HandleFunc("/upload/", uploadFile)
	// http.HandleFunc("/download/", downloadFile)

	// Static resources endpoint
	http.Handle("/files/", http.StripPrefix(strings.TrimRight("/files/", "/"), http.FileServer(http.Dir("./"))))

	// Template endpoint
	http.HandleFunc("/", indexPage)

	http.ListenAndServe("0.0.0.0:8765", nil)
}
