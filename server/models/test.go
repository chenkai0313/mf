package models

type Test struct {
	Id int32 `json:"id"`
}

func (Test) TableName() string {
	return "test"
}
