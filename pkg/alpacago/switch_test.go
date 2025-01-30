package alpacago

import (
	"testing"
)

var sw = NewSwitch(65535, false, "100.69.47.32", "", -1, 0)

func TestNewSwitchBaseURL(t *testing.T) {
	sw := NewSwitch(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = sw.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSwitchBaseURLForHost(t *testing.T) {
	var got string = sw.Alpaca.UrlBase
	var want string = "http://100.69.47.32"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSwitchClientID(t *testing.T) {
	sw := NewSwitch(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = sw.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSwitchTransactionID(t *testing.T) {
	sw := NewSwitch(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = sw.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSwitchDeviceNumber(t *testing.T) {
	sw := NewSwitch(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = sw.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewSwitchGetDescription(t *testing.T) {
	var got, err = sw.GetDescription()

	var want = "ASCOM SwitchV2 Simulator Driver"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewSwitchIsConnected(t *testing.T) {
	var got, err = sw.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if sw.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", sw.Alpaca.ErrorMessage)
	}
}

func TestNewSwitchSetConnected(t *testing.T) {
	var err = sw.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if sw.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", sw.Alpaca.ErrorMessage)
	}
}
