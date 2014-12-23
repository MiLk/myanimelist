package myanimelist

import (
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func GetAnimeList(v *Result, username string) error {
	request := gorequest.New()
	_, body, errs := request.Get("http://myanimelist.net/malappinfo.php?u=" + username).
		End()

	if errs != nil {
		return errs[0]
	}

	err := xml.Unmarshal([]byte(body), v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	return nil
}
