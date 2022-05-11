package rogger

// Hook defines an interface to a log hook.
type Hook interface {
    // Run runs the hook with the event.
    Before(level LogLevel) map[string]string
    After(level LogLevel, message string)
}

func (l *Logger) SetHook(hook Hook) {
    l.hook = hook
}
