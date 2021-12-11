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

func (u *User) GetUserByIp(ip string) (user User, err error) {
	user.Ip = ip
	result := dao.DB.First(&user)
	if result.Error != nil {
		return user, err
	}
	return user, nil
}
