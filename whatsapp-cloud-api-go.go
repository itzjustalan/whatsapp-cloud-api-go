package whatsapp_cloud_api_go

// create a new whatsapp instance
func New(phone, token string, version ApiVersion, protocol HttpProtocol) *whatsApp {
	url := string(protocol) + "://" + FBDOMAIN + "/" + string(version)
	return &whatsApp{
		AccessToken:    token,
		ApiVersion:     version,
		BaseUrl:        url,
		FromPhone:      phone,
		HttpProtocol:   protocol,
		HostDomainName: FBDOMAIN,
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
