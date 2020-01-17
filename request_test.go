package request

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T)  {
	TestingT(t)
}

type MySuite struct {

}
var _ = Suite(&MySuite{})

func (s *MySuite)TestRequest(c *C) {
	_, res, reqErr := Request(Options{Url:"https://bing.com"})
	c.Assert(reqErr, Equals, nil)
	c.Assert(res.StatusCode, Equals, 200)
}