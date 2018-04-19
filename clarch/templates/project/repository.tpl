package project

import "github.com/jinzhu/gorm"

type RepositoryProject struct{}

func (this *RepositoryProject) Transaction(db *gorm.DB, f func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := f(tx); err != nil {
		tx.Rollback()
		if tx.Error != nil {
			return tx.Error
		}
		return err
	}
	tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
