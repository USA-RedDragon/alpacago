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
func (s *Switch) GetDescription() (string, error) {
	return s.Alpaca.GetDescription("switch", s.DeviceNumber)
}

/*
IsConnected() common method to all ASCOM Alpaca compliant devices

@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (s *Switch) IsConnected() (bool, error) {
	return s.Alpaca.GetBooleanResponse("switch", s.DeviceNumber, "connected")
}

/*
SetConnected() common method to all ASCOM Alpaca compliant devices

@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (s *Switch) SetConnected(connected bool) error {
	s.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", s.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", s.Alpaca.TransactionId),
	}

	return s.Alpaca.Put("switch", s.DeviceNumber, "connected", form)
}

/*
GetMaxSwitch()

@returns the number of switch devices managed by this driver
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__maxswitch
*/
func (s *Switch) GetMaxSwitch() (int32, error) {
	return s.Alpaca.GetInt32Response("switch", s.DeviceNumber, "maxswitch")
}

/*
CanAsync()

@returns true if the given switch can operate asynchronously.
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__canasync
*/
func (s *Switch) CanAsync(switchID int32) (bool, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "canasync")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&booleanResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return false, err
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
}

/*
CanWrite()

@returns true if the given switch can be written to.
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__canwrite
*/
func (s *Switch) CanWrite(switchID int32) (bool, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "canwrite")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&booleanResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return false, err
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
}

/*
GetSwitch()

@returns true if the given switch is on.
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__getswitch
*/
func (s *Switch) GetSwitch(switchID int32) (bool, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "getswitch")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&booleanResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return false, err
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
}

/*
GetSwitchDescription()

@returns the description of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__getswitchdescription
*/
func (s *Switch) GetSwitchDescription(switchID int32) (string, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "getswitchdescription")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&stringResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return "", err
	}

	// Return the result:
	result := (resp.Result().(*stringResponse))

	return result.Value, nil
}

/*
GetSwitchName()

@returns the name of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__getswitchname
*/
func (s *Switch) GetSwitchName(switchID int32) (string, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "getswitchname")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&stringResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return "", err
	}

	// Return the result:
	result := (resp.Result().(*stringResponse))

	return result.Value, nil
}

/*
GetSwitchValue()

@returns the value of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__getswitchvalue
*/
func (s *Switch) GetSwitchValue(switchID int32) (float64, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "getswitchvalue")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&float64Response{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return 0, err
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}

/*
MinSwitchValue()

@returns the minimum value of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__minswitchvalue
*/
func (s *Switch) MinSwitchValue(switchID int32) (float64, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "minswitchvalue")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&float64Response{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return 0, err
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}

/*
MaxSwitchValue()

@returns the minimum value of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__maxswitchvalue
*/
func (s *Switch) MaxSwitchValue(switchID int32) (float64, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "maxswitchvalue")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&float64Response{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return 0, err
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}

/*
SetAsync()

@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/put_switch__device_number__setasync
*/

/*
SetAsyncValue()

@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/put_switch__device_number__setasyncvalue
*/

/*
SetSwitch()

@params state bool (the value to be set, true for on, false for off.)
@returns an error or nil, if nil the operation was successful
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/put_switch__device_number__setswitch
*/
func (s *Switch) SetSwitch(state bool) error {
	s.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Relative position to move in degrees from current Position.
		"State":               fmt.Sprintf("%t", state),
		"ClientID":            fmt.Sprintf("%d", s.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", s.Alpaca.TransactionId),
	}

	return s.Alpaca.Put("switch", s.DeviceNumber, "setswitch", form)
}

/*
SetSwitchName()

@params name string (the name to be set)
@returns an error or nil, if nil the operation was successful
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/put_switch__device_number__setswitchname
*/
func (s *Switch) SetSwitchName(name string) error {
	s.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Relative position to move in degrees from current Position.
		"Name":                fmt.Sprintf("%s", name),
		"ClientID":            fmt.Sprintf("%d", s.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", s.Alpaca.TransactionId),
	}

	return s.Alpaca.Put("switch", s.DeviceNumber, "setswitchname", form)
}

/*
SetSwitchValue()

@params value float64 (the value to be set, between MinSwitchValue and MaxSwitchValue.)
@returns an error or nil, if nil the operation was successful
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/put_switch__device_number__setswitchvalue
*/
func (s *Switch) SetSwitchValue(value float64) error {
	s.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Relative position to move in degrees from current Position.
		"Value":               fmt.Sprintf("%f", value),
		"ClientID":            fmt.Sprintf("%d", s.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", s.Alpaca.TransactionId),
	}

	return s.Alpaca.Put("switch", s.DeviceNumber, "setswitchvalue", form)
}

/*
StateChangeComplete()

@returns true if the state of the specified switch is changing.
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__statechangecomplete
*/
func (s *Switch) StateChangeComplete(switchID int32) (bool, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "statechangecomplete")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&booleanResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return false, err
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
}

/*
SwitchStep()

@returns the minimum value of the specified switch
@see https://ascom-standards.org/api/#/Switch%20Specific%20Methods/get_switch__device_number__switchstep
*/
func (s *Switch) SwitchStep(switchID int32) (float64, error) {
	url := s.Alpaca.getEndpoint("switch", s.DeviceNumber, "switchstep")

	querystring := fmt.Sprintf("Id=%d&%s", switchID, s.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := s.Alpaca.Client.R().SetResult(&float64Response{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return 0, err
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}
