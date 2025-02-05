package alpacago

import "fmt"

type ObservingConditions struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewObservingConditions(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *ObservingConditions {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	conditions := ObservingConditions{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &conditions
}

/*
IsConnected() common method to all ASCOM Alpaca compliant devices

@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (c *ObservingConditions) IsConnected() (bool, error) {
	return c.Alpaca.GetBooleanResponse("observingconditions", c.DeviceNumber, "connected")
}

/*
SetConnected() common method to all ASCOM Alpaca compliant devices

@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (c *ObservingConditions) SetConnected(connected bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("observingconditions", c.DeviceNumber, "connected", form)
}

/*
GetDescription() common method to all ASCOM Alpaca compliant devices

@returns the description of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__description
*/
func (c *ObservingConditions) GetDescription() (string, error) {
	return c.Alpaca.GetDescription("observingconditions", c.DeviceNumber)
}

/*
GetAveragePeriod()

@returns the time period over which observations will be averaged.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__averageperiod
*/
func (c *ObservingConditions) GetAveragePeriod() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "averageperiod")
}

/*
GetCloudCover()

@returns the percentage of the sky obscured by cloud
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__cloudcover
*/
func (c *ObservingConditions) GetCloudCover() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "cloudcover")
}

/*
GetDewPoint()

@returns the atmospheric dew point at the observatory reported in °C.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__dewpoint
*/
func (c *ObservingConditions) GetDewPoint() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "dewpoint")
}

/*
GetHumidity()

@returns the atmospheric humidity (%) at the observatory
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__humidity
*/
func (c *ObservingConditions) GetHumidity() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "humidity")
}

/*
GetPressure()

@returns the atmospheric pressure in hectoPascals at the observatory's altitude - NOT reduced to sea level.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__pressure
*/
func (c *ObservingConditions) GetPressure() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "pressure")
}

/*
GetRainRate()

@returns the rain rate (mm/hour) at the observatory.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__rainrate
*/
func (c *ObservingConditions) GetRainRate() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "rainrate")
}

/*
GetSkyBrightness()

@returns the sky brightness (Lux) at the observatory
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__skybrightness
*/
func (c *ObservingConditions) GetSkyBrightness() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "skybrightness")
}

/*
GetSkyQuality()

@returns the sky quality (magnitudes per square arc second) at the observatory.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__skyquality
*/
func (c *ObservingConditions) GetSkyQuality() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "skyquality")
}

/*
GetSkyTemperature()

@returns the the sky temperature(°C) at the observatory.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__skytemperature
*/
func (c *ObservingConditions) GetSkyTemperature() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "skytemperature")
}

/*
GetSeeingStarFWHM()

@returns the seeing at the observatory measured as star full width half maximum (FWHM) in arc secs.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__starfwhm
*/
func (c *ObservingConditions) GetSeeingStarFWHM() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "starfwhm")
}

/*
GetTemperature()

@returns the temperature(°C) at the observatory.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__temperature
*/
func (c *ObservingConditions) GetTemperature() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "temperature")
}

/*
GetWindDirection()

@returns the the wind direction. The returned value must be between 0.0 and 360.0, interpreted according to the metereological standard,
where a special value of 0.0 is returned when the wind speed is 0.0. Wind direction is measured clockwise from north, through east,
where East=90.0, South=180.0, West=270.0 and North=360.0.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__winddirection
*/
func (c *ObservingConditions) GetWindDirection() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "winddirection")
}

/*
GetWindGust()

@returns the peak 3 second wind gust(m/s) at the observatory over the last 2 minutes.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__windgust
*/
func (c *ObservingConditions) GetWindGust() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "windgust")
}

/*
GetWindSpeed()

@returns the wind speed(m/s) at the observatory.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__windspeed
*/
func (c *ObservingConditions) GetWindSpeed() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "windspeed")
}

/*
SetRefresh()

@returns an error or nil, if forces the driver to immediately query its attached hardware to refresh sensor values.
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/put_observingconditions__device_number__refresh
*/
func (c *ObservingConditions) SetRefresh() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("observingconditions", c.DeviceNumber, "refresh", form)
}

/*
GetSensorDescription()

@returns a description of the sensor with the name specified in the SensorName parameter
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__sensordescription
*/
func (c *ObservingConditions) GetSensorDescription(sensorName string) (string, error) {
	url := c.Alpaca.getEndpoint("observingconditions", c.DeviceNumber, "sensordescription")

	querystring := fmt.Sprintf("sensorName=%v&%s", sensorName, c.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := c.Alpaca.Client.R().SetResult(&stringResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return "", err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		c.Alpaca.ErrorNumber = resp.StatusCode()
		c.Alpaca.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*stringResponse))

	return result.Value, nil
}

/*
GetTimeSinceLastUpdate()

@returns the time since the sensor specified in the SensorName parameter was last updated
@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__timesincelastupdate
*/
func (c *ObservingConditions) GetTimeSinceLastUpdate(sensorName string) (float64, error) {
	url := c.Alpaca.getEndpoint("observingconditions", c.DeviceNumber, "timesincelastupdate")

	querystring := fmt.Sprintf("sensorName=%v&%s", sensorName, c.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := c.Alpaca.Client.R().SetResult(&float64Response{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return 0., err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		c.Alpaca.ErrorNumber = resp.StatusCode()
		c.Alpaca.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*float64Response))

	return result.Value, nil
}
