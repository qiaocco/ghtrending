package ghtrending

import "testing"

func TestGhTrending(t *testing.T) {
	trend := New(WithDaily(), WithSpokenLanguageFull("chinese"))
	_, err := trend.FetchRepos()
	if err != nil {
		t.Logf("fetch failed, err=%+v", err)
		return
	}
	//for _, repo := range repos {
	//	//t.Logf("repo=%#v\n", repo)
	//}
}
