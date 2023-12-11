package _interface

type CRUDRepositories interface {
	GetAll() ([]string, error)
	GetByID(id uint) (interface{}, error)
	Create(data interface{}) error
	Update(id uint, data interface{}) error
	Delete(id uint) error
}
