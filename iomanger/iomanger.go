package iomanger

type IoManger interface {
	ReadFile() (prices []string, err error)
	WriteResult(data any) error
}
