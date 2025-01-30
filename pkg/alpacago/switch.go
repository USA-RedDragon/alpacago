package alpacago

import "fmt"

type Switch struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewSwitch(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Switch {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	sw := Switch{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &sw
}

/*
GetDescription() common method to all ASCOM Alpaca compliant devices

@returns the description of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__description
*/
func (r *Switch) GetDescription() (string, error) {
	return r.Alpaca.GetDescription("switch", r.DeviceNumber)
}

/*
IsConnected() common method to all ASCOM Alpaca compliant devices

@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (r *Switch) IsConnected() (bool, error) {
	return r.Alpaca.GetBooleanResponse("switch", r.DeviceNumber, "connected")
}

/*
SetConnected() common method to all ASCOM Alpaca compliant devices

@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (r *Switch) SetConnected(connected bool) error {
	r.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", r.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", r.Alpaca.TransactionId),
	}

	return r.Alpaca.Put("switch", r.DeviceNumber, "connected", form)
}
