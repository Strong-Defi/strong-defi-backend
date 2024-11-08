package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	API "strong-defi-backend/common"
	"strong-defi-backend/config"
	"strong-defi-backend/model"
)

var (
	dao *model.Dao
	app config.App
)

func New(d *model.Dao) {
	dao = d
	app = config.Config.App

}

func NewDao(orm *gorm.DB) *model.Dao {
	d := &model.Dao{Orm: orm}
	return d
}

type CustomContext struct {
	*gin.Context
}

/*封装返回信息*/
type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

/*代理原有的Json方法*/
func (c *CustomContext) JSON(msg API.ApiResponseEnum, obj interface{}) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(msg),
		Data: obj,
		Msg:  API.GetMessage(msg),
	}
	c.Context.JSON(ok, response)
}
func (c *CustomContext) JSON2(api API.ApiResponseEnum, msg string) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(api),
		Data: "",
		Msg:  msg,
	}
	c.Context.JSON(ok, response)
}
func (c *CustomContext) CustomJSON(msg API.ApiResponseEnum, desc string) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(msg),
		Data: "",
		Msg:  desc,
	}
	c.Context.JSON(ok, response)
}

func dataValidate(obj interface{}) error {
	validate := validator.New()
	return validate.Struct(obj)
}
