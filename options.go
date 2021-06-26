package ghtrending

var spokenLangCode map[string]string

func init() {
	spokenLangCode = map[string]string{
		"chinese": "zh",
		"english": "en",
	}
}

type options struct {
	GithubURL  string
	SpokenLang string
	Language   string // programming language
	DateRange  string
}

type option func(*options)

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

func WithDateRange(dr string) option {
	return func(o *options) {
		o.DateRange = dr
	}
}

func WithLanguage(lang string) option {
	return func(o *options) {
		o.Language = lang
	}
}

func WithSpokenLanguageCode(code string) option {
	return func(o *options) {
		o.SpokenLang = code
	}
}

func WithSpokenLanguageFull(lang string) option {
	return func(o *options) {
		o.SpokenLang = spokenLangCode[lang]
	}
}

func WithURL(url string) option {
	return func(o *options) {
		o.GithubURL = url
	}
}
