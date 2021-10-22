package tars

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
	"testing"
)

type hook struct {
}

func (m hook) Before(level rogger.LogLevel) map[string]string {
	return map[string]string{
		"a": "1",
		"B": "bbbb",
	}
}

func (m hook) After(level rogger.LogLevel, message string) {
	fmt.Println(level, message)
}

func TestTLOG(t *testing.T) {
	rogger.SetLevel(rogger.INFO)
	TLOG.SetHook(hook{})
	TLOG.Info("abc")
	TLOG.Errorf("123")
}
