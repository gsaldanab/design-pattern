package builder

import (
	"encoding/xml"
)

type XMLMessageBuilder struct {
	message Message
}

// SetRecipient asigna el receptor del mensaje
func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage asigna el mensaje a enviar
func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// SetRecipient asigna el receptor del mensaje
func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "XML"}, nil

}
