package audit

type Client interface {
	GetStatus() (*Status, error)
}
