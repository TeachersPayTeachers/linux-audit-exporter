// +build !linux

package audit

type noopClient struct{}

func New() (Client, error) {
	return &noopClient{}, nil
}

func (c *noopClient) GetStatus() (*Status, error) {
	return &Status{}, nil
}
