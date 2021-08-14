package builder

import "encoding/json"

type JSONMessageBuilder struct {
	message Message
}

// SetRecipient asigna el receptor del mensaje
func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage asigna el mensaje a enviar
func (b *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// SetRecipient asigna el receptor del mensaje
func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "JSON"}, nil

}
