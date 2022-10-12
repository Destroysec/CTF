package db

import "context"

func (db *Db_mongo) RemoveArray(filter, change interface{}) error {

	_, err := db.collection.UpdateOne(context.TODO(), filter, change)
	return err
}
