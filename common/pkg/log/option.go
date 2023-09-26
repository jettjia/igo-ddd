package log

type Option func(p *LoggerConfig)

func WithLogPath(logPath string) Option {
	return func(s *LoggerConfig) {
		s.LogPath = logPath
	}
}

func WithMaxSize(maxSize int) Option {
	return func(s *LoggerConfig) {
		s.MaxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(s *LoggerConfig) {
		s.MaxBackups = maxBackups
	}
}

func WithMaxAge(maxAge int) Option {
	return func(s *LoggerConfig) {
		s.MaxAge = maxAge
	}
}

func WithLogLevel(logLevel string) Option {
	return func(s *LoggerConfig) {
		s.LogLevel = logLevel
	}
}

func WithLogOut(logOut string) Option {
	return func(s *LoggerConfig) {
		s.LogOut = logOut
	}
}
