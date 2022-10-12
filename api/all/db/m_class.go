package db

import (
	"context"
	"fmt"

	"time"
	//"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	config "api/all/setting"

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
	if err != nil {
		fmt.Print(err)
	}
	db.collection = client.Database("WEB").Collection(ztructDB.DATA.Database.Collection)
	db.regcollection = client.Database("WEB").Collection(ztructDB.DATA.Database.Recollection)

}

type FormUD struct {
	db     *Db_mongo
	Insert struct{}
}
type FormUN struct {
	db     *Db_mongo
	Insert struct{}
}
type INSERTONE interface {
	Db_insertOne()
}

func (F *FormUD) Db_insertOne() {
	_, err := F.db.collection.InsertOne(context.TODO(), F.Insert)
	if err != nil {
		fmt.Print(err)
	}
}
func (F *FormUN) Db_insertOne() {
	_, err := F.db.regcollection.InsertOne(context.TODO(), F.Insert)
	if err != nil {
		fmt.Print(err)
	}
}
