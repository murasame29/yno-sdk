package yno

import (
	"context"
	"fmt"
)

type SearchRouterRequest struct {
	PageSize  int                `json:"PageSize,omitempty" validate:"gte=5,lte=100"`
	Query     *SearchRouterQuery `json:"Query,omitempty"`
	PageToken string             `json:"PageToken,omitempty"`
}

type SearchRouterQuery struct {
	Where *SearchRouterWhere `json:"Where,omitempty"`
}

type SearchRouterWhere struct {
	Equal               *SearchRouterEqualObject        `json:"$eq,omitempty"`
	PartialMatch        *SearchRouterPartialMatchObject `json:"$pm,omitempty"`
	In                  *SearchRouterInObject           `json:"$in,omitempty"`
	InArray             *RouterAssignedObject           `json:"$inArray,omitempty"`
	PartialMatchInArray *RouterAssignedObject           `json:"$pmInArray,omitempty"`
	And                 []SearchRouterWhere             `json:"$and,omitempty"`
	Or                  []SearchRouterWhere             `json:"$or,omitempty"`
}

type SearchRouterEqualObject struct {
	SerialNumber      string       `json:"SerialNumber,omitempty"`
	ModelName         string       `json:"ModelName,omitempty"`
	FirmwareRevision  string       `json:"FirmwareRevision,omitempty"`
	EndpointIpAddress string       `json:"EndpointIpAddress,omitempty"`
	DeviceDescription string       `json:"DeviceDescription,omitempty"`
	DeviceStatus      DeviceStatus `json:"DeviceStatus,omitempty"`
}

type SearchRouterPartialMatchObject struct {
	SerialNumber      string `json:"SerialNumber,omitempty"`
	ModelName         string `json:"ModelName,omitempty"`
	FirmwareRevision  string `json:"FirmwareRevision,omitempty"`
	EndpointIpAddress string `json:"EndpointIpAddress,omitempty"`
	DeviceDescription string `json:"DeviceDescription,omitempty"`
}

type SearchRouterInObject struct {
	SerialNumber      []string       `json:"SerialNumber,omitempty"`
	ModelName         []string       `json:"ModelName,omitempty"`
	FirmwareRevision  []string       `json:"FirmwareRevision,omitempty"`
	EndpointIpAddress []string       `json:"EndpointIpAddress,omitempty"`
	DeviceDescription []string       `json:"DeviceDescription,omitempty"`
	DeviceStatus      []DeviceStatus `json:"DeviceStatus,omitempty"`
}

type RouterAssignedObject struct {
	AssignedLabels []string `json:"AssignedLabels,omitempty"`
	AssignedUsers  []string `json:"AssignedUsers,omitempty"`
}

type SearchRouterResponse struct {
	Meta     MetaData                 `json:"Meta"`
	Data     SearchRouterResponseData `json:"Data"`
	Warnings []Warning                `json:"Warnings,omitempty"`
}

type SearchRouterResponseData struct {
	NextPageToken string                 `json:"NextPageToken,omitempty"`
	Routers       []RouterResponseRouter `json:"Routers"`
}

type RouterResponseRouter struct {
	DeviceStatus      string `json:"DeviceStatus"`
	ModelName         string `json:"ModelName"`
	FirmwareRevision  string `json:"FirmwareRevision"`
	DeviceDescription string `json:"DeviceDescription"`
	EndpointIPAddress string `json:"EndpointIpAddress"`
	SerialNumber      string `json:"SerialNumber"`
	RouterAssignedObject
}

type UpdateRotuerResponse struct {
	Meta MetaData             `json:"Meta"`
	Data RouterResponseRouter `json:"Data"`
}

func (c *ynoClient) SearchRotuer(ctx context.Context, requestBody *SearchRouterRequest) (*SearchRouterResponse, error) {
	var responseBody SearchRouterResponse
	err := c.client.Post(ctx, "routers/_search", requestBody, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (c *ynoClient) UpdateRotuer(ctx context.Context, serialNumber string, requestBody *RouterAssignedObject) (*UpdateRotuerResponse, error) {
	var responseBody UpdateRotuerResponse
	err := c.client.Put(ctx, fmt.Sprintf("routers/%s", serialNumber), requestBody, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
