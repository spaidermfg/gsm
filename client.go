package gsm

import (
	"bufio"
	"container/ring"
	"net"
	"sync"
)

type deliverMsgPart string

type Client struct {
	addr       string
	user       string
	password   string
	accessCode string

	ringCounter *ring.Ring

	// Establishing TCP connection
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer

	submitShortMesCh     chan []string
	deliverNotifyOpCh    chan []string
	deliverShortMesCh    chan []string
	deliverMsgPartCh     chan deliverMsgPart
	deliverMsgCompleteCh chan deliverMsgPart
	closeChan            chan struct{}
	wg                   *sync.WaitGroup
	once                 sync.Once
}

func (c *Client) Close() {
	c.once.Do(func() {
		close(c.closeChan)
	})

	if c.conn != nil {
		c.conn.Close()
	}

	// wait for all goroutines to terminate
	c.wg.Wait()
}
