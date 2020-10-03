package Config

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
	"plexToHue/Plex2U/Hue"
	"plexToHue/Tools"
)

func GetCredentials(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		cred := CreateCredentials()
		SetCredentials(path, cred)
		return cred
	}
	content, err := ioutil.ReadFile(path)
	Tools.PanicOnError(err)
	return string(content)

}

var CallbackWait = make(chan string)
var TempApi *echo.Echo

func CreateCredentials() string {
	fmt.Println("please visit the following link to link your philips hue account: ")
	fmt.Printf("https://api.meethue.com/oauth2/auth?clientid=%v&appid=%v&deviceid=dave&devicename=developer&state=caca&response_type=code", "jbmq2UrN1cIxd6JXeg9EreppyhaYhUEP", "plex2u")
	go waitForCallback()

	return <-CallbackWait
}

func waitForCallback() {
	TempApi = echo.New()
	TempApi.GET("/callback", func(c echo.Context) error {
		token := "jbmq2UrN1cIxd6JXeg9EreppyhaYhUEP"
		secret := "eK3ZvYejE2nAcdFT"
		Hue.AuthGetFromCode(c.Get("code").(string), &token, &secret)
		return c.String(http.StatusOK, "done!")
	})

	TempApi.Start(":8000")
}

func SetCredentials(path string, credential string) {
	ioutil.WriteFile(path, []byte(credential), 0777)
}
