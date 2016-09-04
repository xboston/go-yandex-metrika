package yametrikago

const (

	// @link https://github.com/pikhovkin/yametrikapy/blob/master/yametrikapy/core.py
	HOST        = "https://api-metrika.yandex.ru/"
	OAUTH_TOKEN = "https://oauth.yandex.ru/token"

	_COUNTERS   = "counters"
	_COUNTER    = "counter/%s"
	_GOALS      = _COUNTER + "/goals"
	_GOAL       = _COUNTER + "/goal/%d"
	_FILTERS    = _COUNTER + "/filters"
	_FILTER     = _COUNTER + "/filter/%d"
	_OPERATIONS = _COUNTER + "/operations"
	_OPERATION  = _COUNTER + "/operation/%d"
	_GRANTS     = _COUNTER + "/grants"
	_GRANT      = _COUNTER + "/grant/%s"
	_DELEGATES  = "delegates"
	_DELEGATE   = "delegate/%s"
	_ACCOUNTS   = "accounts"
	_ACCOUNT    = "account/%s"

	_STAT = "stat"

	_STAT_TRAFFIC          = _STAT + "/traffic"
	_STAT_TRAFFIC_SUMMARY  = _STAT_TRAFFIC + "/summary"
	_STAT_TRAFFIC_DEEPNESS = _STAT_TRAFFIC + "/deepness"
	_STAT_TRAFFIC_HOURLY   = _STAT_TRAFFIC + "/hourly"
	_STAT_TRAFFIC_LOAD     = _STAT_TRAFFIC + "/load"

	_STAT_SOURCES                  = _STAT + "/sources"
	_STAT_SOURCES_SUMMARY          = _STAT_SOURCES + "/summary"
	_STAT_SOURCES_SITES            = _STAT_SOURCES + "/sites"
	_STAT_SOURCES_SEARCH_ENGINES   = _STAT_SOURCES + "/search_engines"
	_STAT_SOURCES_PHRASES          = _STAT_SOURCES + "/phrases"
	_STAT_SOURCES_MARKETING        = _STAT_SOURCES + "/marketing"
	_STAT_SOURCES_DIRECT           = _STAT_SOURCES + "/direct"
	_STAT_SOURCES_DIRECT_SUMMARY   = _STAT_SOURCES_DIRECT + "/summary"
	_STAT_SOURCES_DIRECT_PLATFORMS = _STAT_SOURCES_DIRECT + "/platforms"
	_STAT_SOURCES_DIRECT_REGIONS   = _STAT_SOURCES_DIRECT + "/regions"
	_STAT_SOURCES_TAGS             = _STAT_SOURCES + "/tags"

	_STAT_CONTENT           = _STAT + "/content"
	_STAT_CONTENT_POPULAR   = _STAT_CONTENT + "/popular"
	_STAT_CONTENT_ENTRANCE  = _STAT_CONTENT + "/entrance"
	_STAT_CONTENT_EXIT      = _STAT_CONTENT + "/exit"
	_STAT_CONTENT_TITLES    = _STAT_CONTENT + "/titles"
	_STAT_CONTENT_URL_PARAM = _STAT_CONTENT + "/url_param"
	_STAT_CONTENT_USER_VARS = _STAT_SOURCES + "/user_vars"
	_STAT_CONTENT_ECOMMERCE = _STAT_CONTENT + "/ecommerce"

	_STAT_GEO = _STAT + "/geo"

	_STAT_DEMOGRAPHY            = _STAT + "/demography"
	_STAT_DEMOGRAPHY_AGE_GENDER = _STAT_DEMOGRAPHY + "/age_gender"
	_STAT_DEMOGRAPHY_STRUCTURE  = _STAT_DEMOGRAPHY + "/structure"

	_STAT_TECH             = _STAT + "/tech"
	_STAT_TECH_BROWSERS    = _STAT_TECH + "/browsers"
	_STAT_TECH_OS          = _STAT_TECH + "/os"
	_STAT_TECH_DISPLAY     = _STAT_TECH + "/display"
	_STAT_TECH_MOBILE      = _STAT_TECH + "/mobile"
	_STAT_TECH_FLASH       = _STAT_TECH + "/flash"
	_STAT_TECH_SILVERLIGHT = _STAT_TECH + "/silverlight"
	_STAT_TECH_DOTNET      = _STAT_TECH + "/dotnet"
	_STAT_TECH_JAVA        = _STAT_TECH + "/java"
	_STAT_TECH_COOKIES     = _STAT_TECH + "/cookies"
	_STAT_TECH_JAVASCRIPT  = _STAT_TECH + "/javascript"
)
