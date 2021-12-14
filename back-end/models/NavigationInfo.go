package models

import (
	"Navigation-Web/dao"
	"Navigation-Web/models/dto"
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

func (n *NavigationInfo) GetAll(pageinfo *dto.PageInfo) (page *dto.PageInfo,err error) {
	var count int64
	if err = dao.DB.Model(&NavigationInfo{}).Count(&count).Error; err != nil {
		return nil, err
	}
	var navigationInfos [] NavigationInfo
	result:=pageinfo.GetPageDB().Order("weight desc").Find(&navigationInfos)
	if result.Error != nil {
		return nil, err
	}
	pageinfo.Total=count
	pageinfo.Data=navigationInfos
	return pageinfo, nil
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
