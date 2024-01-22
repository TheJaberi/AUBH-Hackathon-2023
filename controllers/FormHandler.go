package AMJ

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	// "github.com/otiai10/gosseract/v2"
)

// FormHandler handles HTTP requests for the form submission.
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseMultipartForm(100 << 20) // 100 MB max memory
	if err != nil {
		fmt.Println("Error Parsing the Multipart Form:", err)
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}

	// Map to store the paths of the uploaded images
	uploadedFiles := make(map[string]string)

	// Process the uploaded files
	for _, inputName := range []string{"paperUpload", "keyUpload"} {
		file, handler, err := r.FormFile(inputName)
		if err != nil {
			fmt.Printf("Error Retrieving the File for input %s: %v\n", inputName, err)
			continue // Continue to process the next file
		}
		defer file.Close()

		// Create the img directory if it doesn't exist
		imgDir := filepath.Join(".", "img")
		os.MkdirAll(imgDir, os.ModePerm) // Create all directories along the path if necessary

		// Create a new file in the img directory
		dstFilePath := filepath.Join(imgDir, handler.Filename)
		dst, err := os.Create(dstFilePath)
		if err != nil {
			fmt.Printf("Unable to create the file for writing: %v\n", err)
			continue // Continue to process the next file
		}
		fmt.Println(handler.Filename)
		defer dst.Close()

		// Copy the uploaded file to the filesystem at the specified destination
		if _, err := io.Copy(dst, file); err != nil {
			fmt.Printf("Unable to write the file: %v\n", err)
			continue // Continue to process the next file
		}

		// Store the path of the uploaded file
		uploadedFiles[inputName] = dstFilePath
	}

	// // Perform OCR using gosseract on the uploaded images
	// for key, filePath := range uploadedFiles {
	// 	client := gosseract.NewClient()
	// 	defer client.Close()
	// 	if err := client.SetImage(filePath); err != nil {
	// 		fmt.Printf("Error setting image for OCR: %v\n", err)
	// 		continue // Continue with the next image
	// 	}
	// 	text, err := client.Text()
	// 	if err != nil {
	// 		fmt.Printf("Error performing OCR on image %s: %v\n", key, err)
	// 		continue // Continue with the next image
	// 	}
	// 	fmt.Printf("Extracted Text for %s: %s\n", key, text)
	// 	// Do something with the extracted text, like storing it in a database or file
	// }

	// Render the template or redirect to a success page
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
