package builder_test

import (
	"testing"

	builder "github.com/gsaldanab/design-pattern/3.builder"
)

func TestSender_BuildMessage(t *testing.T) {
	// json := &builder.JSONMessageBuilder{}
	xmltype := &builder.XMLMessageBuilder{}
	sender := &builder.Sender{}

	sender.SetBuilder(xmltype)
	msg, err := sender.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("ocurrio un error: %s", err.Error())
	}

	t.Log(msg.Format, string(msg.Body))
}
