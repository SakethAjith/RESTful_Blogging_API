package models

type Blogs struct {
	Id      int64  `json:"id" gorm:"autoIncrement;primaryKey"`
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content"`
}
