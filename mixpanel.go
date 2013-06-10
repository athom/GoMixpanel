// This package implements the Mixpanel API in Go
//
package gomixpanel

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var DefaultMixPanel = NewMixpanel()

func SetApiToken(apiToken string) {
	DefaultMixPanel.SetApiToken(apiToken)
}

func Track(name string, properties map[string]interface{}) {
	SendEvent(name, properties)
}

func TrackWithCallback(name string, properties map[string]interface{}, failCallBack func()) {
	ok, err := SendEvent(name, properties)
	if err != nil || !ok {
		return
	}
	failCallBack()
}

func SendEvent(name string, properties map[string]interface{}) (ok bool, err error) {
	return DefaultMixPanel.SendEvent(name, properties)
}

type Mixpanel struct {
	apiToken string
}

func NewMixpanel() *Mixpanel {
	return &Mixpanel{}
}

func (this *Mixpanel) ApiToken() string {
	return this.apiToken
}

func (this *Mixpanel) SetApiToken(apiToken string) *Mixpanel {
	this.apiToken = apiToken
	return this
}

func (this *Mixpanel) NewEvent() *MixpanelEvent {
	e := NewMixpanelEvent()
	e.SetMixpanel(this)
	return e
}
func (this *Mixpanel) Track(name string, properties map[string]interface{}) {
	this.SendEvent(name, properties)
	return
}

func (this *Mixpanel) TrackWithCallback(name string, properties map[string]interface{}, failCallBack func()) {
	ok, err := this.SendEvent(name, properties)
	if err != nil || !ok {
		return
	}
	failCallBack()
}

func (this *Mixpanel) SendEvent(name string, properties map[string]interface{}) (success bool, err error) {
	e := this.NewEvent()
	e.SetName(name)
	e.SetProperties(properties)
	return e.Send()
}

type MixpanelEvent struct {
	mixpanel   *Mixpanel
	name       string
	properties map[string]interface{}
	apiToken   string
}

func NewMixpanelEvent() *MixpanelEvent {
	m := &MixpanelEvent{}
	m.properties = map[string]interface{}{}
	return m
}

func (this *MixpanelEvent) Mixpanel() *Mixpanel {
	return this.mixpanel
}

func (this *MixpanelEvent) SetMixpanel(mixpanel *Mixpanel) *MixpanelEvent {
	this.mixpanel = mixpanel
	return this
}

func (this *MixpanelEvent) Name() string {
	return this.name
}

func (this *MixpanelEvent) SetName(name string) *MixpanelEvent {
	this.name = name
	return this
}

func (this *MixpanelEvent) Properties() map[string]interface{} {
	return this.properties
}

func (this *MixpanelEvent) SetProperties(properties map[string]interface{}) *MixpanelEvent {
	this.properties = properties
	return this
}

func (this *MixpanelEvent) SetProperty(name string, value interface{}) *MixpanelEvent {
	this.properties[name] = value
	return this
}

func (this *MixpanelEvent) Send() (success bool, err error) {
	m := map[string]interface{}{}
	m["event"] = this.name

	p := map[string]interface{}{}
	for k, v := range this.properties {
		p[k] = v
	}
	p["token"] = this.mixpanel.ApiToken()
	m["properties"] = p

	//TODO handle errors
	data, err := json.Marshal(m)
	if err != nil {
		return false, err
	}

	encodedData := base64.URLEncoding.EncodeToString(data)
	resp, err := http.Get("http://api.mixpanel.com/track?data=" + encodedData)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body) == "1", err
}
