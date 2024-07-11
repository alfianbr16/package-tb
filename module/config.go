package module

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRINGS")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "TBAlfian",
}

var MongoConn = atdb.MongoConnect(MongoInfo)