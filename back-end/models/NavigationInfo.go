package models

import (
	"Navigation-Web/dao"
	"time"
)

type NavigationInfo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Logo      string    `json:"logo"`
	Desc      string    `json:"desc"`
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"update_time"`
}

var NavigationInfofunc = new(NavigationInfo)

func (n *NavigationInfo) GetAll() (navigationInfos []NavigationInfo, err error) {
	result := dao.DB.Order("weight desc").Find(&navigationInfos)
	if result.Error != nil {
		return nil, err
	}
	return navigationInfos, nil
}

func (n *NavigationInfo) GetOne(id int) (navigationInfo *NavigationInfo, err error) {
	result := dao.DB.First(&navigationInfo, id)
	if result.Error != nil {
		return navigationInfo, err
	}
	return navigationInfo, nil
}

func (n *NavigationInfo) Create(info *NavigationInfo) (err error) {
	info.CreatedAt = time.Now()
	info.UpdatedAt = time.Now()
	result := dao.DB.Create(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func (n *NavigationInfo) Update(info *NavigationInfo) (err error) {
	info.UpdatedAt = time.Now()
	result := dao.DB.Save(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func (n *NavigationInfo) Delete(info *NavigationInfo) (err error) {
	result := dao.DB.Delete(info)
	if result.Error != nil {
		return err
	}
	return nil
}
