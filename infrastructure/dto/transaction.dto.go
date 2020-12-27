package dto

//TransactionCreateDTO used when a new transcation need to be created
type TransactionCreateDTO struct {
	StudentID uint64 `form:"student_id" json:"student_id" binding:"required"`
	BookID    uint64 `form:"book_id" json:"book_id" binding:"required"`
}
