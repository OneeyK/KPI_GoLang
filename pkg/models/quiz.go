package models

type Quiz struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"      binding:"required"`
	Questions []Question `json:"questions,omitempty"`
}

type Question struct {
	ID          int      `json:"id,omitempty"`
	QuizID      int      `json:"quiz_id,omitempty"`
	Description string   `json:"description,omitempty" binding:"required"`
	Options     []Option `json:"options,omitempty"`
}

type Option struct {
	ID         int    `json:"id,omitempty"`
	QuestionID int    `json:"question_id,omitempty"`
	Content    string `json:"content,omitempty"     binding:"required"`
	Correct    bool   `json:"correct"`
}

type Answer struct {
	QuestionID int `json:"question_id" binding:"required"`
	OptionID   int `json:"option_id"   binding:"required"`
}
