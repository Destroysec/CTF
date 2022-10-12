package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type FormUD struct {
	Db                     *Db_mongo
	Insert, Update, Filter interface{}
}
type FormUN struct {
	Db                     *Db_mongo
	Insert, Update, Filter interface{}
}
type INSERTONE interface {
	Db_insertOne()
	Db_FindtOne(c chan bson.D)
	Db_FixOneStuck()
	Db_FindALL(c chan []bson.M)
}

func (F *FormUD) Db_insertOne() {
	_, err := F.Db.collection.InsertOne(context.TODO(), F.Insert)
	if err != nil {
		fmt.Print(err)
	}
}
func (F *FormUN) Db_insertOne() {
	_, err := F.Db.regcollection.InsertOne(context.TODO(), F.Insert)
	if err != nil {
		fmt.Print(err)
	}
}
func Db_insertOne(i INSERTONE) {
	i.Db_insertOne()
}

func (F *FormUD) Db_FindtOne(c chan bson.D) {
	var result bson.D
	err := F.Db.collection.FindOne(context.TODO(), F.Insert).Decode(&result)
	if err != nil {
		c <- result

	}
	c <- result

}
func (F *FormUN) Db_FindtOne(c chan bson.D) {
	var result bson.D
	err := F.Db.regcollection.FindOne(context.TODO(), F.Insert).Decode(&result)
	if err != nil {
		c <- result
	}
	c <- result
}
func Db_FindtOne(i INSERTONE, c chan bson.D) {
	i.Db_FindtOne(c)
}
func (F *FormUD) Db_FixOneStuck() {

	_, err := F.Db.collection.UpdateOne(
		context.Background(),
		F.Filter,
		F.Update,
	)
	if err != nil {
		fmt.Println(err)
	}
}
func (F *FormUN) Db_FixOneStuck() {

	_, err := F.Db.collection.UpdateOne(
		context.Background(),
		F.Filter,
		F.Update,
	)
	if err != nil {
		fmt.Println(err)
	}
}
func Db_FixOneStuck(i INSERTONE) {
	i.Db_FixOneStuck()
}

func (F *FormUD) Db_FindALL(c chan []bson.M) {

	axc, err := F.Db.collection.Find(context.TODO(), F.Filter)
	var results []bson.M
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.M
	if err = F.Db.collection.FindOne(context.TODO(), F.Filter).Decode(&result); err != nil {
		c <- nil
	}
	c <- results
}
func (F *FormUN) Db_FindALL(c chan []bson.M) {

	axc, err := F.Db.regcollection.Find(context.TODO(), F.Filter)
	var results []bson.M
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.M
	if err = F.Db.collection.FindOne(context.TODO(), F.Filter).Decode(&result); err != nil {
		c <- nil
	}
	c <- results

}

func Db_FindALL(i INSERTONE, c chan []bson.M) {
	i.Db_FindALL(c)

}
