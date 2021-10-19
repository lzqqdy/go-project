package models

import (
	"github.com/jinzhu/gorm"
)

type Test struct {
	Model

	Name  string `json:"name"`
	State int    `json:"state"`
}

// GetTest gets a list of test based on paging and constraints
func GetTest(pageNum int, pageSize int, maps interface{}) ([]Test, error) {
	var (
		test []Test
		err  error
	)
	err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&test).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return test, nil
}

// GetTestTotal counts the total number of test based on the constraint
func GetTestTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Test{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
