package dao

import (
	"admin/internal/model"
	"admin/pkg/conf"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Log struct {
	Qid     int64
	Name    string
	Message string
	Time    time.Time
	New_que model.Questions
	Old_que model.Questions
}

type Mongo struct {
	collection *mongo.Collection
	dao        *Dao
}

func NewMongo(dao *Dao, yaml *conf.Yaml) (*Mongo, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/basic?authSource=admin&authMechanism=SCRAM-SHA-1",
		yaml.Mongo.Username, yaml.Mongo.Password, yaml.Mongo.Host, yaml.Mongo.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	collection := client.Database("basic").Collection("log")
	return &Mongo{
		collection: collection,
		dao:        dao,
	}, nil
}

func (m *Mongo) InsertLog(questions model.Questions, name string, operate string) error {
	switch operate {
	case "add":
		if _, err := m.collection.InsertOne(context.TODO(), Log{
			Qid:     questions.Id,
			Name:    name,
			Message: "新增",
			Time:    time.Now(),
			New_que: questions,
		}); err != nil {
			return err
		}
		return nil
	case "delete":
		if _, err := m.collection.InsertOne(context.TODO(), Log{
			Qid:     questions.Id,
			Name:    name,
			Message: "删除",
			Time:    time.Now(),
			New_que: questions,
		}); err != nil {
			return err
		}
		return nil
	case "update":
		q := model.Questions{
			Id:       questions.Id,
			Question: questions.Question,
			Answer:   questions.Answer,
			Tag:      questions.Tag,
		}
		m.dao.db.First(&q)
		if _, err := m.collection.InsertOne(context.TODO(), Log{
			Qid:     questions.Id,
			Name:    name,
			Message: "修改",
			Time:    time.Now(),
			New_que: questions,
			Old_que: q,
		}); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("error")
	}
}

func (m *Mongo) FindLogs() ([]*Log, error) {
	var result []*Log
	cur, err := m.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem Log
		if err := cur.Decode(&elem); err != nil {
			return nil, err
		}
		result = append(result, &elem)
	}
	_ = cur.Close(context.TODO())
	return result, nil
}
