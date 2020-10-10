package Services

import (
	"github.com/gin-gonic/gin"
	. "icxl/DTO/Requests"
	. "icxl/Dao"
	"icxl/Entitys"
	"unsafe"
)

func GetStudents(context *gin.Context)  {


	var accounts []Entitys.Account

	DB.Find(&accounts)

	context.IndentedJSON(200,accounts)
}

func PutStudents(context *gin.Context)  {
	var requestDTO PUTAccountDTO
	context.Bind(&requestDTO)
	account:= AccountDTOToEntity((*AccountDTO)(unsafe.Pointer(&requestDTO)))
	DB.Save(&account)
	context.IndentedJSON(200,account)
}
func PostStudents(context *gin.Context)  {
	var requestDTO POSTAccountDTO
	context.Bind(&requestDTO)
	account:= AccountDTOToEntity((*AccountDTO)(unsafe.Pointer(&requestDTO)))

	DB.Create(&account)
	context.IndentedJSON(200,account)
}
func DeleteStudents(context *gin.Context)  {
	var requestDTO DELETEAccountDTO
	context.Bind(&requestDTO)
	account:= AccountDTOToEntity((*AccountDTO)(unsafe.Pointer(&requestDTO)))

	DB.Delete(&account)
	context.JSON(200,"ok")
}
