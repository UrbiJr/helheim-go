package main

import (
	"fmt"
	"log"
	"net/http"

	helheim_go "github.com/UrbiJr/helheim-go"
)

const YourApiKey = "2d276869-d44a-4e27-be05-19100afc0bd1"

func main() {
	helheimClient, err := helheim_go.ProvideClient(YourApiKey, false, true, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("helheim client initiated")

	versionResponse, err := helheimClient.Version()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(fmt.Sprintf("you are using helheim version %s", versionResponse.Version))

	options := helheim_go.CreateSessionOptions{
		Browser: helheim_go.BrowserOptions{
			Browser:  "chrome",
			Mobile:   false,
			Platform: "windows",
		},
		Captcha: helheim_go.CaptchaOptions{
			Provider: "vanaheim",
		},
	}
	session, err := helheimClient.NewSession(options)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("session response:")
	log.Println(session)

	balance, err := helheimClient.GetBalance()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("balance response:")
	log.Println(balance)

	_, err = session.SetProxy("http://127.0.0.1:8888")
	if err != nil {
		log.Println(err)
		return
	}

	kasadaOptions := helheim_go.KasadaHooksOptions{
		map[string]helheim_go.KasadaHookMethod{
			"mobile.api.prod.veve.me": helheim_go.KasadaHookMethod{
				map[string][]string{
					"POST": []string{
						"/graphql",
						"/api/auth/*",
					},
				},
			},
		},
	}
	_, err = session.SetKasadaHooks(kasadaOptions)
	if err != nil {
		log.Println(err)
		return
	}

	reqOpts := helheim_go.RequestOptions{
		Method: http.MethodGet,
		Url:    "https://mobile.api.prod.veve.me/149e9513-01fa-4fb0-aad4-566afd725d1b/2d206a39-8ed7-437e-a3be-862e0f06eea3/fp",
		Options: map[string]interface{}{
			"headers": map[string]string{
				"test": "gallina",
			},
		},
	}
	resp, err := session.Request(reqOpts)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request response status code:")
	log.Println(resp.Response.StatusCode)

	log.Println("request response:")
	log.Println(resp)

	err = session.Delete()

	if err != nil {
		log.Println(err)
		return
	}
}
