package db

import (
	"context"
	"fmt"

	//"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db Db_mongo) Db_InsertOne(Insert map[string]string) {

	_, err := db.collection.InsertOne(context.TODO(), Insert)
	if err != nil {
		fmt.Print(err)
	}
}

type DATA struct {
	Email   string `bson:"Email" json:"Email"`
	Subdata struct {
		Password, Username, Tag, UserId string
	}
}

func (db Db_mongo) Db_FixOneStuck(filter, update interface{}) {

	_, err := db.collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println(err)
	}
}
func (db Db_mongo) Db_InsertOneS(Insert interface{}) {

	_, err := db.collection.InsertOne(context.TODO(), Insert)
	if err != nil {
		fmt.Print(err)
	}
}
func (db Db_mongo) Db_FindtOne(dfkdf string, Username interface{}, c chan primitive.D) error {
	var result bson.D
	f := bson.D{{dfkdf, Username}}

	err := db.collection.FindOne(context.TODO(), f).Decode(&result)
	if err != nil {
		c <- result
		return err
	}
	c <- result
	return nil
}
func (db Db_mongo) Db_FindALL(dfkdf string, something interface{}) ([]primitive.M, error) {

	f := bson.D{{dfkdf, something}}
	coll := db.collection
	axc, err := coll.Find(context.TODO(), f)
	var results []bson.M
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.M
	if err = coll.FindOne(context.TODO(), f).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
func (db Db_mongo) Db_FindALLD(dfkdf string, asdsd string, som interface{}, something interface{}) ([]primitive.D, error) {

	f := bson.D{{dfkdf, something}, {asdsd, som}}
	coll := db.collection
	axc, err := coll.Find(context.TODO(), f)
	var results []bson.D
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.D
	if err = coll.FindOne(context.TODO(), f).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
func (db Db_mongo) Db_FindALLM(dfkdf string, something interface{}, merify string) ([]primitive.M, error) {

	f := bson.M{dfkdf: bson.M{merify: something}}
	coll := db.collection
	axc, err := coll.Find(context.TODO(), f)
	var results []bson.M
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		fmt.Print(err)
	}

	var result bson.M
	if err = coll.FindOne(context.TODO(), f).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
func (db Db_mongo) Db_Delete(something interface{}) error {
	coll := db.collection
	f, err := coll.DeleteOne(context.TODO(), something)
	fmt.Print(f)
	//.Remove(something)
	return err
}
