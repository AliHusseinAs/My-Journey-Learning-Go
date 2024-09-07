package main

import "time"

type files struct {
	File_name   string    `json:"file_name"`
	File_data   []byte    `json:"file_data"`
	Upload_Date time.Time `json:"upload_date"`
	Size        int64     `json:"size"`
	File_type   string    `json:"file_type"`
}
