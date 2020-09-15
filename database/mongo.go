package database
//
//import (
//	"context"
//	"go.mongodb.org/mongo-driver/mongo/readpref"
//	"log"
//	"time"
//
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//// MgoCli mongodb client
//var MgoCli *mongo.Client
//
//const (
//	DBName = "embedci"
//	Collection = "cfgjson"
//	URI = "mongodb://root@172.16.248.111:27017,172.16.248.111:27018,172.16.248.111:27019/?authSource=admin"
//	Username = "root"
//	Password = "root"
//)
//
//func InitMongo(){
//	// 连接
//	clientoption := options.Client().ApplyURI(URI).SetAuth(options.Credential{Username: Username, Password: Password}).SetMaxPoolSize(50)
//	MgoCli, err := mongo.NewClient(clientoption)
//	if err != nil {
//		log.Fatal(err)
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	err = MgoCli.Connect(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	err = MgoCli.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func DissConnect() {
//	if err := MgoCli.Disconnect(nil); err != nil {
//		panic(err)
//	}
//}
//
//func GetConnect(db, collection string) *mongo.Collection {
//	return MgoCli.Database(db).Collection(collection)
//}


import (
"context"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"log"
"time"
)
const (
	DBName = "embedci"
	Collection = "cfgjson"
	URI = "mongodb://root@172.16.248.111:27017,172.16.248.111:27018,172.16.248.111:27019/?authSource=admin"
	Username = "root"
	Password = "root"
)
type Database struct {
	Mongo *mongo.Client
}

var DB *Database

//初始化
func InitMongo() {
	DB = &Database{
		Mongo: SetConnect(URI),
	}
}

//连接设置
func SetConnect(URI string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 连接池
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI).
		SetAuth(options.Credential{Username: Username, Password: Password}).
		SetMaxPoolSize(20))
	if err != nil {
		log.Println(err)
	}
	return client
}

type Mgo struct {
	database   string
	collection string
}
