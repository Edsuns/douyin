package dbx

import (
	"errors"
	"gorm.io/gorm"
)

func SpinOptimisticLock[T any](tx *gorm.DB, id any, modify func(*T)) error {

	var err error
	var entity T
	// spin
	for i := 0; i < 3; i++ {
		err = tx.First(&entity, id).Error
		if err != nil {
			return err
		}
		modify(&entity)
		err = tx.Save(&entity).Error
		if err == nil {
			// modification affected, finish spin
			break
		}
		// error other than ErrRecordNotFound happens
		// return the error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// if modification failed, err != nil
	return err
}
