package models

import (
	"Navigation-Web/dao"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `json:"id"`
	Ip          string `json:"ip"`
	BrowserInfo string `json:"browser_info"`
}

//用戶模型中暫時只会有两种操作 创建和查询

func (u *User) Created(user *User) (err error) {
	if dao.DB.Create(&user).Error != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByIp(ip string) (user *User, err error) {
	if dao.DB.Where("ip =?", ip).Find(&user).Error != nil {
		return nil, err
	}
	return user, nil
}
