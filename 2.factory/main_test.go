package factory_test

import (
	"testing"

	"github.com/gsaldanab/design-pattern/2.factory/factory"
)

func TestFactory_Msql(t *testing.T) {
	var s int = 1

	conn := factory.Factory(s)
	if conn == nil {
		t.Fatal("Motor no válido")
	}
	conn.Connect()
	conn.GetNow()
	conn.Close()
	t.Log("exito")
}

func TestFactory_Postgres(t *testing.T) {
	var s int = 2

	conn := factory.Factory(s)
	if conn == nil {
		t.Fatal("Motor no válido")
	}
	conn.Connect()
	conn.GetNow()
	conn.Close()
	t.Log("exito")
}

func TestFactory_MotorNoValido(t *testing.T) {
	var s int = 3

	conn := factory.Factory(s)
	if conn == nil {
		t.Log("Motor no válido")
	}

}
