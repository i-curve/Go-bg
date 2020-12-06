package diary_service

import "template/model"

type Diary struct {
	ID       int
	Username string
	Title    string
	Text     string
}

func (a *Diary) CreateDiary() bool {
	return model.CreateDiary(a.Username, a.Title, a.Text)
}
func (a *Diary) ModifyDiary() bool {
	return model.ModifiedDiary(a.ID, a.Title, a.Text)
}
func (a *Diary) GetDiarys(pageNum, pageSize int) ([]*model.Diary, error) {
	return model.GetDiarys(pageNum, pageSize, a.getMaps())
}

func (a *Diary) GetDiaryCount() (int, error) {
	return model.GetDiaryCount(a.getMaps())
}
func (a *Diary) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	// maps["deleted_on"] = 0
	maps["username"] = a.Username
	return maps
}
