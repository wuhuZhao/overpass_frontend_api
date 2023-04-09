package dao

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Idl struct {
	ID         uint   `gorm:"primaryKey"`
	Source     string `form:"source,required"`
	Name       string `form:"name,required"`
	Version    int
	Use        bool
	Author     string `form:"author,required"`
	CreateTime time.Time
	Tag        string
}

// TableName set tableName in mysql
func (i Idl) TableName() string {
	return "idl_config"
}

// IdlDao dao层查询
type IdlDao struct {
	db *gorm.DB
}

// NewIdlDao ...
func NewIdlDao(db *gorm.DB) *IdlDao {
	return &IdlDao{
		db: db,
	}
}

// Migrate 迁移数据表
func (i *IdlDao) Migrate() {
	i.db.AutoMigrate(&Idl{})
}

// FindById ...
func (i *IdlDao) FindById(idl *Idl, id int) error {
	tx := i.db.First(idl, id)
	return tx.Error
}

// FindAll ...
func (i *IdlDao) FindAll(idls []*Idl, pageSize, pageNum int) error {
	tx := i.db.Limit(pageSize).Offset(pageNum*pageSize).Select("id", "name", "version", "use", "author", "create_time", "tag").Find(&idls)
	return tx.Error
}

func (i *IdlDao) FindByName(idl *Idl, name string) error {
	tx := i.db.Where("name = ?", name).Order("version desc").Limit(1).First(idl)
	return tx.Error
}

// UpdateOneById ...
func (i *IdlDao) UpdateOneById(idl *Idl) error {
	tx := i.db.Save(idl)
	if tx.RowsAffected != 1 {
		return fmt.Errorf("rows afffect not equal to 1")
	}
	return tx.Error
}

// InsertOne ...
func (i *IdlDao) InsertOne(idl *Idl) error {
	tx := i.db.Create(idl)
	return tx.Error
}

// DeleteOne ...
func (i *IdlDao) DeleteOne(id uint) error {
	tx := i.db.Delete(&Idl{}, id)
	return tx.Error
}
