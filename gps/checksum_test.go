package gps

import (
	"testing"
)

func TestChecksum(t *testing.T) {

	validGPRMC := "$GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"
	invalidGPRMC := "$GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*76"

	if !Checksum(validGPRMC) {
		t.Errorf("%v\n", "Checksum not correct!")
	}

	if Checksum(invalidGPRMC) {
		t.Errorf("%v\n", "Invalid checksum passed check!")
	}

}
