package controller

import (
	"testing"

	"github.com/dangviethung096/core"
)

func TestMain(m *testing.M) {
	core.Init("../config.json")
	m.Run()
}
