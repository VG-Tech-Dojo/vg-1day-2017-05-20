package bot

import (
	"regexp"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// messageを受け取り、条件を満たすか判定するインターフェース
	Checker interface {
		Check(*model.Message) bool
	}

	// 正規表現を満たす場合true、そうでない場合falseを返すchecker
	RegexpChecker struct {
		regexp *regexp.Regexp
	}
)

// 正規表現を満たす場合true、そうでない場合falseを返す
func (c *RegexpChecker) Check(m *model.Message) bool {
	return c.regexp.MatchString(m.Body)
}

// NewRegexpCheckerのインスタンス生成はこの関数を使う
func NewRegexpChecker(pattern string) *RegexpChecker {
	r := regexp.MustCompile(pattern)
	return &RegexpChecker{
		regexp: r,
	}
}
