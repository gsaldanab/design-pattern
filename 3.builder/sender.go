package builder

type Sender struct {
	builder MessageBuilder
}

func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	return s.builder.SetRecipient(recipient).SetMessage(message).Build()
}
