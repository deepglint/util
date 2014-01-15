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
	to_labels []string, title string, body string) (err error) {
	auth := smtp.PlainAuth(
		"",
		from_addr,
		from_pwd,
		from_smtp,
	)

	from := mail.Address{from_alias, from_addr}
	var tos []mail.Address
	var toaddrs []string

	header := make(map[string]string)
	header["From"] = from.String()
	for i := 0; i < len(to_labels); i++ {
		splits := strings.Split(to_labels[i], ":")
		to := mail.Address{splits[1], splits[0]}
		if strings.EqualFold(header["To"], "") {
			header["To"] = to.String()
		} else {
			header["To"] = header["To"] + ";" + to.String()
		}
		tos = append(tos, to)
		toaddrs = append(toaddrs, to.Address)
	}
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
		toaddrs,
		[]byte(message),
	)
	if err != nil {
		return
	}
	return
}
