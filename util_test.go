package chinaums

import (
	"testing"
	"time"
)

func TestAuthorization(t *testing.T) {
	authorization := Authorization("abc", time.Now())
	t.Logf("auth:%v", authorization)
}
