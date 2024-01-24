package helper

func EnsureNotEmpty(fields ...string) error {
	for _, field := range fields {
		if field == "" {
			return NewMicroMissingParameterErr(NotificationsServiceID)
		}
	}
	return nil
}
