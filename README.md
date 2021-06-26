# Usage

install package:
```shell
go get -u github.com/qiaocco/ghtrending
```

fetch trending repositories：
```go
func main() {
	trending := ghtrending.New()
	repos, err := trending.FetchRepos()
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Printf("repo=%#v\n", repo)
	}
}
```

fetch trending developers:
```go
developers, err := trending.FetchDevelopers()
```