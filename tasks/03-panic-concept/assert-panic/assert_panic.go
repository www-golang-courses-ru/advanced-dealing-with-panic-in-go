package coolassertlib

type TestingT interface {
	Errorf(format string, args ...interface{})
	Helper()
}

func AssertPanics(t TestingT, f func()) bool {
	// Реализуй меня.
	return false
}

func AssertNotPanics(t TestingT, f func()) bool {
	// Реализуй меня.
	return false
}
