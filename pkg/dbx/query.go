package dbx

import "gorm.io/gorm"

func Exists(tx *gorm.DB, model interface{}, conds ...interface{}) (exists bool, err error) {
	err = tx.Model(model).Select("count(*) > 0").First(&exists, conds...).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
