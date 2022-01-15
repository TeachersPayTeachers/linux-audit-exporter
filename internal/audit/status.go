package audit

type Status struct {
	Backlog               uint32
	BacklogLimit          uint32
	BacklogWaitTime       uint32
	BacklogWaitTimeActual uint32
	Enabled               uint32
	Failure               uint32
	Lost                  uint32
	RateLimit             uint32
}
