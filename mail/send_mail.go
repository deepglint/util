package mail

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(String string) string {
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func SendMail(from_addr string, from_alias string, from_pwd string, from_smtp string,
	to_addr string, to_alias string, title string, body string) (err error) {
	auth := smtp.PlainAuth(
		"",
		from_addr,
		from_pwd,
		from_smtp,
	)

	from := mail.Address{from_alias, from_addr}
	to := mail.Address{to_alias, to_addr}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err = smtp.SendMail(
		from_smtp+":25",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		return
	}
	return
}
