package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/alfianbr16/package-tb/model"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMakanan(db *mongo.Database, col string, hewan model.Hewan, jenismakanan string, bahan string, berat string, rasa string,merk string, harga string) (insertedID primitive.ObjectID, err error) {
	makanan := bson.M{
		"hewan":    hewan,
		"jenismakanan":     jenismakanan,
		"bahan":    bahan,
		"berat":    berat,
		"rasa":    rasa,
		"merk":    merk,
		"harga":    harga,
		"tanggal": primitive.NewDateTimeFromTime(time.Now().UTC()),
	}
	result, err := db.Collection(col).InsertOne(context.Background(), makanan)
	if err != nil {
		fmt.Printf("InsertMakanan: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllMakanan(db *mongo.Database, col string) (data []model.MakananHewan) {
	makanan := db.Collection(col)
	filter := bson.M{}
	cursor, err := makanan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetMakananFromID(_id primitive.ObjectID, db *mongo.Database, col string) (makananhewan model.MakananHewan, errs error) {
	makanan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := makanan.FindOne(context.TODO(), filter).Decode(&makananhewan)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return makananhewan, fmt.Errorf("no data found for ID %s", _id)
		}
		return makananhewan, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return makananhewan, nil
}

func UpdateMakanan(db *mongo.Database, col string, id primitive.ObjectID, hewan model.Hewan, jenismakanan string, bahan string, berat string, rasa string,merk string, harga string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
		"hewan":    hewan,
		"jenismakanan":     jenismakanan,
		"bahan":    bahan,
		"berat":    berat,
		"rasa":    rasa,
		"merk":    merk,
		"harga":    harga,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMakanan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteMakananByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	makanan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := makanan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}