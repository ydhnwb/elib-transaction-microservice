package entity

//Transaction represents transactions table in database
type Transaction struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	StudentID uint64  `gorm:"not null" json:"-"`
	Student   Student `json:"student" gorm:"foreignkey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookID    uint64  `gorm:"not null" json:"-"`
	Book      Book    `json:"book" gorm:"foreignkey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
