package entity

import "time"

type Kurir struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	phone      string    `gorm:"type:varchar(255)" json:"phone"`
	noKtp      string    `gorm:"type:varchar(255)" json:"noKtp"`
	fotoKtp    string    `gorm:"type:varchar(255)" json:"fotoKtp"`
	alamat     string    `gorm:"uniqueIndex;type:varchar(255)" json:"alamat"`
	birthdate  string    `gorm:"type:varchar(255)" json:"birthdate"`
	password   string    `gorm:"->;<-;not null" json:"-"`
	Token      string    `gorm:"-" json:"token,omitempty"`
	created_at time.Time `gorm:"<-:create"`
	updated_at time.Time `gorm:"<-:update"`
}
