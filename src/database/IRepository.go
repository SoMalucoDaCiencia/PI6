package database

type IRepository[K, T any] interface {
	Insert(obj T) error
	Select(obj *[]T) error
	SelectById(id K, obj T) error
	Count(obj T) (int64, error)
	DeleteById(obj T) error
}
