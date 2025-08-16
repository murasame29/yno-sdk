package yno

import "context"

type GetDeviceStatsRequest struct {
	Type         *DeviceStatType `json:"Type,omitempty"`
	SerialNumber *string         `json:"SerialNumber,omitempty"`
	StartTime    *int            `json:"StartTime,omitempty"`
	EndTime      *int            `json:"EndTime,omitempty"`
	Period       *int            `json:"Period,omitempty"`
	Statistics   *Statistic      `json:"Statistics,omitempty"`
	SearchAfter  *string         `json:"SearchAfter,omitempty"`
	Parameters   Parameter       `json:"Parameters,omitempty"`
}

func (r GetDeviceStatsRequest) Validate() error {
	if r.Type == nil {
		return ValidateErrorRequired{"Type"}
	}

	if r.SerialNumber == nil {
		return ValidateErrorRequired{"SerialNumber"}
	}

	if r.StartTime == nil {
		return ValidateErrorRequired{"StartTime"}
	}

	if *r.StartTime >= 0 {
		return ValidateErrorNotMatch{"StartTime", ">=0"}
	}

	if r.EndTime == nil {
		return ValidateErrorRequired{"EndTime"}
	}

	if *r.StartTime >= 0 {
		return ValidateErrorNotMatch{"StartTime", ">=0"}
	}

	if r.Statistics == nil {
		return ValidateErrorRequired{"Statistics"}
	}

	if r.Parameters != nil {
		return r.Parameters.Validate()
	}

	return nil
}

type Parameter interface {
	Validate() error
}

type CpuUtilizationParameter struct {
	CpuId *int `json:"CpuId,omitempty"`
}

func (p CpuUtilizationParameter) Validate() error {
	if p.CpuId == nil {
		return ValidateErrorRequired{"CpuId"}
	}

	return nil
}

type AmountOfTrafficParameter struct {
	Direction *TrafficDirection `json:"Direction,omitempty"`
	Interface *string           `json:"Interface,omitempty"`
}

func (p AmountOfTrafficParameter) Validate() error {
	if p.Direction == nil {
		return ValidateErrorRequired{"Direction"}
	}

	if p.Interface == nil {
		return ValidateErrorRequired{"Interface"}
	}

	return nil
}

type NumberOfFastPathFlowsParameter struct {
	IpVersion *IPVersion `json:"IpVersion,omitempty"`
}

func (p NumberOfFastPathFlowsParameter) Validate() error {
	if p.IpVersion == nil {
		return ValidateErrorRequired{"IpVersion"}
	}

	return nil
}

type GetDeviceStatsResponse struct {
	Meta MetaData `json:"Meta"`
}

type GetDeviceStatsResponseData struct {
	NextSearchAfter  string            `json:"NextSearchAfter"`
	DeviceStatistics []DeviceStatistic `json:"DeviceStatistics"`
}

type DeviceStatistic struct {
	Timestamp int `json:"Ts"`
	Value     int `json:"Val"`
}

func (c *ynoClient) GetDeviceStatistic(ctx context.Context, requestBody *GetDeviceStatsRequest) (*GetDeviceStatsResponse, error) {
	if err := requestBody.Validate(); err != nil {
		return nil, err
	}

	var responseBody GetDeviceStatsResponse
	err := c.client.Post(ctx, "routers/_search", requestBody, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
