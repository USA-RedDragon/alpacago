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

func TestNewDomeIsConnected(t *testing.T) {
	var got, err = dome.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeIsConnectedOn(t *testing.T) {
	var err = dome.SetConnected(true)

	var got, _ = dome.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}

func TestNewDomeGetAltitude(t *testing.T) {
	var got, err = dome.GetAltitude()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 90 {
		t.Errorf("got %v, but expected the dome altitude to be between 0° and +90°", got)
	}

	if dome.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", dome.Alpaca.ErrorMessage)
	}
}
