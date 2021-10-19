package logic

import (
	"go-project/models"
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
