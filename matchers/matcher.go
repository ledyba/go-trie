package matchers

type Matcher interface {
	Match(string) bool
	Contains(string) bool
}
