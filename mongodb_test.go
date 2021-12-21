package tools_test

import (
	"fmt"
	"log"

	"testing"

	"github.com/xiaoliuxiao6/tools"
	"go.mongodb.org/mongo-driver/bson"
)

type Blog struct {
	Name string
	Age  int
}

// 基本使用示例
func TestUsaged(t *testing.T) {

	dbName := "insertDB"
	collectionName := "haikus"

	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 插入结构体 - 单条数据
	bloger := Blog{
		Name: "Tom",
		Age:  18,
	}
	mongoClient.InsertOne(dbName, collectionName, bloger)

	// 插入结构体 - 多条数据
	bloger1 := Blog{
		Name: "Tom1",
		Age:  18,
	}
	bloger2 := Blog{
		Name: "Tom2",
		Age:  18,
	}
	blogers := make([]interface{}, 0)
	blogers = append(blogers, bloger1)
	blogers = append(blogers, bloger2)
	fmt.Println(blogers)
	mongoClient.InsertMany(dbName, collectionName, blogers)

	// 插入单个文档
	doc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	mongoClient.InsertOne(dbName, collectionName, doc)

	// 插入多个文档
	docs := []interface{}{
		bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}},
		bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}},
	}
	mongoClient.InsertMany(dbName, collectionName, docs)
}

// 插入单条数据
func TestInsertOne(t *testing.T) {

	dbName := "insertDB"
	collectionName := "haikus"

	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 插入结构体 - 单条数据
	bloger := Blog{
		Name: "Tom",
		Age:  18,
	}
	mongoClient.InsertOne(dbName, collectionName, bloger)
	result := mongoClient.InsertOne(dbName, collectionName, bloger)
	// log.Printf("插入文档的 ID：%v\n", result.InsertedID)
	// fmt.Println(len(result.InsertedID))
	if result == nil {
		fmt.Println("插入数量为空")
	}
}

// 插入多条数据
func TestInsertMany(t *testing.T) {

	dbName := "insertDB"
	collectionName := "haikus"

	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 插入结构体 - 多条数据
	bloger1 := Blog{
		Name: "Tom1",
		Age:  18,
	}
	bloger2 := Blog{
		Name: "Tom2",
		Age:  18,
	}
	blogers := make([]interface{}, 0)
	blogers = append(blogers, bloger1)
	blogers = append(blogers, bloger2)
	// fmt.Println(blogers)
	mongoClient.InsertMany(dbName, collectionName, blogers)
}

// 创建索引
func TestAddIndex(t *testing.T) {
	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 设置是否为唯一索引
	mongoClient.Options.SetUnique(true)

	// // 单字段索引（方式1）
	// aaa := map[string]interface{}{
	// 	"myfieldname_type1": 1,
	// }
	// mongoClient.AddIndex("mydb", "mycollection111", aaa) // to descending set it to -1

	// // 单字段索引（方式2）
	// mongoClient.AddIndex("mydb", "mycollection111", bson.M{"myfieldname_type2": 1}) // to descending set it to -1

	// // 符合索引
	// mongoClient.AddIndex("mydb", "mycollection222", bson.D{{"myFirstField", 1}, {"mySecondField", -1}}) // to descending set it to -1

	// // 文本索引
	// mongoClient.AddIndex("mydb", "mycollection333", bson.D{{"myFirstTextField", "text"}, {"mySecondTextField", "text"}})

	// 插入多个文档
	// 符合索引
	mongoClient.AddIndex("mydb", "mycollection333", bson.D{{"myFirstField", 1}, {"mySecondField", -1}}) // to descending set it to -1

	// mongoClient.set
	docs := []interface{}{
		// 	bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}, {"_id", "111"}},
		// 	bson.D{{"myFirstField", "bbb"}, {"mySecondField", "bbb"}, {"_id", "111"}},

		bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}},
		bson.D{{"myFirstField", "aaa"}, {"mySecondField", "aaa"}},
	}
	mongoClient.InsertMany("mydb", "mycollection333", docs)
}

// 查找单条数据
func TestFindOne(t *testing.T) {
	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 解析返回值用
	type AutoGenerated struct {
		Timestamp        string        `json:"timestamp"`
		BaseFeePerGas    string        `json:"baseFeePerGas"`
		LogsBloom        string        `json:"logsBloom"`
		MixHash          string        `json:"mixHash"`
		Size             string        `json:"size"`
		TotalDifficulty  string        `json:"totalDifficulty"`
		Hash             string        `json:"hash"`
		Miner            string        `json:"miner"`
		Nonce            string        `json:"nonce"`
		Sha3Uncles       string        `json:"sha3Uncles"`
		GasLimit         string        `json:"gasLimit"`
		ParentHash       string        `json:"parentHash"`
		ReceiptsRoot     string        `json:"receiptsRoot"`
		StateRoot        string        `json:"stateRoot"`
		TransactionsRoot string        `json:"transactionsRoot"`
		Uncles           []interface{} `json:"uncles"`
		Difficulty       string        `json:"difficulty"`
		ExtraData        string        `json:"extraData"`
		GasUsed          string        `json:"gasUsed"`
		Number           string        `json:"number"`
	}
	var ret AutoGenerated

	// 查询方式1（手动写 Map 类型过滤器）
	filter := map[string]interface{}{
		// "number": "0xd27b8c",
	}

	// 排序（可选）
	sort := map[string]interface{}{
		"number": -1,
	}
	opts := mongoClient.FindOneOptions.SetSort(sort)

	// 执行查询
	notFind, err := mongoClient.FindOne("eth", "BlockByNumberResultInfo", filter, &ret, opts)
	// notFind, err := mongoClient.FindOne("eth", "BlockByNumberResultInfo", filter, &ret)
	if err != nil {
		fmt.Println("FindOne 运行结果错误")
	}
	if notFind == true {
		fmt.Println("没有找到结果")
	}
	// fmt.Println(ret.Number)
	tools.StructPrint(ret)
}

// 查找多条数据
func TestFind(t *testing.T) {
	// 建立连接
	mongoClient := tools.New("mongodb://127.0.0.1")
	err := mongoClient.InitMongoDB()
	if err != nil {
		log.Panicln(err)
	}

	// 查询方式1（手动写 Map 类型过滤器）
	filter := map[string]interface{}{
		"blockNumber": "0xcf8500",
	}

	// 排序（可选）
	sort := map[string]interface{}{
		"blockNumber": -1,
	}
	opts := mongoClient.FindOptions.SetSort(sort)

	// 执行查询
	results, err := mongoClient.Find("eth", "BlockByNumberResultTransactions", filter, opts)
	if err != nil {
		fmt.Println("Find 运行结果错误")
	}

	for _, result := range results {
		fmt.Println(result["hash"])
	}
}

// // 测试事务（只有 4.0 以上版本且副本集群可用）
// func TestTransaction(t *testing.T) {

// 	// 建立连接
// 	dbName := "insertDB"
// 	collectionName := "haikus"
// 	mongoClient := tools.New("mongodb://127.0.0.1")
// 	err := mongoClient.InitMongoDB()
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	type Stu struct {
// 		Name string `bson:"_id"`
// 		Age  int
// 	}

// 	// 插入结构体 - 单条数据
// 	bloger := Stu{
// 		Name: "Tom",
// 		Age:  20,
// 	}
// 	result := mongoClient.InsertOne(dbName, collectionName, bloger)
// 	if result == nil {
// 		fmt.Println("插入数量为空")
// 	} else {
// 		fmt.Println(result.InsertedID)
// 	}
// }
