package verify

import (
	"4-order-api/configs"
	"4-order-api/pkg/req"
	"4-order-api/pkg/util"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
	"sync"
)

type VerifyHandlerDeps struct {
	*configs.VerifyMailConfig
}
type VerifyHandler struct {
	*configs.VerifyMailConfig
}

var mailVault = make(map[string]string)
var mu sync.Mutex

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		VerifyMailConfig: deps.VerifyMailConfig,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.VerifyMail())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendMailRequest](&w, r)
		if err != nil {
			return
		}
		hash := util.GenHash(body.Email)
		link := fmt.Sprintf("http://localhost:8081/verify/%s", hash)

		mu.Lock()
		mailVault[hash] = body.Email
		mu.Unlock()

		sendMail(body.Email, link)
		fmt.Println(mailVault)
		fmt.Println(body)
		fmt.Println(hash)
		fmt.Println(link)
		fmt.Printf("Send to %s", body.Email)
	}
}

func sendMail(to, link string) {
	e := email.NewEmail()
	config := configs.LoadVerifyConfig()
	e.From = config.Email
	e.To = []string{to}
	e.Subject = "Подтверждение почты"
	e.Text = []byte(fmt.Sprintf("Перейдите по ссылке: %s", link))
	err := e.Send(config.Address,
		smtp.PlainAuth("", config.Email, config.Password, config.Address))
	if err != nil {
		log.Println("Ошибка отправки письма")
	}
}

func (handler *VerifyHandler) VerifyMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyMailRequest](&w, r)
		if err != nil {
			return
		}
		hash := r.PathValue("hash")
		mu.Lock()
		_, exist := mailVault[hash]
		if exist {
			mu.Unlock()
			fmt.Fprintf(w, "false")
		} else {
			delete(mailVault, hash)
			mu.Unlock()
			fmt.Fprintf(w, "true")
		}
		fmt.Println(body)
		fmt.Println("Verify")
	}

}
