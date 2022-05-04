package alpacago

import "testing"

var dome = NewDome(65535, true, "alpaca.observerly.com", "", -1, 0)

func TestNewDomeBaseURL(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = dome.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeBaseURLForHost(t *testing.T) {
	var got string = dome.Alpaca.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeClientID(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = dome.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeTransactionID(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = dome.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDomeDeviceNumber(t *testing.T) {
	dome := NewDome(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = dome.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
