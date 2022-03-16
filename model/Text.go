package model

import (
	"gorm.io/gorm/clause"
	"time"
)

type Text struct {
	UUID      string `gorm:"column:uuid;type:char(32);not null;unique;primaryKey"`
	Path      string `gorm:"column:path;type:varchar(128);not null;unique"`
	Content   string `gorm:"column:content;type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Text) TableName() string {
	return "ogm_config_Text"
}

type TextDAO struct {
	conn *Conn
}

func NewTextDAO(_conn *Conn) *TextDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &TextDAO{
		conn: conn,
	}
}

func (this *TextDAO) Count() (int64, error) {
	var count int64
	err := this.conn.DB.Model(&Text{}).Count(&count).Error
	return count, err
}

func (this *TextDAO) Insert(_entity *Text) error {
	return this.conn.DB.Create(_entity).Error
}

func (this *TextDAO) Update(_entity *Text) error {
	// 只更新非零值
	return this.conn.DB.Updates(_entity).Error
}

func (this *TextDAO) Upsert(_entity *Text) error {
	// 在冲突时，更新除主键以外的所有列到新值。
	return this.conn.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(_entity).Error
}

func (this *TextDAO) Get(_uuid string) (*Text, error) {
	var entity Text
	err := this.conn.DB.Where("uuid = ?", _uuid).First(&entity).Error
	return &entity, err
}

func (this *TextDAO) FindByPath(_path string) (*Text, error) {
	var entity Text
	err := this.conn.DB.Where("path = ?", _path).First(&entity).Error
	return &entity, err
}

func (this *TextDAO) List(_offset int64, _count int64) (int64, []*Text, error) {
	var entityAry []*Text
	count := int64(0)
	db := this.conn.DB.Model(&Text{})
	// db = db.Where("key = ?", value)
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *TextDAO) Delete(_uuid string) error {
	return this.conn.DB.Where("uuid = ?", _uuid).Delete(&Text{}).Error
}

func (this *TextDAO) Search(_offset int64, _count int64, _path string) (int64, []*Text, error) {
	var entityAry []*Text
	count := int64(0)
	db := this.conn.DB.Model(&Text{})
	db = db.Where("path LIKE ?", "%"+_path+"%")
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}
