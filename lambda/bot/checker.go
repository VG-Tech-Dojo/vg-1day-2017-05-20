package bot

import (
	"regexp"

	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/lambda/model"
)

type (
	// Checker はmessageを受け取り、条件を満たすか判定するインターフェースです
	Checker interface {
		Check(*model.Message) bool
	}

	// RegexpChecker は 正規表現を満たす場合true、そうでない場合falseを返す構造体です
	RegexpChecker struct {
		regexp *regexp.Regexp
	}
)

// Check は正規表現を満たす場合true、そうでない場合falseを返します
func (c *RegexpChecker) Check(m *model.Message) bool {
	return c.regexp.MatchString(m.Body)
}

// NewRegexpChecker は新しいRegexpChecker構造体のポインタを返します
func NewRegexpChecker(pattern string) *RegexpChecker {
	r := regexp.MustCompile(pattern)
	return &RegexpChecker{
		regexp: r,
	}
}
