package email

import (
	"fmt"
	"net/url"

	"gopkg.in/mailgun/mailgun-go.v1"
)

const (
	welcomeSubject = "Welcome to Gallery!"
	resetSubject   = "Instructions for resetting your password."
	resetBaseURL   = "https://www.skmd.xyz/reset"
)

const welcomeText = `Hi there!
Welcome to our awesome website!

Thanks for signing up!
Domen`

const welcomeHTML = `Hi there!<br/>
Welcome to <a href="https://duckduckgo.com">DuckDuckGo.com</a>!`

const resetTextTmpl = `Hi there!

It appears that you have requested a password reset. If this was you, please follow the link below to update your password:

%s

If you are asked for a token, please use the following value:

%s

If you didn't request a password reset you can safely ignore this email.

skmd.xyz
`

const resetHTMLTmpl = `Hi there!<br>
<br>
It appears that you have requested a password reset. If this was you, please follow the link below to update your password:<br>
<br>
<a href="%s">%s</a><br>
<br>
If you are asked for a token, please use the following value:<br>
<br>
<b>%s</b><br>
<br>
If you didn't request a password reset you can safely ignore this email.<br>
<br>
skmd.xyz
`

func WithMailgun(domain, apiKey, publicKey string) ClientConfig {
	return func(c *Client) {
		mg := mailgun.NewMailgun(domain, apiKey, publicKey)
		c.mg = mg
	}
}

func WithSender(name, email string) ClientConfig {
	return func(c *Client) {
		c.from = buildEmail(name, email)
	}
}

type ClientConfig func(*Client)

func NewClient(opts ...ClientConfig) *Client {
	client := Client{
		from: "support@skmd.xyz",
	}
	for _, opt := range opts {
		opt(&client)
	}
	return &client
}

type Client struct {
	from string
	mg   mailgun.Mailgun
}

func (c *Client) Welcome(toName, toEmail string) error {
	message := mailgun.NewMessage(c.from, welcomeSubject, welcomeText, buildEmail(
		toName, toEmail,
	))
	message.SetHtml(welcomeHTML)
	_, _, err := c.mg.Send(message)
	return err
}

func (c *Client) ResetPw(toEmail, token string) error {
	v := url.Values{}
	v.Set("token", token)
	resetURL := resetBaseURL + "?" + v.Encode()
	resetText := fmt.Sprintf(resetTextTmpl, resetURL, token)
	message := mailgun.NewMessage(c.from, resetSubject, resetText, toEmail)
	resetHTML := fmt.Sprintf(resetHTMLTmpl, resetURL, resetURL, token)
	message.SetHtml(resetHTML)
	_, _, err := c.mg.Send(message)
	return err
}

func buildEmail(name, email string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}
