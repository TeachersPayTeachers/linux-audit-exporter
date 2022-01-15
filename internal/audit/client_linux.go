// +build linux

package audit

import (
	"fmt"

	libaudit "github.com/elastic/go-libaudit/v2"
)

type linuxClient struct {
	auditClient *libaudit.AuditClient
}

func New() (Client, error) {
	//	if os.Geteuid() != 0 {
	//		return nil, ErrNotRoot
	//	}

	auditClient, err := libaudit.NewMulticastAuditClient(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create multicast audit client: %w", err)
	}

	return &linuxClient{auditClient}, nil
}

func (c *linuxClient) GetStatus() (*Status, error) {
	s, err := c.auditClient.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("failed to get audit status: %w", err)
	}

	return &Status{
		Backlog:         s.Backlog,
		BacklogLimit:    s.BacklogLimit,
		BacklogWaitTime: s.BacklogWaitTime,
		Enabled:         s.Enabled,
		Failure:         s.Failure,
		Lost:            s.Lost,
		RateLimit:       s.RateLimit,
	}, nil
}
