package postman

// Collection represents a Postman Collection v2.1
type Collection struct {
	Info     Info       `json:"info"`
	Item     []Item     `json:"item"`
	Event    []Event    `json:"event,omitempty"`
	Variable []Variable `json:"variable,omitempty"`
	Auth     *Auth      `json:"auth,omitempty"`
}

// Info contains metadata about the collection
type Info struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Schema      string `json:"schema"`
	Version     string `json:"version,omitempty"`
}

// Item represents either a folder or a request
type Item struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Item        []Item     `json:"item,omitempty"`    // For folders
	Request     *Request   `json:"request,omitempty"` // For requests
	Event       []Event    `json:"event,omitempty"`
	Variable    []Variable `json:"variable,omitempty"`
}

// Request represents an HTTP request
type Request struct {
	Method      string   `json:"method"`
	Header      []Header `json:"header,omitempty"`
	Body        *Body    `json:"body,omitempty"`
	URL         URL      `json:"url"`
	Auth        *Auth    `json:"auth,omitempty"`
	Description string   `json:"description,omitempty"`
}

// URL represents a request URL
type URL struct {
	Raw      string     `json:"raw"`
	Protocol string     `json:"protocol,omitempty"`
	Host     []string   `json:"host,omitempty"`
	Path     []string   `json:"path,omitempty"`
	Query    []Query    `json:"query,omitempty"`
	Variable []Variable `json:"variable,omitempty"`
}

// Header represents an HTTP header
type Header struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
}

// Body represents the request body
type Body struct {
	Mode       string       `json:"mode"` // raw, urlencoded, formdata, file, graphql
	Raw        string       `json:"raw,omitempty"`
	Options    *BodyOptions `json:"options,omitempty"`
	URLEncoded []KeyValue   `json:"urlencoded,omitempty"`
	FormData   []KeyValue   `json:"formdata,omitempty"`
}

// BodyOptions contains options for the body
type BodyOptions struct {
	Raw *RawOptions `json:"raw,omitempty"`
}

// RawOptions contains options for raw body
type RawOptions struct {
	Language string `json:"language,omitempty"` // json, text, javascript, html, xml
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	Type        string `json:"type,omitempty"` // text, file (for formdata)
}

// Query represents a URL query parameter
type Query struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
}

// Variable represents a variable
type Variable struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Auth represents authentication configuration
type Auth struct {
	Type   string         `json:"type"` // apikey, bearer, basic, oauth2, etc.
	Bearer []AuthKeyValue `json:"bearer,omitempty"`
	APIKey []AuthKeyValue `json:"apikey,omitempty"`
	Basic  []AuthKeyValue `json:"basic,omitempty"`
}

// AuthKeyValue represents auth key-value pairs
type AuthKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type,omitempty"`
}

// Event represents a pre-request or test script
type Event struct {
	Listen string `json:"listen"` // prerequest, test
	Script Script `json:"script"`
}

// Script represents executable code
type Script struct {
	Type string   `json:"type"` // text/javascript
	Exec []string `json:"exec"` // Array of script lines
}
