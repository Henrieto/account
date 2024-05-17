package auth

import "time"

func ExpiredOtp(otp_expiry time.Time) bool {
	return time.Now().After(otp_expiry)
}
