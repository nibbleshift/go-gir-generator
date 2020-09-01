package gst

import (
	C "gopkg.in/check.v1"
	"os"
	"strings"
	"testing"
)

type gstreamer struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&gstreamer{})
}
