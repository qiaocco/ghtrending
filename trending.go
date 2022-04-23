package ghtrending

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Repository struct {
	Author  string
	Name    string
	Link    string
	Desc    string
	Lang    string
	Stars   int
	Forks   int
	Add     int      // 新增星星数量
	BuiltBy []string // 贡献者
}

type Developer struct {
	Name        string
	Username    string
	PopularRepo string
	Desc        string
}

// Fetcher 设计接口
// 我想提供一个类型和一个创建该类型对象的方法，然后
//调用对象的FetchRepos()和FetchDevelopers()方法就可以获取仓库和开发者列表。
//但是我不希望用户了解这个类型的细节
type Fetcher interface {
	FetchRepos() ([]*Repository, error)
	FetchDevelopers() ([]*Developer, error)
}

type trending struct {
	opts options
}

func New(opts ...option) Fetcher {
	return &trending{
		opts: loadOptions(opts...),
	}
}

func loadOptions(opts ...option) options {
	o := options{
		GithubURL: "https://github.com",
	}
	fmt.Printf("option=%#v\n", o)
	for _, option := range opts {
		option(&o)
		fmt.Printf("option=%#v\n", o)
	}

	return o
}

func (t trending) FetchDevelopers() ([]*Developer, error) {
	resp, err := http.Get(fmt.Sprintf("%s/trending/developers/%s?since=%s",
		t.opts.GithubURL, t.opts.Language, t.opts.DateRange))

	if err != nil {
		log.Printf("get err: %v\n\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("NewDocumentFromReader err:%v\n", err)
		return nil, err
	}

	developers := make([]*Developer, 0, 10)
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		developer := &Developer{}
		developer.Username = strings.TrimSpace(s.Find("div>div>h1>a").Text())
		developer.Name = strings.TrimSpace(s.Find("div>div>p>a").Text())
		developer.PopularRepo = strings.TrimSpace(s.Find("div>div>article>h1>a").Text())
		developer.Desc = strings.TrimSpace(s.Find("div>div>article").Children().Last().Text())

		developers = append(developers, developer)
	})
	return developers, nil
}

func (t trending) FetchRepos() ([]*Repository, error) {
	resp, err := http.Get(fmt.Sprintf("%s/trending/%s?spoken_language_code=%s&since=%s",
		t.opts.GithubURL, t.opts.Language, t.opts.SpokenLang, t.opts.DateRange))
	if err != nil {
		log.Printf("get err: %v\n\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("NewDocumentFromReader err:%v\n", err)
		return nil, err
	}

	repos := make([]*Repository, 0, 10)
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		repo := &Repository{}
		// author name link
		titleSel := s.Find("h1 a")
		repo.Author = strings.TrimSpace(strings.Trim(titleSel.Find("span").Text(), "/\n"))
		repo.Name = strings.TrimSpace(titleSel.Contents().Last().Text())
		relativeLink, _ := titleSel.Attr("href")
		if len(relativeLink) > 0 {
			repo.Link = "https://github.com" + relativeLink
		}

		// desc
		repo.Desc = strings.TrimSpace(s.Find("p").Text())

		var langIdx, addIdx, builtByIdx int
		spanSel := s.Find("div>span")
		if spanSel.Size() == 2 {
			// language not exist
			langIdx = -1
			builtByIdx = 0
			addIdx = 1
		} else {
			langIdx = 0
			builtByIdx = 1
			addIdx = 2
		}

		// language
		if langIdx >= 0 {
			repo.Lang = strings.TrimSpace(spanSel.Eq(langIdx).Text())
		} else {
			repo.Lang = "unknown"
		}

		// add
		addParts := strings.SplitN(strings.TrimSpace(spanSel.Eq(addIdx).Text()), " ", 2)
		repo.Add, _ = strconv.Atoi(addParts[0])

		// builtby
		spanSel.Eq(builtByIdx).Find("a>img").Each(func(i int, img *goquery.Selection) {
			src, _ := img.Attr("src")
			repo.BuiltBy = append(repo.BuiltBy, src)
		})

		// stars forks
		aSel := s.Find("div>a")
		starStr := strings.TrimSpace(aSel.Eq(-2).Text())
		forkStr := strings.TrimSpace(aSel.Eq(-1).Text())
		repo.Stars, _ = strconv.Atoi(strings.Replace(starStr, ",", "", -1))
		repo.Forks, _ = strconv.Atoi(strings.Replace(forkStr, ",", "", -1))

		repos = append(repos, repo)
	})
	return repos, nil
}

func TrendingRepos(opts ...option) ([]*Repository, error) {
	return New(opts...).FetchRepos()
}

func TrendingDevelopers(opts ...option) ([]*Developer, error) {
	return New(opts...).FetchDevelopers()
}
