package domain

type Config struct {
	// Account             string
	Label               string
	Server              string
	Proxy               string
	Domain              string
	Username            string
	Password            string
	AuthID              string
	DisplayName         string
	DialingPrefix       string
	DialPlan            string
	HideCID             int
	VoicemailNumber     string
	Transport           string
	PublicAddr          string
	SRTP                string
	RegisterRefresh     int
	KeepAlive           int
	Publish             int
	ICE                 int
	AllowRewrite        int
	DisableSessionTimer int
}

func NewConfig() *Config {
	return &Config{
		HideCID:             0,
		RegisterRefresh:     300,
		KeepAlive:           15,
		Publish:             0,
		ICE:                 0,
		AllowRewrite:        0,
		DisableSessionTimer: 0,
	}
}
