package tools

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	client     *mongo.Client
	collection *mongo.Collection
	uri        string
	Options    options.IndexOptions
}

func New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

func (s *Session) InitMongoDB() error {
	var ClientOpts = options.Client().
		// 基本设置
		SetConnectTimeout(10 * time.Second).     // 连接超时
		SetHosts([]string{"10.100.0.31:27017"}). // 指定服务器地址
		SetMaxPoolSize(10).                      // 连接池连接数 - 最大
		SetMinPoolSize(1)                        // 连接池连接数 - 最小

	// 创建客户端
	client, err := mongo.Connect(context.TODO(), ClientOpts)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

// 创建索引
// 参考：https://stackoverflow.com/questions/56759074/how-do-i-create-a-text-index-in-mongodb-with-golang-and-the-mongo-go-driver
func (s *Session) AddIndex(dbName string, collectionName string, indexKeys interface{}) error {

	// 手动实现，可以用
	// aaa := options.Index()
	// aaa.SetUnique(true)

	coll := s.client.Database(dbName).Collection(collectionName)
	indexName, err := coll.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    indexKeys,
		Options: &s.Options,
		// Options: options.Index().SetUnique(true),	// 原始格式
	})
	if err != nil {
		return err
	}
	fmt.Println(indexName)
	return nil
}

// 插入一条数据
func (s *Session) InsertOne(dbName, collectionName string, doc interface{}) {
	coll := s.client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("主键冲突")
			return
		}
		log.Panicln(err)
	}
	log.Printf("插入文档的 ID：%v\n", result.InsertedID)
}

// 插入多条数据
func (s *Session) InsertMany(dbName, collectionName string, doc []interface{}) {
	coll := s.client.Database(dbName).Collection(collectionName)

	result, err := coll.InsertMany(context.TODO(), doc)
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			log.Panicln(err)
			// log.Println("主键冲突")
			// return
		}

	}

	totalCount := len(doc)
	count := len(result.InsertedIDs)
	log.Printf("传入文档数量：%v, 插入文档数量：%v", totalCount, count)
}

// // 查找数据
// func (s *Session) FindOne(dbName, collectionName string, doc []interface{}) {
// 	coll := s.client.Database(dbName).Collection(collectionName)
// 	result := coll.FindOne(context.TODO(), doc)
// 	log.Printf("插入文档数量：%v", result)
// }
