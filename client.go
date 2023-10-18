package gsm

import "container/ring"

type Client struct {
	addr       string
	user       string
	password   string
	accessCode string

	ringCounter *ring.Ring
}
