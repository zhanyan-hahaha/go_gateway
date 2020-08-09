package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/zhaojiasanxing/go_gateway/dto"
	"github.com/zhaojiasanxing/go_gateway/public"
	"time"
)

type Admin struct {
	Id int `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName string `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt string `json:"salt" gorm:"column:salt" description:"盐" `
	Password string `json:"password" gorm:"column:password" description:"密码"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"created_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *Admin) TableName() string {
	return "gateway_admin"
}

func (t *Admin)LoginCheck(tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := t.Find(tx,(&Admin{UserName:param.UserName, IsDelete:0}))
	if err != nil{
		return nil, err
	}
	//param.Password
	//adminIndo.Salt
	saltPassword := public.GinSaltPassword(adminInfo.Salt, param.Password)
	if saltPassword != adminInfo.Password{
		return nil, errors.New("密码错误")
	}
	return adminInfo,nil
}

func (t *Admin)Find(tx *gorm.DB, search *Admin) (*Admin, error) {
	admin := &Admin{}
	err := tx.Where(search).Find(admin).Error
	if err != nil{
		return nil,err
	}
	return admin, nil
}

func (t *Admin)Save(tx *gorm.DB) error {
	err := tx.Save(t).Error
	return err
}
