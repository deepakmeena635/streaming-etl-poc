package internal

type Sensor struct {
	Building             int       `json:"building"`
	Floor                int       `json:"floor"`
	Section              int       `json:"section"`
	RackId               int       `json:"rack-id"`
	Temp                 int       `json:"rack-temperature"`
	NearestCooling       []Cooling `json:"nearest-cooling"`
	RackTempContribution int       `json:"rack-temp-contribution"`
}

type Cooling struct {
	Building                int `json:"building"`
	Floor                   int `json:"floor"`
	Section                 int `json:"section"`
	LowestTemp              int `json:"lowest-temp"`
	CoolingSystemId         int `json:"hvac-id"`
	WattPerDegreeDifference int `json:"watt-per-degree-difference"`
}
