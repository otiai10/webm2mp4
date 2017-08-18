package goavcodec

import (
	"io/ioutil"
	"os/exec"
)

const (
	avconv = "avconv"
)

// Client ...
type Client struct {
	bin    string
	StdOut []byte
	StdErr []byte
}

// NewClient looks path for `ffmpeg` and returns initialized Client.
func NewClient() (*Client, error) {
	bin, err := exec.LookPath(avconv)
	if err != nil {
		return nil, err
	}
	return &Client{
		bin: bin,
	}, nil
}

// Convert just converts src to dest with using `ffmpeg -i`
func (c *Client) Convert(src, dest string) error {
	cmd := exec.Command(c.bin, "-y", "-i", src, dest)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if c.StdErr, err = ioutil.ReadAll(stderr); err != nil {
		return err
	}
	if c.StdOut, err = ioutil.ReadAll(stdout); err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
