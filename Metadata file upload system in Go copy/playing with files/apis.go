package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getData(GET *gin.Context) {
	rows, err := db.Query("SELECT file_name, file_data, upload_data, size, file_type")
	if err != nil {
		GET.JSON(http.StatusInternalServerError, gin.H{"Error ": err.Error()})
		return
	}
	defer rows.Close()

	var data []files

	for rows.Next() {
		var dt files
		if err := rows.Scan(&dt.File_name, &dt.File_data, &dt.Upload_Date, &dt.Size, &dt.File_type); err != nil {
			GET.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		data = append(data, dt)
	}

	GET.JSON(http.StatusOK, data)
}

func addFiles(post *gin.Context) {
	// we get the file
	file, _, err := post.Request.FormFile("file")
	if err != nil {
		post.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// we read the file
	fileData, err := io.ReadAll(file)
	if err != nil {
		post.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	// getting metadata

	fileName := post.Request.FormValue("file_name")
	fileType := post.Request.FormValue("file_type")
	uploadDate := time.Now()
	fileSize := int64(len(fileData))

	db, err := connect()
	if err != nil {
		post.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO files (file_name, file_data, upload_date, size, file_type) VALUES (?, ?, ?, ?, ?)",
		fileName, fileData, uploadDate, fileSize, fileType)
	if err != nil {
		post.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	post.JSON(http.StatusOK, "data added")

}
