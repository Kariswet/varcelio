package service

import (
	"log"
	"varcelio/model"
	"varcelio/util"

	"github.com/google/uuid"
)

type UserService struct {
	util *util.MongoUtil
}

func NewUserService() *UserService {
	return &UserService{
		util: util.NewMongoUtil(),
	}
}

func (us *UserService) AddUser(param model.UserModel) (resp model.Response) {
	param.Id = uuid.NewString()
	param.Password = util.HashPassword(param.Password)

	newId, _ := us.util.InsertOne(param)

	resp.Data = newId

	return
}

func (us *UserService) GetOneUser(key, value string) (resp model.UserModel) {
	if err := us.util.FindOne(key, value, &resp); err != nil {
		log.Println(err)
		return
	}
	return
}

func (us *UserService) GetAll() (resp []model.UserModel) {
	// if err := us.util.Find(key, value, &resp); err != nil {
	// log.Println(err)
	// return
	// }
	return

}
