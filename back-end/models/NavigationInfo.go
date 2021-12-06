package models

import (
	"Navigation-Web/dao"
	"Navigation-Web/utils"
	"time"
)

type NavigationInfo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Logo      string    `json:"logo"`
	CreatedAt utils.JsonTime `json:"created_time"`
	UpdatedAt utils.JsonTime `json:"update_time"`
}


func GetAll() (navigationInfos []NavigationInfo, err error) {
	result := dao.DB.Find(&navigationInfos)
	if result.Error != nil {
		return nil, err
	}
	return navigationInfos, nil
}

func GetOne(id int) (navigationInfo *NavigationInfo, err error) {
	result := dao.DB.First(&navigationInfo, id)
	if result.Error != nil {
		return navigationInfo, err
	}
	return navigationInfo, nil
}

func Create(info *NavigationInfo) (err error) {
	info.CreatedAt=utils.JsonTime(time.Now())
	info.UpdatedAt=utils.JsonTime(time.Now())
	result := dao.DB.Create(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func Update(info *NavigationInfo) (err error) {
	info.UpdatedAt=utils.JsonTime(time.Now())
	result := dao.DB.Save(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func Delete(info *NavigationInfo) (err error) {
	result := dao.DB.Delete(info)
	if result.Error != nil {
		return err
	}
	return nil
}
