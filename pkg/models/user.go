package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"     binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type Dashboard struct {
	QuizID int `json:"quiz,omitempty"  binding:"required" gorm:"primaryKey"`
	UserID int `json:"user,omitempty"  binding:"required"`
	Score  int `json:"score,omitempty"`
}
