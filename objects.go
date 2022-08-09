package whatsapp_cloud_api_go

type whatsApp struct {
	AccessToken    string
	ApiVersion     ApiVersion
	BaseUrl        string
	FromPhone      string
	HostDomainName string
	HttpProtocol   HttpProtocol
	EndPoints      endPoints
}

type endPoints struct {
	Register,
	Deregister,
	Profile,
	Messages,
	Media,
	Request_code,
	Verify_code string
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/reference
type MessageObject struct {
	Audio            *MediaObject         `json:"audio,omitempty"`
	Contacts         *ContactObject       `json:"contacts,omitempty"`
	Document         *MediaObject         `json:"document,omitempty"`
	Image            *MediaObject         `json:"image,omitempty"`
	Interactive      *InteractiveObject   `json:"interactive,omitempty"`
	Location         *LocationObject      `json:"location,omitempty"`
	MessagingProduct MessagingProductType `json:"messagingProduct,omitempty"`
	PreviewUrl       bool                 `json:"preview_url,omitempty"`
	RecipientType    RecipientType        `json:"recipient_type,omitempty"`
	Status           string               `json:"status,omitempty"`
	Sticker          *MediaObject         `json:"sticker,omitempty"`
	Template         *TemplateObject      `json:"template,omitempty"`
	Text             *TextObject          `json:"text,omitempty"`
	To               string               `json:"to,omitempty"`
	Type             MessageType          `json:"type,omitempty"`
}

type ContactUrlObject struct {
	Url  string         `json:"url,omitempty"`
	Type ContactUrlType `json:"type,omitempty"`
}

type ContactOrgObject struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type ContactNameObject struct {
	FormattedName string `json:"formattedName,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	MiddleName    string `json:"middleName,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

type ContactAddressObject struct {
	Street      string             `json:"street,omitempty"`
	City        string             `json:"city,omitempty"`
	State       string             `json:"state,omitempty"`
	Zip         string             `json:"zip,omitempty"`
	Country     string             `json:"country,omitempty"`
	CountryCode string             `json:"countryCode,omitempty"`
	Type        ContactAddressType `json:"type,omitempty"`
}

type ContactEmailObject struct {
	Email string           `json:"email,omitempty"`
	Type  ContactEmailType `json:"type,omitempty"`
}

type ContactPhoneObject struct {
	WA_ID string           `json:"wA_ID,omitempty"`
	Phone string           `json:"phone,omitempty"`
	Type  ContactPhoneType `json:"type,omitempty"`
}

type ContactObject struct {
	Address  *ContactAddressObject `json:"address,omitempty"`
	Birthday string                `json:"birthday,omitempty"`
	Emails   []*ContactEmailObject `json:"emails,omitempty"`
	Name     *ContactNameObject    `json:"name,omitempty"`
	Org      *ContactOrgObject     `json:"org,omitempty"`
	Phones   []*ContactPhoneObject `json:"phones,omitempty"`
	Urls     []*ContactUrlObject   `json:"urls,omitempty"`
}

type InteractiveObjectBodyObject struct {
	Text string `json:"text,omitempty"`
}

type InteractiveObjectFooterObject struct {
	Text string `json:"text,omitempty"`
}

type InteractiveObject struct {
	Action string                         `json:"action,omitempty"`
	Type   InteractiveType                `json:"type,omitempty"`
	Header *HeaderObject                  `json:"header,omitempty"`
	Body   *InteractiveObjectBodyObject   `json:"body,omitempty"`
	Footer *InteractiveObjectFooterObject `json:"footer,omitempty"`
}

type ButtonObject struct {
	ID    string     `json:"id,omitempty"`
	Title string     `json:"title,omitempty"`
	Type  ButtonType `json:"type,omitempty"`
}

type ActionObject struct {
	Button   string           `json:"button,omitempty"`
	Buttons  []*ButtonObject  `json:"buttons,omitempty"`
	Sections []*SectionObject `json:"sections,omitempty"`
}

type HeaderObject struct {
	Type     HeaderType   `json:"type,omitempty"`
	Text     string       `json:"text,omitempty"`
	Image    *MediaObject `json:"image,omitempty"`
	Document *MediaObject `json:"document,omitempty"`
	Video    *MediaObject `json:"video,omitempty"`
}

type SectionRowObject struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type SectionObject struct {
	Title string              `json:"title,omitempty"`
	Rows  []*SectionRowObject `json:"rows,omitempty"`
}

type LocationObject struct {
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
}

type MediaObject struct {
	ID       string `json:"id,omitempty"`
	Link     string `json:"link,omitempty"`
	Caption  string `json:"caption,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type TextObject struct {
	Body       string `json:"body,omitempty"`
	PreviewUrl bool   `json:"previewUrl,omitempty"`
}

type ButtonParameterObject struct {
	Type    ButtonParameterType `json:"type,omitempty"`
	Payload interface{}         `json:"payload,omitempty"`
	Text    string              `json:"text,omitempty"`
}

type ComponentObject struct {
	Type       ComponentType      `json:"type,omitempty"`
	SubType    ComponentSubType   `json:"subType,omitempty"`
	Parameters []*ParameterObject `json:"parameters,omitempty"`
	Index      string             `json:"index,omitempty"`
}

type CurrencyObject struct {
	FallbackValue string `json:"fallbackValue,omitempty"`
	Code          string `json:"code,omitempty"`
	Amount1000    string `json:"amount1000,omitempty"`
}

type DateTimeObject struct {
	FallbackValue string `json:"fallbackValue,omitempty"`
}

type ParameterObject struct {
	Text     string          `json:"text,omitempty"`
	Type     ParameterType   `json:"type,omitempty"`
	Currency *CurrencyObject `json:"currency,omitempty"`
	DateTime *DateTimeObject `json:"dateTime,omitempty"`
	Image    *MediaObject    `json:"image,omitempty"`
	Document *MediaObject    `json:"document,omitempty"`
	Video    *MediaObject    `json:"video,omitempty"`
}

type TemplateObject struct {
	Name       string             `json:"name,omitempty"`
	Language   string             `json:"language,omitempty"`
	Components []*ComponentObject `json:"components,omitempty"`
	Namespace  string             `json:"namespace,omitempty"`
}
