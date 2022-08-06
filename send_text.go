package whatsapp_cloud_api_go

import (
	"bytes"
	"encoding/json"
	"errors"
	// "io/ioutil"
	// "log"
	"net/http"
)

func (wa *whatsApp) SendTextMessage(phone, text string, preview bool) error {
	d := make(map[string]interface{})
	t := make(map[string]interface{})

	d["messaging_product"] = "whatsapp"
	d["recipient_type"] = "individual"
	t["preview_url"] = preview
	d["type"] = "text"
	t["body"] = text
	d["to"] = phone
	d["text"] = t
	j, _ := json.Marshal(d)
	p := bytes.NewBuffer(j)

	res, err := http.Post(wa.EndPoints.Messages, "application/json", p)
	if err != nil || res.StatusCode != 200 {
		return errors.New("request failed")
	}
	// defer res.Body.Close()
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return errors.New("failed reading body")
	// }
	// log.Println(string(b), "bbbbbbbbb")
	return nil
}
