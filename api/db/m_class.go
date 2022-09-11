package db

import (
	"context"
	"fmt"

	"time"
	//"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	config "api/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBStarterConfig struct {
	DATA *config.Data_Config
}
type Db_mongo struct {
	url           string
	collection    *mongo.Collection
	regcollection *mongo.Collection
}

func (db *Db_mongo) Db_start(ztructDB DBStarterConfig) {
	db.url = ztructDB.DATA.Database.Url
	client, err := mongo.NewClient(options.Client().ApplyURI(db.url))
	if err != nil {
		fmt.Print(err)
	}
	ctx, cte := context.WithTimeout(context.Background(), 10*time.Second)
	defer cte()

	err = client.Connect(ctx)
	db.collection = client.Database("WEB").Collection(ztructDB.DATA.Database.Collection)
	db.regcollection = client.Database("WEB").Collection(ztructDB.DATA.Database.Recollection)

}
