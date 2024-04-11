package repositories

type IGenericRepository[T any] interface {
	GetAll() ([]*T, error)
	GetById(id string) (*T, error)
	Create(data *T) (*T, error)
	Update(data *T) (*T, error)
	Delete(data *T) (*T, error)
}