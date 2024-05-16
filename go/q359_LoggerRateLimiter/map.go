package q359_LoggerRateLimiter

type Logger struct {
    m map[string]int      
}

func Constructor() Logger {
    return Logger{
        m: map[string]int{},
    }
}

func (this *Logger) shouldPrintMessage(timestamp int, message string) bool {
    if timestamp < this.m[message] {
		return false
	}

	this.m[message] = timestamp + 10
	return true
}