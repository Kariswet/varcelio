package util

import (
	"errors"
	"log"
	"os"
	"varcelio/config"
	"varcelio/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MongoUtil struct {
	Config *config.MongoConfig
}

func NewMongoUtil() *MongoUtil {
	return &MongoUtil{
		Config: config.NewMongoConfig(os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_DB_COLLECTION")),
	}
}

func (mu *MongoUtil) InsertOne(param interface{}) (newId any, err error) {
	client, err := mu.Config.Connect()
	if err != nil {
		log.Println("client failed")
		return
	}
	defer mu.Config.Disconnect(client)

	coll := client.Database(mu.Config.DbName).Collection(mu.Config.Collection)
	result, err := coll.InsertOne(mu.Config.Ctx, param)
	log.Println(result)
	if err != nil {
		log.Println("insert failed")
		return
	}

	return result.InsertedID, nil
}

func (mu *MongoUtil) BaseFindOne(filter bson.M, param interface{}) (err error) {
	client, err := mu.Config.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer mu.Config.Disconnect(client)

	coll := client.Database(mu.Config.DbName).Collection(mu.Config.Collection)

	result := coll.FindOne(mu.Config.Ctx, filter)
	if result.Err() != nil {
		log.Println(result.Err(), filter)
		err = errors.New("data not found")
		return
	}

	if err := result.Decode(param); err != nil {
		log.Println(err)
	}
	return
}

func (mu *MongoUtil) FindOne(key, value string, param interface{}) (err error) {
	return mu.BaseFindOne(bson.M{key: value}, param)
}

func (mu *MongoUtil) BaseFind(filter bson.M, data interface{}) (resp model.Response, err error) {
	client, err := mu.Config.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer mu.Config.Disconnect(client)

	coll := client.Database(mu.Config.DbName).Collection(mu.Config.Collection)
	cursor, err := coll.Find(mu.Config.Ctx, filter)
	if err != nil {
		log.Println(err)
		return
	}

	if err = cursor.All(mu.Config.Ctx, data); err != nil {
		log.Println(err)
		return
	}

	return
}

func (mu *MongoUtil) Find(param model.User_Search) (err error) {
	// filter := bson.M{}

	// _, err = mu.BaseFind(filter, )
	return
}

// func (mu *MongoUtil) BaseGetAll(param model.User_Search) (data []model.UserModel, resp model.MetadataResponse) {
// 	client, err := mu.Config.Connect()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer mu.Config.Disconnect(client)
// 	coll := client.Database(mu.Config.DbName).Collection(mu.Config.Collection)

// 	filter := bson.M{}
// 	listFilterAnd := []bson.M{}
// 	param.HandleFilter(&listFilterAnd)

// 	if len(listFilterAnd) > 0 {
// 		filter["$and"] = listFilterAnd
// 	}
// 	resp.Pagination, resp.Message = coll.Find()
// }
