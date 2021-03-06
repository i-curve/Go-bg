package model

import (
	"fmt"
)

type Diary struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int
	ModifiedOn int
	Username   string `json:"username"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

// func GetDiaryById(id int) {
// 	var diary
// }
func CreateDiary(username, title, text string) bool {
	var diary = &Diary{
		Username: username,
		Title:    title,
		Text:     text,
	}
	err := db.Create(&diary).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func ModifiedDiary(id int, title, text string) bool {
	// var diary = &Diary{
	// 	Title: title,
	// 	Text:  text,
	// }
	err := db.Model(&Diary{}).Where("id = ?", id).Update(map[string]interface{}{"title": title, "text": text})
	if err.Error != nil {
		return false
	}
	return true
}
func GetDiarys(pageNum, pageSize int, maps interface{}) ([]*Diary, error) {
	var diarys []*Diary
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&diarys).Error
	if err != nil {
		return nil, err
	}
	return diarys, nil
}
func GetDiaryCount(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Diary{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
