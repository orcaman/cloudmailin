package cloudmailin

import "encoding/json"

// Response instances contain the data of the JSON response form CloudMailin
type Response struct {
	Headers     Headers      `json:"headers,omitempty"`
	Envelope    Envelope     `json:"envelope,omitempty"`
	Plain       *string      `json:"plain,omitempty"`
	HTML        *string      `json:"html,omitempty"`
	ReplyPlain  *string      `json:"reply_plain,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Headers type holds Envelope data
type Headers struct {
	Received             []*string
	Date                 *string
	From                 *string
	To                   *string
	MessageID            *string `json:"Message-ID,omitempty"`
	Subject              *string
	MimeVersion          *string `json:"Mime-Version,omitempty"`
	ContentType          *string `json:"Content-Type,omitempty"`
	DKIMSignature        *string `json:"DKIM-Signature,omitempty"`
	XGoogleDKIMSignature *string `json:"X-Google-DKIM-Signature,omitempty"`
	XGmMessageState      *string `json:"X-Gm-Message-State,omitempty"`
	XReceived            *string `json:"X-Received,omitempty"`
}

// Envelope type holds Envelope data
type Envelope struct {
	To         *string   `json:"to,omitempty"`
	Recipients []*string `json:"recipients,omitempty"`
	From       *string   `json:"from,omitempty"`
	HeloDomain *string   `json:"helo_domain,omitempty"`
	RemoteIP   *string   `json:"remote_ip,omitempty"`
	SFP        SPF       `json:"spf,omitempty"`
}

// SPF type holds SPF data
type SPF struct {
	Result *string `json:"result,omitempty"`
	Domain *string `json:"domain,omitempty"`
}

// Attachment represents an email attachment
type Attachment struct {
	Content     []byte  `json:"content,omitempty"`
	Filename    *string `json:"file_name,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Size        *string `json:"size,omitempty"`
	Disposition *string `json:"disposition,omitempty"`
	ContentID   *string `json:"content_id,omitempty"`
}

// Parse parsed the incoming email data and returns a Parsed instance
func Parse(data []byte) (*Response, error) {
	var r *Response

	err := json.Unmarshal(data, &r)

	if err != nil {
		return nil, err
	}

	return r, nil
}
