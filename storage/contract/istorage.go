package contract

type ReturnData struct {
	Data []any
}

type IStorage interface {
	Exist(data string) bool
	Store(data interface{}) bool
	GetAllRecordes() any
}
