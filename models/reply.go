// models/reply.go

package models

type Reply struct {
	BaseModel
	Text           string `json:"text" binding:"required"`
	Reported       string `json:"reported" gorm:"default:false"`
	DeletePassword string `json:"delete_password" gorm:"not null" binding:"required`
	ThreadID       uint   `json:"thread_id"`
}
