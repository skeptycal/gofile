package errorlogger

type loggerFunc interface {
	Error(args ...interface{})
}

// SetLoggerFunc allows setting of the logger function.
// The default is log.Error(err), which is compatible with
// the standard library log package and logrus.
//
// The function signature must be
//  func(args ...interface{}).
func (e *errorLogger) SetLoggerFunc(fn loggerFunc) {
	// switch v := fn.(type) {
	// case loggerFunc:
	// 	e.SetLoggerFunc(fn)
	// default:
	// 	log.Info(fmt.Errorf("%v is not a loggerFunc ... using default", v))
	// 	e.SetLoggerFunc(logrus.New())
	// }

	e.logFunc = fn
}
