package entity

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(250);"`
	FileName    string `json:"file_name" gorm:"type:varchar(250);"`
	MimeType    string `json:"mime_type" gorm:"type:varchar(30);"`
	Collection  string `json:"collection" gorm:"type:varchar(50);"`
	StorageDisk string `json:"storage_disk" gorm:"type:varchar(50);"`
	Url         string `json:"url" gorm:"type:varchar(255);"`
}
