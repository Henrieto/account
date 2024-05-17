package notifiers

type Msg struct {
	Subject string
	Body    string
}

type EmailPayload struct {
	From string
	To   []string
	Msg  *Msg
}
