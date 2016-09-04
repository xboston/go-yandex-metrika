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

var (
	httpClient = *http.DefaultClient
)

// Metrika - основеой объект работы с метрикой
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

// NewMetrikaFromToken - создание инстанции с токеном доступа
func NewMetrikaFromToken(token string) (metrika *Metrika) {

	return &Metrika{
		Token: token,
	}
}

// NewMetrikaFromCode - создание инстанции с кодом доступа
func NewMetrikaFromCode(code, clientID, clientSecret string) (metrika *Metrika) {

	metrika = &Metrika{
		Code: code,

		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return metrika
}

// NewMetrika - создание инстанции
func NewMetrika(clientID, clientSecret, username, password string) (metrika *Metrika) {

	metrika = &Metrika{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	}

	return metrika
}

// SetDebug - включение-выключение отладки
func (m *Metrika) SetDebug(flag bool) {

	m.Debug = flag
}

// обработка результатов авторизации
func (m *Metrika) authorizeHandle(resultBody []byte) (err error) {

	var tokenResult ResultToken

	if err := json.Unmarshal(resultBody, &tokenResult); err == nil && tokenResult.AccessToken != "" {

		m.Token = tokenResult.AccessToken
		return nil
	}

	var tokenError ResultError

	if err := json.Unmarshal(resultBody, &tokenError); err == nil {
		return errors.New(tokenError.Error + ": " + tokenError.ErrorDescription)
	}

	return errors.New("Error authorizeHandle")
}

// Authorize - авторизация
func (m *Metrika) Authorize() (err error) {

	// при имеющемся токене все остальные движения не нужны
	if m.Token != "" {
		return
	}

	params := url.Values{
		"client_id":     {m.ClientID},
		"client_secret": {m.ClientSecret},
	}

	if m.Code != "" {
		// OAuth окно выдаёт этот самый code
		params.Add("grant_type", "authorization_code")
		params.Add("code", m.Code)
	} else {
		// или через логин - пароль
		params.Add("grant_type", "password")
		params.Add("username", m.Username)
		params.Add("password", m.Password)
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

	return m.authorizeHandle(body)
}

func (m *Metrika) getURI(methodname, id string, params url.Values) (uri string) {

	// можно и в Header указывать
	params.Add("oauth_token", m.Token)

	uri = fmt.Sprintf("%s%s.json", HOST, methodname)

	if m.Debug {
		log.Println(uri)
	}

	if id != "" {
		uri = fmt.Sprintf(uri, id)
	}

	if len(params) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, params.Encode())
	}

	if m.Debug {
		log.Println(uri)
	}

	return uri
}

// GetCounterList - список всех доступных счетчиков
func (m *Metrika) GetCounterList() (counterList CounterList, err error) {

	resp, err := httpClient.Get(m.getURI(_COUNTERS, "", url.Values{}))
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

// GetCounter - полная информация о счетчике
func (m *Metrika) GetCounter(counterID string) (counterData CounterData, err error) {

	resp, err := httpClient.Get(m.getURI(_COUNTER, counterID, url.Values{}))
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

// GetStatTrafficSummary - данные по трафику
func (m *Metrika) GetStatTrafficSummary(counterID string) (statTraffic StatTraffic, err error) {

	resp, err := http.Get(m.getURI(_STAT_TRAFFIC_SUMMARY, "", url.Values{"id": {counterID}}))
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
