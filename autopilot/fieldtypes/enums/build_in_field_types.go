package enums

type BuildInFieldType string

const (
	BuildInFieldType_YES_NO             BuildInFieldType = "Twilio.YES_NO"
	BuildInFieldType_NUMBER             BuildInFieldType = "Twilio.NUMBER"
	BuildInFieldType_PHONE_NUMBER       BuildInFieldType = "Twilio.PHONE_NUMBER"
	BuildInFieldType_DATE               BuildInFieldType = "Twilio.DATE"
	BuildInFieldType_TIME               BuildInFieldType = "Twilio.TIME"
	BuildInFieldType_FIRST_NAME         BuildInFieldType = "Twilio.FIRST_NAME"
	BuildInFieldType_LAST_NAME          BuildInFieldType = "Twilio.LAST_NAME"
	BuildInFieldType_EMAIL              BuildInFieldType = "Twilio.EMAIL"
	BuildInFieldType_MONTH              BuildInFieldType = "Twilio.MONTH"
	BuildInFieldType_DAY_OF_WEEK        BuildInFieldType = "Twilio.DAY_OF_WEEK"
	BuildInFieldType_US_STATE           BuildInFieldType = "Twilio.US_STATE"
	BuildInFieldType_COUNTRY            BuildInFieldType = "Twilio.COUNTRY"
	BuildInFieldType_CITY               BuildInFieldType = "Twilio.CITY"
	BuildInFieldType_CURRENCY           BuildInFieldType = "Twilio.CURRENCY"
	BuildInFieldType_COUNTRY_ISO_ALPHA2 BuildInFieldType = "Twilio.COUNTRY_ISO_ALPHA2"
	BuildInFieldType_COUNTRY_ISO_ALPHA3 BuildInFieldType = "Twilio.COUNTRY_ISO_ALPHA3"
	BuildInFieldType_LANGUAGE           BuildInFieldType = "Twilio.LANGUAGE"
	BuildInFieldType_NUMBER_SEQUENCE    BuildInFieldType = "Twilio.NUMBER_SEQUENCE"
	BuildInFieldType_ALPHANUMERIC       BuildInFieldType = "Twilio.ALPHANUMERIC"
)

var buildInFields = map[string]bool{
	"Twilio.YES_NO":             true,
	"Twilio.NUMBER":             true,
	"Twilio.PHONE_NUMBER":       true,
	"Twilio.DATE":               true,
	"Twilio.TIME":               true,
	"Twilio.FIRST_NAME":         true,
	"Twilio.LAST_NAME":          true,
	"Twilio.EMAIL":              true,
	"Twilio.MONTH":              true,
	"Twilio.DAY_OF_WEEK":        true,
	"Twilio.US_STATE":           true,
	"Twilio.COUNTRY":            true,
	"Twilio.CITY":               true,
	"Twilio.CURRENCY":           true,
	"Twilio.COUNTRY_ISO_ALPHA2": true,
	"Twilio.COUNTRY_ISO_ALPHA3": true,
	"Twilio.LANGUAGE":           true,
	"Twilio.NUMBER_SEQUENCE":    true,
	"Twilio.ALPHANUMERIC":       true,
}

func IsBuildInFieldType(buildInFieldType string) bool {
	if _, ok := buildInFields[buildInFieldType]; ok {
		return true
	}
	return false
}
