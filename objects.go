package whatsapp_cloud_api_go

type ParameterType string
type ComponentType string
type ComponentSubType string
type ButtonParameterType string
type HeaderType string
type ButtonType string
type InteractiveType string
type ContactAddressType string
type ContactEmailType string
type ContactPhoneType string
type ContactUrlType string
type MessagingProductType string
type RecipientType string
type MessageType string

type MessageObject struct {
	Audio            MediaObject
	Contacts         ContactObject
	Document         MediaObject
	Image            MediaObject
	Interactive      InteractiveObject
	Location         LocationObject
	MessagingProduct MessagingProductType
	PreviewUrl       bool
	RecipientType    RecipientType
	Status           string
	Sticker          MediaObject
	Template         TemplateObject
	Text             TextObject
	To               string
	Type             MessageType
}

type ContactUrlObject struct {
	Url  string
	Type ContactUrlType
}

type ContactOrgObject struct {
	Company,
	Department,
	Title string
}

type ContactNameObject struct {
	FormattedName,
	FirstName,
	LastName,
	MiddleName,
	Suffix,
	Prefix string
}

type ContactAddressObject struct {
	Street,
	City,
	State,
	Zip,
	Country,
	CountryCode string
	Type ContactAddressType
}

type ContactEmailObject struct {
	Email string
	Type  ContactEmailType
}

type ContactPhoneObject struct {
	WA_ID string
	Phone string
	Type  ContactPhoneType
}

type ContactObject struct {
	Address  ContactAddressObject
	Birthday string
	Emails   []ContactEmailObject
	Name     ContactNameObject
	Org      ContactOrgObject
	Phones   []ContactPhoneObject
	Urls     []ContactUrlObject
}

type InteractiveObjectBodyObject struct {
	Text string
}

type InteractiveObjectFooterObject struct {
	Text string
}

type InteractiveObject struct {
	Action string
	Type   InteractiveType
	Header HeaderObject
	Body   InteractiveObjectBodyObject
	Footer InteractiveObjectFooterObject
}

type ButtonObject struct {
	ID    string
	Title string
	Type  ButtonType
}

type ActionObject struct {
	Button   string
	Buttons  []ButtonObject
	Sections []SectionObject
}

type HeaderObject struct {
	Type     HeaderType
	Text     string
	Image    MediaObject
	Document MediaObject
	Video    MediaObject
}

type SectionRowObject struct {
	ID,
	Title,
	Description string
}

type SectionObject struct {
	Title string
	Rows  []SectionRowObject
}

type LocationObject struct {
	Name      string
	Address   string
	Longitude float32
	Latitude  float32
}

type MediaObject struct {
	ID,
	Link,
	Caption,
	FileName,
	Provider string
}

type TextObject struct {
	Body       string
	PreviewUrl bool
}

type ButtonParameterObject struct {
	Type    ButtonParameterType
	Payload interface{}
	Text    string
}

type ComponentObject struct {
	Type       ComponentType
	SubType    ComponentSubType
	Parameters []ParameterObject
	Index      string
}

type CurrencyObject struct {
	FallbackValue string
	Code          string
	Amount1000    string
}

type DateTimeObject struct {
	FallbackValue string
}

type ParameterObject struct {
	Text     string
	Type     ParameterType
	Currency CurrencyObject
	DateTime DateTimeObject
	Image    MediaObject
	Document MediaObject
	Video    MediaObject
}

type TemplateObject struct {
	Name       string
	Language   string
	Components []ComponentObject
	Namespace  string
}

const (
	ParameterTypeText     ParameterType = "text"
	ParameterTypeCurrency               = "currency"
	ParameterTypeDateTime               = "date_time"
	ParameterTypeImage                  = "image"
	ParameterTypeDocument               = "document"
	ParameterTypeVideo                  = "video"
)

const (
	ComponentTypeHeader ComponentType = "header"
	ComponentTypeBody                 = "body"
	ComponentTypeButton               = "button"
)

const (
	ComponentSubTypeQuickReply ComponentSubType = "quick_reply"
	ComponentSubTypeUrl                         = "url"
)

const (
	ButtonParameterTypePayload ButtonParameterType = "payload"
	ButtonParameterTypeText                        = "text"
)

const (
	HeaderTypeText     HeaderType = "text"
	HeaderTypeVideo               = "video"
	HeaderTypeImage               = "image"
	HeaderTypeDocument            = "document"
)

const ButtonTypeReply ButtonType = "reply"

const (
	InteractiveTypeList   InteractiveType = "list"
	InteractiveTypeButton                 = "button"
)

const (
	ContactAddressTypeHome ContactAddressType = "HOME"
	ContactAddressTypeWork                    = "WORK"
)

const (
	ContactEmailTypeHome ContactEmailType = "HOME"
	ContactEmailTypeWork                  = "WORK"
)

const (
	ContactPhoneTypeHome ContactPhoneType = "HOME"
	ContactPhoneTypeWork                  = "WORK"
)

const (
	ContactUrlTypeHome ContactUrlType = "HOME"
	ContactUrlTypeWork                = "WORK"
)

const MessagingProductTypeWhatsApp MessagingProductType = "whatsapp"

const RecipientTypeIndividual RecipientType = "individual"

const (
	MessageTypeAudio       MessageType = "audio"
	MessageTypeContacts                = "contacts"
	MessageTypeDocument                = "document"
	MessageTypeImage                   = "image"
	MessageTypeInteractive             = "interactive"
	MessageTypeLocation                = "location"
	MessageTypeSticker                 = "sticker"
	MessageTypeTemplate                = "template"
	MessageTypeText                    = "text"
)
