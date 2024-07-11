package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hewan struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis 	        string             	`bson:"jenis,omitempty" json:"jenis,omitempty"`
	Umur 	  		string 			 	`bson:"umur,omitempty" json:"umur,omitempty"`
	Ras 			string 			 	`bson:"ras,omitempty" json:"ras,omitempty"`
}

type MakananHewan struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Hewan           Hewan             	`bson:"hewan,omitempty" json:"hewan,omitempty"`
	JenisMakanan    string             	`bson:"jenismakanan,omitempty" json:"jenismakanan,omitempty"`
	Bahan         	string             	`bson:"bahan,omitempty" json:"bahan,omitempty"`
	Berat           string             	`bson:"berat,omitempty" json:"berat,omitempty"`
	Rasa         	string             	`bson:"rasa,omitempty" json:"rasa,omitempty"`
	Merk		    string             	`bson:"merk,omitempty" json:"merk,omitempty"`
	Harga		    string             	`bson:"harga,omitempty" json:"harga,omitempty"`
	Tanggal  		primitive.DateTime 	`bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}