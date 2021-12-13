package models

import (
	"Navigation-Web/dao"
	"Navigation-Web/models/dto"
	"encoding/base64"
	"strings"
	"time"
)

type Note struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Title        string    `json:"title"`
	Tags         string    `json:"tags"` //以逗号分隔开的数组
	Content      string    `json:"content"`
	Text         string    `json:"text"`
	EditStatus   uint      `json:"edit_status"`
	CreatedAt    time.Time `json:"create_time"`
	DeleteStatus uint      `json:"delete_status"`
	UpdateAt     time.Time `json:"update_time"`
}


func (n *Note) GetOne(id int) (note *Note, err error) {
	result := dao.DB.Where("delete_status != ?", 1).First(&note, id)
	content, err := base64.StdEncoding.DecodeString(note.Content)
	note.Content = string(content)
	if result.Error != nil {
		return note, err
	}
	return note, nil
}

func (note *Note) GetAll() (notes []Note, err error) {
	result := dao.DB.Where("delete_status != ?", 1).Find(&notes)
	if err = result.Error; err != nil {
		return nil, err
	}
	return notes, nil
}
func(n *Note) GetListByQueryDto(noteDto *dto.NoteQueryDto) (notes []Note,err error){
	result := dao.DB.Where(&Note{DeleteStatus: 0,EditStatus: 1}).Find(&notes)
	if err = result.Error; err != nil {
		return nil, err
	}
	return notes,nil
}

func (n *Note) Create(note *Note) (err error) {
	note.CreatedAt = time.Now()
	note.UpdateAt = time.Now()
	note.Text=strings.ReplaceAll(note.Text,"&nbsp;","")
	note.Content = base64.StdEncoding.EncodeToString([]byte(note.Content))
	if err = dao.DB.Create(&note).Error; err != nil {
		return err
	}
	return nil
}

func (n *Note) Update(note *Note) (err error) {
	note.UpdateAt = time.Now()
	note.Content = base64.StdEncoding.EncodeToString([]byte(note.Content))
	oldStr:="&nbsp;"
	note.Text=strings.ReplaceAll(note.Text,oldStr,"")
	if err = dao.DB.Save(&note).Error; err != nil {
		return err
	}
	return nil
}

// func (n *Note) Delete(note *Note) (err error) {
// 	var deleteNote Note
// 	deleteNote.ID=note.ID
// 	deleteNote.DeleteAt=time.Now()
// 	deleteNote.DeleteStatus=1
// 	if err = dao.DB.Save(&deleteNote).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
