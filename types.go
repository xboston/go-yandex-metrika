package yametrikago

type ResultError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type ResultToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	UID         int    `json:"uid"`
}

type CounterList struct {
	Counters []Counter `json:"counters"`
	Rows     int       `json:"rows"`
}

type Counter struct {
	ID         string `json:"id"`
	CodeStatus string `json:"code_status"`
	Name       string `json:"name"`
	OwnerLogin string `json:"owner_login"`
	Permission string `json:"permission"`
	Site       string `json:"site"`
	Type       string `json:"type"`
}

type CounterData struct {
	Counter struct {
		Code        string `json:"code"`
		CodeOptions struct {
			Async         int `json:"async"`
			Clickmap      int `json:"clickmap"`
			Denial        int `json:"denial"`
			ExternalLinks int `json:"external_links"`
			Informer      struct {
				ColorArrow int    `json:"color_arrow"`
				ColorEnd   string `json:"color_end"`
				ColorStart string `json:"color_start"`
				ColorText  int    `json:"color_text"`
				Enabled    int    `json:"enabled"`
				Indicator  string `json:"indicator"`
				Size       int    `json:"size"`
				Type       string `json:"type"`
			} `json:"informer"`
			Params             int    `json:"params"`
			TrackHash          int    `json:"track_hash"`
			Ut                 int    `json:"ut"`
			Visor              int    `json:"visor"`
			WebvisorLoadPlayer string `json:"webvisor_load_player"`
			XMLSite            int    `json:"xml_site"`
		} `json:"code_options"`
		CodeStatus          string `json:"code_status"`
		FilterRobots        int    `json:"filter_robots"`
		ID                  string `json:"id"`
		MaxDetailedGoals    int    `json:"max_detailed_goals"`
		MaxGoals            int    `json:"max_goals"`
		MaxRetargetingGoals int    `json:"max_retargeting_goals"`
		Monitoring          struct {
			Emails           []string `json:"emails"`
			EnableMonitoring int      `json:"enable_monitoring"`
			EnableSms        int      `json:"enable_sms"`
			SmsAllowed       int      `json:"sms_allowed"`
			SmsTime          string   `json:"sms_time"`
		} `json:"monitoring"`
		Name                 string      `json:"name"`
		OwnerLogin           string      `json:"owner_login"`
		Permission           string      `json:"permission"`
		Site                 string      `json:"site"`
		TimeZoneName         string      `json:"time_zone_name"`
		Type                 string      `json:"type"`
		VisitThreshold       int         `json:"visit_threshold"`
		WebvisorArchEnabled  int         `json:"webvisor_arch_enabled"`
		WebvisorArchType     string      `json:"webvisor_arch_type"`
		WebvisorArchiveLinks interface{} `json:"webvisor_archive_links"`
	} `json:"counter"`
}

type StatTraffic struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Date1 string `json:"date1"`
	Date2 string `json:"date2"`
	Data  []struct {
		Date        string  `json:"date"`
		Denial      float64 `json:"denial"`
		Depth       float64 `json:"depth"`
		ID          string  `json:"id"`
		NewVisitors int64   `json:"new_visitors"`
		PageViews   int64   `json:"page_views"`
		VisitTime   int64   `json:"visit_time"`
		Visitors    int64   `json:"visitors"`
		Visits      int64   `json:"visits"`
	} `json:"data"`
	Totals struct {
		Denial      float64 `json:"denial"`
		Depth       float64 `json:"depth"`
		NewVisitors int64   `json:"new_visitors"`
		PageViews   int64   `json:"page_views"`
		VisitTime   int64   `json:"visit_time"`
		Visitors    int64   `json:"visitors"`
		Visits      int64   `json:"visits"`
	} `json:"totals"`
}

type Traffic struct {
	ID    string        `json:"id"`
	Data  []TrafficDate `json:"data"`
	Date1 string        `json:"date1"`
	Date2 string        `json:"date2"`
	Rows  int           `json:"rows"`
	Goals []interface{} `json:"goals"`
	Max   struct {
		Denial          float64 `json:"denial"`
		Depth           float64 `json:"depth"`
		NewVisitors     int     `json:"new_visitors"`
		NewVisitorsPerc int     `json:"new_visitors_perc"`
		PageViews       int     `json:"page_views"`
		VisitTime       int     `json:"visit_time"`
		Visitors        int     `json:"visitors"`
		Visits          int     `json:"visits"`
	} `json:"max"`
	Min struct {
		Denial          float64 `json:"denial"`
		Depth           float64 `json:"depth"`
		NewVisitors     int     `json:"new_visitors"`
		NewVisitorsPerc float64 `json:"new_visitors_perc"`
		PageViews       int     `json:"page_views"`
		VisitTime       int     `json:"visit_time"`
		Visitors        int     `json:"visitors"`
		Visits          int     `json:"visits"`
	} `json:"min"`
	Totals struct {
		Denial          float64 `json:"denial"`
		Depth           float64 `json:"depth"`
		NewVisitors     int     `json:"new_visitors"`
		NewVisitorsPerc int     `json:"new_visitors_perc"`
		PageViews       int     `json:"page_views"`
		VisitTime       int     `json:"visit_time"`
		Visitors        int     `json:"visitors"`
		Visits          int     `json:"visits"`
	} `json:"totals"`
}

type TrafficDate struct {
	Date            string  `json:"date"`
	Denial          float64 `json:"denial"`
	Depth           float64 `json:"depth"`
	ID              string  `json:"id"`
	NewVisitors     int     `json:"new_visitors"`
	NewVisitorsPerc float64 `json:"new_visitors_perc"`
	PageViews       int     `json:"page_views"`
	VisitTime       int     `json:"visit_time"`
	Visitors        int     `json:"visitors"`
	Visits          int     `json:"visits"`
	Wday            int     `json:"wday"`
}
