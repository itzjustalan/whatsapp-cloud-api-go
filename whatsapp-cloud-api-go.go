package whatsapp_cloud_api_go

// type WhatsApp interface{}

type whatsApp struct {
	AccessToken    string
	ApiVersion     ApiVersion
	BaseUrl        string
	FromPhone      string
	HostDomainName string
	HttpProtocol   HttpProtocol
	EndPoints      endPoints
}

type HttpProtocol string
type ApiVersion string

type endPoints struct {
	Register,
	Deregister,
	Profile,
	Messages,
	Media,
	Request_code,
	Verify_code string
}

func New(phone, token string, version ApiVersion, protocol HttpProtocol) *whatsApp {
	domain := "graph.facebook.com"
	url := string(protocol) + "://" + domain + "/" + string(version)
	return &whatsApp{
		AccessToken:    token,
		ApiVersion:     version,
		BaseUrl:        url,
		FromPhone:      phone,
		HttpProtocol:   protocol,
		HostDomainName: domain,
		EndPoints: endPoints{
			Register:     url + "/" + phone + "/register",
			Deregister:   url + "/" + phone + "/deregister",
			Media:        url + "/" + phone + "/media",
			Messages:     url + "/" + phone + "/messages",
			Request_code: url + "/" + phone + "/request_code",
			Verify_code:  url + "/" + phone + "/verify_code",
			Profile:      url + "/" + phone + "/whatsapp_business_profile",
		},
	}
}

// add a and b
func Add(a int, b int) int {
	// return the sum
	return a + b
}
