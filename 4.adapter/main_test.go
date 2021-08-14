package adapter_test

import (
	"testing"

	adapter "github.com/gsaldanab/design-pattern/4.adapter"
)

func TestMover(t *testing.T) {
	tipos := [2]string{"automovil", "bicicleta"}

	for _, tipo := range tipos {
		t.Log("tipo:", tipo)
		transportAdapter := adapter.Factory(tipo)
		transportAdapter.Mover()
	}
}
