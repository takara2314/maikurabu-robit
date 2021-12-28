package processes

// 操作ロックを切り替える
func Lock(status string, isLock *bool) error {
	switch status {
	case "on":
		*isLock = true
	case "off":
		*isLock = false
	default:
		if *isLock {
			*isLock = false
		} else {
			*isLock = true
		}
	}

	return nil
}
