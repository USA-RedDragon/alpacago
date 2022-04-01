package alpaca

import (
	"testing"
)

var conditions = NewObservingConditions(65535, true, "alpaca.observerly.com", "", -1, 0)

func TestNewObservingConditionsBaseURL(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = conditions.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsBaseURLForHost(t *testing.T) {
	var got string = conditions.Alpaca.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsClientID(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = conditions.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsTransactionID(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = conditions.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsDeviceNumber(t *testing.T) {
	conditions := NewObservingConditions(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = conditions.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewObservingConditionsSetConnectedOn(t *testing.T) {
	var err = conditions.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsIsConnectedOn(t *testing.T) {
	var got, err = conditions.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsSetConnectedOff(t *testing.T) {
	var err = conditions.SetConnected(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsIsConnectedOff(t *testing.T) {
	var got, err = conditions.IsConnected()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t wanted %t", got, want)
	}

	if rotator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", conditions.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetAveragePeriod(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetAveragePeriod()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewObservingConditionsGetCloudCoverage(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetCloudCover()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 100 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetDewPoint(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetDewPoint()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 || got > 1000 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}

func TestNewObservingConditionsGetHumidity(t *testing.T) {
	conditions.SetConnected(true)

	var got, err = conditions.GetHumidity()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 || got > 100 {
		t.Errorf("got %f", got)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}
