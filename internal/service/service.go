package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	API "github.com/strong-defi/strong-defi-backend/common"
	plog "github.com/strong-defi/strong-defi-backend/pkg/log"
	"net/http"
)

var (
	log plog.Logger
)

type Service struct {
	EtherService *EtherService
	StakeService *StakeService
	UserService  *UserService
}

func New(logger plog.Logger,
	etherService *EtherService,
	stakeService *StakeService,
	userService *UserService) *Service {
	log = logger.With("module", "service")
	return &Service{
		EtherService: etherService,
		StakeService: stakeService,
		UserService:  userService,
	}

}

type Context struct {
	*gin.Context
}

// Response /*封装返回信息*/
type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

// JSON /*代理原有的Json方法*/
func (c *Context) JSON(msg API.ApiResponseEnum, obj interface{}) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(msg),
		Data: obj,
		Msg:  API.GetMessage(msg),
	}
	c.Context.JSON(ok, response)
}

// JSON2 /*代理原有的Json方法*/
func (c *Context) JSON2(api API.ApiResponseEnum, msg string) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(api),
		Data: "",
		Msg:  msg,
	}
	c.Context.JSON(ok, response)
}

// CustomJSON 定制 JSON
func (c *Context) CustomJSON(msg API.ApiResponseEnum, desc string) {
	ok := http.StatusOK
	response := &Response{
		Code: API.GetCode(msg),
		Data: "",
		Msg:  desc,
	}
	c.Context.JSON(ok, response)
}

// CustomJSON2 数据校验
func dataValidate(obj interface{}) error {
	validate := validator.New()
	return validate.Struct(obj)
}
