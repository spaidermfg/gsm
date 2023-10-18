package gsm

import (
	"bufio"
	"net"
)

func (c *Client) Connect() error {
	c.initRefNum()

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn

	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)

	if _, err = c.writer.Write(createLogReq(c.nextRefNum(), c.user, c.password)); err != nil {
		return err
	}

	if err = c.writer.Flush(); err != nil {
		return err
	}

	resp, err := c.reader.ReadString(ETX)
	if err != nil {
		return err
	}

	if err = parseSessionResp(resp); err != nil {
		return err
	}
	return nil
}

func createLogReq(num []byte, user, pwd string) []byte {
	return []byte{}
}

// if credentials are invalid, return error
// otherwise it returns nil
func parseSessionResp(resp string) error {
	return nil
}
