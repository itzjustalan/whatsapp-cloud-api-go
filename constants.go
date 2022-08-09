package whatsapp_cloud_api_go

const FBDOMAIN = "graph.facebook.com"

type HttpProtocol string
type ApiVersion string
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

const (
	HTTP  HttpProtocol = "http"
	HTTPS              = "https"
)

const (
	V1_0  ApiVersion = "v1.0"
	V2_0             = "v2.0"
	V3_0             = "v3.0"
	V4_0             = "v4.0"
	V5_0             = "v5.0"
	V6_0             = "v6.0"
	V7_0             = "v7.0"
	V8_0             = "v8.0"
	V9_0             = "v9.0"
	V10_0            = "v10.0"
	V11_0            = "v11.0"
	V12_0            = "v12.0"
	V13_0            = "v13.0"
	V14_0            = "v14.0"
)

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
