package main

type Mail struct {
	Domain      string
	Host        string
	Port        string
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From       string
	FromName   string
	To         string
	Subject    string
	Attachment []string
	Data       any
	DataMap    map[string]any
}


func (m *Mail) SendSMTPMessage(msg Message) error{
	if msg.From == ""{
		msg.From = m.FromAddress
	}

	if msg.FromName == ""{
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message":msg.Data,
	}

	formattedMessage,err := m.buildHTMLMessage(msg)

}

func (m *Mail) buildHTMLMessage(msg string) (string error){
	templateToRender := "./templates/mail.html.gohtml"
}