package whatsapp_cloud_api_go

type WhatsApp interface{}

type whatsApp struct {
	AccessToken    string
	ApiVersion     ApiVersion
	BaseUrl        string
	FromPhone      string
	HostDomainName string
	HttpProtocol   HttpProtocol
}

type HttpProtocol string
type ApiVersion string

func New(phone, token string, version ApiVersion, protocol HttpProtocol) WhatsApp {
	domain := "graph.facebook.com"
	url := string(protocol) + "://" + domain + "/" + string(version)
	return whatsApp{
		AccessToken:    token,
		ApiVersion:     version,
		BaseUrl:        url,
		FromPhone:      phone,
		HttpProtocol:   protocol,
		HostDomainName: domain,
	}
}

// add a and b
func Add(a int, b int) int {
	// return the sum
	return a + b
}
