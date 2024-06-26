package models

type Project struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Link        string `json:"link"`
}