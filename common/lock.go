package common

// Lock switches start command lock flag
func Lock(status string) error {
	switch status {
	case "on":
		RobitState.StartLocked = true
	case "off":
		RobitState.StartLocked = false
	default:
		if RobitState.StartLocked {
			RobitState.StartLocked = false
		} else {
			RobitState.StartLocked = true
		}
	}

	return nil
}
