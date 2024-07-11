package packagetbku

import (
	"fmt"
	"testing"

	"github.com/alfianbr16/package-tb/model"
	"github.com/alfianbr16/package-tb/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertMakanan(t *testing.T) {
	hewan := model.Hewan{
		Jenis: "Kucing",
		Umur:  "7",
		Ras:   "Persia",
	}
	jenismakanan := "basah"
	bahan := "Daging Sapi"
	berat := "1"
	rasa := "Abon"
	merk := "sehatpet"
	harga := "40000"
	insertedID, err := module.InsertMakanan(module.MongoConn, "makananhewan", hewan, jenismakanan, bahan, berat, rasa, merk, harga)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetMakananFromID(t *testing.T) {
	id := "668f9c9f08369bd6d44afa85"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetMakananFromID(objectID, module.MongoConn, "makananhewan")
	if err != nil {
		t.Fatalf("error calling GetMakananFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAllMakanan(t *testing.T) {
	data := module.GetAllMakanan(module.MongoConn, "makananhewan")
	fmt.Println(data)
}

func TestDeleteMakananByID(t *testing.T) {
	id := "668f9b42b6a568ecb22f916c" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMakananByID(objectID, module.MongoConn, "makananhewan")
	if err != nil {
		t.Fatalf("error calling DeleteMakananByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetMakananFromID
	_, err = module.GetMakananFromID(objectID, module.MongoConn, "makananhewan")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
