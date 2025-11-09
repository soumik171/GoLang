package auth

// private, cannot use outside the folder
func extractSession() string {
	return "loggedIN"
}

// call the private function inside as, it cann't be called from the outside
func GetSession() string {
	return extractSession()
}
