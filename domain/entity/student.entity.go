package entity

//Student entity represents students table in database
type Student struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	NIM  string `gorm:"type:varchar(255)" json:"nim"`
}
