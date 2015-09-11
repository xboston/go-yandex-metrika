package yametrikago

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var httpClient = http.DefaultClient

type Metrika struct {
	ClientID     string
	ClientSecret string

	Username string
	Password string

	Token string
	Code  string

	UserAgent string
	Debug     bool
}

// тут еще задавать httpClient
func NewMetrika(clientId, clientSecret, username, password, token, code string) (metrika Metrika) {

	metrika = Metrika{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		Token:        token,
		Code:         code,
	}

	//metrika.UserAgent = "YametrikaGo"
	//metrika.Debug = false

	return
}

// активация отладки
func (this *Metrika) SetDebug(flag bool) {

	this.Debug = flag
}

// обработка результатов авторизации
func (this *Metrika) authorizeHandle(resultBody []byte) (err error) {

	var tokenResult ResultToken

	if err := json.Unmarshal(resultBody, &tokenResult); err == nil && tokenResult.AccessToken != "" {

		this.Token = tokenResult.AccessToken
		return nil
	}

	var tokenError ResultError

	if err := json.Unmarshal(resultBody, &tokenError); err == nil {
		return errors.New(tokenError.Error + ": " + tokenError.ErrorDescription)
	}

	return errors.New("Error authorizeHandle")
}

// авторизация
func (this *Metrika) Authorize() (err error) {

	if this.Token != "" {
		return
	}

	params := url.Values{"client_id": {this.ClientID}, "client_secret": {this.ClientSecret}}

	if this.Code != "" {
		// OAuth окно выдаёт этот самый code
		params.Add("grant_type", "authorization_code")
		params.Add("code", this.Code)
	} else {
		// или через логин - пароль
		params.Add("grant_type", "password")
		params.Add("username", this.Username)
		params.Add("password", this.Password)
	}

	resp, err := httpClient.PostForm(OAUTH_TOKEN, params)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = this.authorizeHandle(body)

	return err
}

func (this *Metrika) Auth() (err error) {

	if this.Token == "" {

		err = this.Authorize()
		return nil
	}

	return nil
}

func (this *Metrika) getURI(methodname, id string, params url.Values) (uri string) {

	// его надо в Header указывать
	params.Add("oauth_token", this.Token)

	uri = fmt.Sprintf("%s%s.json", HOST, methodname)

	if this.Debug {
		log.Println(uri)
	}

	if id != "" {
		uri = fmt.Sprintf(uri, id)
	}

	if len(params) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, params.Encode())
	}

	if this.Debug {
		log.Println(uri)
	}

	return uri
}

// список всех доступных счетчиков
func (this *Metrika) GetCounterList() (counterList CounterList, err error) {

	resp, err := httpClient.Get(this.getURI(_COUNTERS, "", url.Values{}))
	defer resp.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &counterList)

	return counterList, err
}

// полная информация о счетчике
func (this *Metrika) GetCounter(counterId string) (counterData CounterData, err error) {

	resp, err := httpClient.Get(this.getURI(_COUNTER, counterId, url.Values{}))
	defer resp.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &counterData)

	return counterData, err
}

// данные по трафику
func (this *Metrika) GetStatTrafficSummary(counterId string) (statTraffic StatTraffic, err error) {

	resp, err := http.Get(this.getURI(_STAT_TRAFFIC_SUMMARY, "", url.Values{"id": {counterId}}))
	defer resp.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &statTraffic)

	return statTraffic, err
}
