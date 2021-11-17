package logic

import (
	"context"
	"fmt"
	"go-project/models"
	"go-project/pkg/es"
	"go-project/pkg/logger"
	"go-project/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Test struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Test) Count() (int, error) {
	return models.GetTestTotal(t.getMaps())
}
func (t *Test) GetAll() ([]models.Test, error) {
	var (
		test []models.Test
	)
	test, err := models.GetTest(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	return test, nil
}

func (t *Test) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
func TestEs() {
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	type Article struct {
		Title   string    // 文章标题
		Content string    // 文章内容
		Author  string    // 作者
		Created time.Time // 发布时间
	}
	// 定义一篇博客
	blog := Article{Title: "golang es教程", Content: "go如何操作ES", Author: "test", Created: time.Now()}

	// 使用client创建一个新的文档
	put1, err := es.Client.Index().
		Index("blogs"). // 设置索引名称
		Id("1").        // 设置文档id
		BodyJson(blog). // 指定前面声明struct对象
		Do(ctx)         // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		logger.Logger("app.index.es").Error(err)
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
}
func TestMongoDB() {
	// 2, 选择数据库
	database := mongodb.Client.Database("test")
	// 3, 选择表my_collection
	collection := database.Collection("a1")

	type Test struct {
		Name string `bson:"name"` // 定义字段
		Age  int64  `bson:"age"`
	}
	// 4, 插入记录(bson)
	record := &Test{
		Name: "lzqqdy",
		Age:  18,
	}
	var (
		err    error
		result *mongo.InsertOneResult
	)
	if result, err = collection.InsertOne(context.Background(), record); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	docId := result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
}
