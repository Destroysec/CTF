package db

import (
	"context"
	"fmt"

	//"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//async find docs
func (db Db_mongo) Db_FindtOne_UniDentify(dfkdf string, Username interface{}, c chan primitive.D) error {
	var result bson.D
	f := bson.D{{dfkdf, Username}}

	err := db.regcollection.FindOne(context.TODO(), f).Decode(&result)
	if err != nil {
		c <- result
		return err
	}
	c <- result
	return nil
}

//async updata docs
func (db Db_mongo) Db_FixOneStuck_UniDentify(filter, update interface{}) {

	_, err := db.regcollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println(err)
	}
}

//
func (db Db_mongo) Db_Delete_UniDentify(something interface{}) error {

	_, err := db.regcollection.DeleteOne(context.TODO(), something)

	//.Remove(something)
	return err
}
func (db Db_mongo) Db_DeleteMany_UniDentify(something interface{}) error {

	f, err := db.regcollection.DeleteMany(context.TODO(), something)
	fmt.Print(f)
	//.Remove(something)
	return err
}
func (db Db_mongo) Db_FindALLD_UniDentify(dfkdf bson.D) ([]primitive.D, error) {

	axc, err := db.regcollection.Find(context.TODO(), dfkdf)
	var results []bson.D
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.D
	if err = db.regcollection.FindOne(context.TODO(), dfkdf).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
func (db Db_mongo) Db_FindALLunD(dfkdf string, asdsd string, som interface{}, something interface{}) ([]primitive.D, error) {

	f := bson.D{{dfkdf, something}, {asdsd, som}}

	axc, err := db.regcollection.Find(context.TODO(), f)
	var results []bson.D
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.D
	if err = db.regcollection.FindOne(context.TODO(), f).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
