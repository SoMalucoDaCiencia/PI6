package database

import (
	"PI6/src/models"

	"gorm.io/gorm"
)

type RegisterRepoImpl struct {
	IRepository[uint64, models.Register]
}

func (this *RegisterRepoImpl) Insert(obj models.Register) error {
	var db *gorm.DB
	if obj.Id == nil {
		db = db.Create(obj)
	} else {
		db1 := db.Session(&gorm.Session{})
		db = db1.Save(obj)
	}
	return db.Error
}

func (this *RegisterRepoImpl) Select(objs *[]models.Register) error {
	db, err := GetConn()
	if err != nil {
		return err
	}

	var target int64
	return db.Find(&target).Error
}

func (this *RegisterRepoImpl) SelectById(id uint64, obj *models.Register) error {
	db, err := GetConn()
	if err != nil {
		return err
	}

	db = db.Model(obj).Select(obj)
	return db.First(obj, id).Error
}

func (this *RegisterRepoImpl) SelectByQuery(query string, objs *[]RegisterRepoImpl) error {
	db, err := GetConn()
	if err != nil {
		return err
	}

	db = db.Raw(query).Scan(&objs)
	if db.Error == nil || db.Error.Error() == "sql: no rows in result set" {
		return nil
	}
	return db.Error
}

func (this *RegisterRepoImpl) Count() (int64, error) {
	db, err := GetConn()
	if err != nil {
		return -1, err
	}

	var target int64
	return target, db.Count(&target).Error
}

func (this *RegisterRepoImpl) DeleteById(obj models.Register) error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(obj).Delete(obj, obj.Id).Error
}

func (this *RegisterRepoImpl) UpdadeById(obj int64, properties map[string]any) error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(models.Register{}).Where("id = ?", obj).Updates(properties).Error
}
