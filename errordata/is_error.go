package errordata

func IsError(err error) bool {
	if err == nil || err == SUCCESS {
		return true
	}
	return false
}
