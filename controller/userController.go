package controller

import (
	"log"
	"time"
	"varcelio/model"
	"varcelio/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router  *gin.RouterGroup
	service *service.UserService
}

func NewUserController(router *gin.RouterGroup) *UserController {
	uc := &UserController{router: router, service: service.NewUserService()}

	user := uc.router.Group("/dev/user")
	user.POST("/add", uc.AddUser)
	user.GET("/get-one", uc.GetOneUser)

	return uc
}

// @Tags user
// @Accept json
// @Param parameter body model.UserModel true "PARAM"
// @Produce json
// @Success 201 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /v1/dev/user/add [post]
func (uc *UserController) AddUser(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	var param model.UserModel
	if err := ctx.BindJSON(&param); err != nil {
		log.Println(err)
		return
	}
	resp = uc.service.AddUser(param)
}

// @Tags user
// @Accept json
// @Param id query string true "ID"
// @Produce json
// @Success 200 {object} object{meta_data=model.MetadataResponse} "OK"
// @Router /v1/dev/user/get-one [get]
// @Security JWT
func (uc *UserController) GetOneUser(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Data = uc.service.GetOneUser("_id", ctx.Query("id"))
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	resp := model.Response{}
	defer SetMetadataResponse(ctx, time.Now(), &resp)

	resp.Data = uc.service.GetAll()
}