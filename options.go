package ghtrending

var spokenLangCode map[string]string

func init() {
	spokenLangCode = map[string]string{
		"chinese": "zh",
		"english": "en",
	}
}

// 选项模式/函数式模式 functional option
type options struct {
	GithubURL  string
	SpokenLang string
	Language   string // programming language
	DateRange  string
}

type option func(*options)

// WithDaily 定义三个DateRange
func WithDaily() option {
	return func(o *options) {
		o.DateRange = "daily"
	}
}

func WithWeekly() option {
	return func(o *options) {
		o.DateRange = "weekly"
	}
}

func WithMonthly() option {
	return func(o *options) {
		o.DateRange = "monthly"
	}
}

// WithDateRange 通用DateRange
func WithDateRange(dr string) option {
	return func(o *options) {
		o.DateRange = dr
	}
}

// WithLanguage 编程语言选项
func WithLanguage(lang string) option {
	return func(o *options) {
		o.Language = lang
	}
}

// WithSpokenLanguageCode 本地语言选项: 代码 cn
func WithSpokenLanguageCode(code string) option {
	return func(o *options) {
		o.SpokenLang = code
	}
}

// WithSpokenLanguageFull 本地语言选项: 国家 chinese
func WithSpokenLanguageFull(lang string) option {
	return func(o *options) {
		o.SpokenLang = spokenLangCode[lang]
	}
}

// WithURL 设置github的URL
func WithURL(url string) option {
	return func(o *options) {
		o.GithubURL = url
	}
}
