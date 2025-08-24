package yno

import (
	"context"
	"fmt"

	"github.com/murasame29/yno-sdk/client"
)

type SearchRouterRequest struct {
	PageSize  *int               `json:"PageSize,omitempty"`
	Query     *SearchRouterQuery `json:"Query,omitempty"`
	PageToken *string            `json:"PageToken,omitempty"`
}

func (p *SearchRouterRequest) Validate() error {
	if p.PageSize != nil {
		if *p.PageSize < 5 || *p.PageSize < 100 {
			return ValidateErrorNotMatch{"PageSize", " 5 <= x <= 100"}
		}
	}

	return nil
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

func (p RouterAssignedObject) Validate() error {
	if p.AssignedLabels != nil && len(p.AssignedLabels) == 0 {
		return ValidateErrorNotMatch{"AssignedLabels", "not empty"}
	}

	if p.AssignedUsers != nil && len(p.AssignedUsers) == 0 {
		return ValidateErrorNotMatch{"AssignedUsers", "not empty"}
	}

	return nil
}

type SearchRouterResponse struct {
	Meta     MetaData                 `json:"Meta"`
	Data     SearchRouterResponseData `json:"Data"`
	Warnings []Warning                `json:"Warnings"`
}

type SearchRouterResponseData struct {
	NextPageToken string                 `json:"NextPageToken"`
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

func (c *YNOClient) SearchRotuer(ctx context.Context, requestBody *SearchRouterRequest, opts ...OptionFunc) (*SearchRouterResponse, error) {
	if err := requestBody.Validate(); err != nil {
		return nil, err
	}

	var clientOpts []client.Option
	for _, optFunc := range opts {
		clientOpts = optFunc(clientOpts)
	}

	var responseBody SearchRouterResponse
	err := c.client.Post(ctx, "routers/_search", requestBody, &responseBody, clientOpts...)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (c *YNOClient) UpdateRotuer(ctx context.Context, serialNumber string, requestBody *RouterAssignedObject, opts ...OptionFunc) (*UpdateRotuerResponse, error) {
	if err := requestBody.Validate(); err != nil {
		return nil, err
	}

	var clientOpts []client.Option
	for _, optFunc := range opts {
		clientOpts = optFunc(clientOpts)
	}

	var responseBody UpdateRotuerResponse
	err := c.client.Put(ctx, fmt.Sprintf("routers/%s", serialNumber), requestBody, &responseBody, clientOpts...)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
