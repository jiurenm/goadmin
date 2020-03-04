package dao

import (
	"admin/internal/model"
	"admin/internal/server/thrift"
	"admin/pkg/conf"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Dao struct {
	db     *gorm.DB
	node   *snowflake.Node
	thrift *thrift.Thrift
}

func NewDao(thrift *thrift.Thrift, db *gorm.DB) *Dao {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil
	}
	return &Dao{db: db, node: node, thrift: thrift}
}

func New(yaml *conf.Yaml) (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		yaml.Mysql.Username, yaml.Mysql.Password, yaml.Mysql.Host, yaml.Mysql.Port, yaml.Mysql.Dbname)
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(yaml.Mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(yaml.Mysql.MaxOpenConns)
	db.LogMode(false)
	return db, nil
}

func (d *Dao) Close() error {
	return d.db.Close()
}

func (d *Dao) GetAll() (*[]model.Questions, error) {
	var q []model.Questions
	if err := d.db.Find(&q).Error; err != nil {
		return &q, err
	}
	return &q, nil
}

func (d *Dao) GetAllQuestion() (*[]model.Question, error) {
	var q []model.Question
	if err := d.db.Find(&q).Error; err != nil {
		return &q, err
	}
	return &q, nil
}

func (d *Dao) Add(questions model.Questions) error {
	questions.Id = d.node.Generate().Int64()
	questions.Tag = d.thrift.GetKeyWord(questions.Question)
	if !d.db.NewRecord(questions) {
		if err := d.db.Create(&questions).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (d *Dao) AddQuestion(question model.Question) error {
	question.Id = d.node.Generate().Int64()
	if !d.db.NewRecord(question) {
		if err := d.db.Create(&question).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (d *Dao) Delete(questions model.Questions) error {
	if err := d.db.Delete(&questions).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) DeleteQuestion(question model.Question) error {
	if err := d.db.Delete(&question).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) Update(questions model.Questions) error {
	questions.Tag = d.thrift.GetKeyWord(questions.Tag)
	if err := d.db.Save(&questions).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UpdateQuestion(question model.Question) error {
	if err := d.db.Save(&question).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) AddBatch(questions []model.Questions) error {
	sql := "INSERT INTO `questions`(`id`, `question`, `answer`, `tag`) VALUES "
	for index, value := range questions {
		if len(questions)-1 == index {
			sql += fmt.Sprintf("(%d,'%s','%s','%s');",
				d.node.Generate().Int64(), value.Question, value.Answer, d.thrift.GetKeyWord(value.Question))
		} else {
			sql += fmt.Sprintf("(%d,'%s','%s','%s'),",
				d.node.Generate().Int64(), value.Question, value.Answer, d.thrift.GetKeyWord(value.Question))
		}
	}
	return d.db.Exec(sql).Error
}
