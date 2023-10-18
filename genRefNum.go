package gsm

import (
	"container/ring"
	"fmt"
)

// generate valid transaction reference numbers
// ringing from 00 ~ 99

func initRefNum(c *Client) {
	ringCounter := ring.New(MaxRefNum)
	for j := 0; j < MaxRefNum; j++ {
		ringCounter.Value = []byte(fmt.Sprintf("%02d", j))
		ringCounter = ringCounter.Next()
	}

	c.ringCounter = ringCounter
}

func nextRefNum(c *Client) []byte {
	refNum := (c.ringCounter.Value).([]byte)
	c.ringCounter = c.ringCounter.Next()
	return refNum
}
