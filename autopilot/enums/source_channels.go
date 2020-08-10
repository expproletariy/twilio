package enums

type SourceChannel string

func (l SourceChannel) String() string {
	return string(l)
}

const (
	SourceChannelVoice           SourceChannel = "voice"
	SourceChannelSMS             SourceChannel = "sms"
	SourceChannelChat            SourceChannel = "chat"
	SourceChannelAlexa           SourceChannel = "alexa"
	SourceChannelGoogleAssistant SourceChannel = "google-assistant"
	SourceChannelSlack           SourceChannel = "slack"
)
