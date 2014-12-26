/*
Author: Aalekh Nigam@aalekh.nigam@gmail.com
*/
package gomailer 

import ("fmt"
		"reflect"
		"net/smtp"
		"strings"
		"net/mail"
		"encoding/base64"
		"log" )

func encodeRFC2047(String string) string{
// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func authentication (email , passcode , smtpserver string) smtp.Auth{

	auth := smtp.PlainAuth(
	"",
	email,
	passcode,
	smtpserver,
	)
	//fmt.Println(reflect.TypeOf(auth))
	return auth
}

type mailInfo struct {
	senderName string
	senderMail string
	recipientName string
	recipientMail string
	title string
	body string
	contentType string
	server string
	//port string
}

func sendMail(mai_two mailInfo, auth smtp.Auth) {

	from := mail.Address{mai_two.senderName, mai_two.senderMail}
	to := mail.Address{mai_two.recipientName, mai_two.recipientMail}
	title := mai_two.title
 
	body := mai_two.body;
 	fmt.Println("aaa2")
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	if mai_two.contentType == "plain" {
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
	} else if mai_two.contentType == "html" {
		header["Content-Type"] = "text/html; charset=\"utf-8\""
	}

	//header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
 	fmt.Println("aaa3")
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
 	//s := []string{":", mai_two.port};
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	fmt.Println("aaa4")
	err := smtp.SendMail(
		mai_two.server + ":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	fmt.Println("aaa5")
	if err != nil {
		log.Fatal(err)
	}
}

/*Example

func main() {
	//a := "aalekh.nigam@gmail.com"
	//smtpServer := "smtp.mandrillapp.com"
	//auth := smtp.PlainAuth(
	//"",
	//"aalekh.nigam@gmail.com",
	//"unique api key",
	//smtpServer,
//)

	auth_two := authentication ("aalekh.nigam@gmail.com" , "vI2qYIqqdPfpJnnSFGvhSA" , "smtp.mandrillapp.com")	
	fmt.Println(reflect.TypeOf(auth_two))
	mai := mailInfo{"Aalekh Nigam", "aalekh.nigam@gmail.com","Aalekhwa", "aalekh1993@rediffmail.com", "Hello Mail", "Yo bwoy", "plain", "smtp.mandrillapp.com"}
	fmt.Println("aaa")
	sendMail(mai, auth_two)
}