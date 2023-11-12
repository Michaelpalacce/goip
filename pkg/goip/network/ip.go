package network

// GetPublicIp will fetch the public IP of the
// machine that is running goip
func GetPublicIp() ([]byte, error) {
	return GetBody("https://icanhazip.com")
}
