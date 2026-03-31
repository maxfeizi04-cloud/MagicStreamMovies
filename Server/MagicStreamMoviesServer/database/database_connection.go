package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client 是全局共享的 MongoDB 客户端。
// 包初始化时会先执行 DBInstance()，建立连接后保存在这里，后续整个项目都复用这一个客户端。
var Client *mongo.Client = DBInstance()

// DBInstance 用于创建 MongoDB 客户端实例。
//
// 对新人来说，可以把这个函数理解为“数据库连接初始化入口”：
// - 先从 .env 文件读取环境变量。
// - 再读取 MONGODB_URI 连接字符串。
// - 最后通过官方 MongoDB 驱动创建客户端。
//
// 依赖的环境变量：
// - MONGODB_URI：MongoDB 的完整连接地址。
//
// 如果本地没有 .env 文件，会打印警告日志。
// 这是因为有些部署环境会直接注入环境变量，不一定依赖 .env 文件。
func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	MongoDb := os.Getenv("MONGODB_URI")
	if MongoDb == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	fmt.Println("MONGODB_URI:", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil
	}

	return client
}

// OpenCollection 根据集合名称返回对应的 MongoDB collection 对象。
//
// 可以把它理解为“打开某张表”，虽然 MongoDB 严格来说不是关系型数据库，
// 但对初学者而言，把 collection 暂时类比成“表”会更容易理解。
//
// 依赖的环境变量：
// - DATABASE_NAME：当前项目要使用的数据库名称。
func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME:", databaseName)
	collection := Client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}
	return collection
}
