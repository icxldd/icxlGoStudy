package Requests

import (
	"gorm.io/gorm"
	"icxl/Entitys"
)

type AccountDTO struct{
	gorm.Model
	Name string `form:"Name"`
	PassWord string `form:"PassWord"`
}

type POSTAccountDTO struct {
	AccountDTO
}

type PUTAccountDTO struct {
	AccountDTO
}

type DELETEAccountDTO struct {
	gorm.Model
}

func AccountDTOToEntity(AccountDTO *AccountDTO)  *Entitys.Account{
	var account *Entitys.Account=new(Entitys.Account)
	//var account *Entitys.Account=&Entitys.Account{}
	account.ID = AccountDTO.ID
	account.PassWord = AccountDTO.PassWord
	account.Name = AccountDTO.Name
	return account
}






