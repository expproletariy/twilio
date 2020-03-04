package enums

type Language string

func (l Language) String() string {
	return string(l)
}

const (
	LanguageEn Language = "en-US"
)
