package yno

import (
	"context"
	"fmt"
	"strconv"
)

type CreateTaskRequest struct {
	Type       *string        `json:"Type,omitempty"`
	Timeout    *int           `json:"Timeout,omitempty"`
	Parameters *TaskParameter `json:"Parameters,omitempty"`
}

func (p CreateTaskRequest) Validate() error {
	if p.Type == nil {
		return ValidateErrorRequired{"Type"}
	}

	if p.Timeout != nil {
		if *p.Timeout < 60 || 1800 < *p.Timeout {
			return ValidateErrorNotMatch{"Timeout", " 60 <= x <= 1800"}
		}
	}

	if p.Parameters == nil {
		return ValidateErrorRequired{"Parameters"}
	}

	return p.Validate()
}

type TaskParameter struct {
	SerialNumbers []string `json:"SerialNumbers,omitempty"`
	Commands      []string `json:"Commands,omitempty"`
}

func (p TaskParameter) Validate() error {
	if p.SerialNumbers == nil {
		return ValidateErrorRequired{"SerialNumbers"}
	}

	if len(p.SerialNumbers) == 0 || len(p.SerialNumbers) > 1000 {
		return ValidateErrorNotMatch{"SerialNumbers", "0 <= x <= 1000"}
	}

	if p.Commands == nil {
		return ValidateErrorRequired{"Commands"}
	}

	if len(p.Commands) == 0 || len(p.Commands) > 100 {
		return ValidateErrorNotMatch{"Commands", "0 <= x <= 100"}
	}

	return nil
}

type CreateTaskResponse struct {
	Meta MetaData                `json:"Meta"`
	Data *CreateTaskResponseData `json:"Data,omitempty"`
}

type CreateTaskResponseData struct {
	TaskId string `json:"TaskId"`
}

type GetExecuteTaskQuery struct {
	PageSize  *int
	PageToken *string
}

func (p GetExecuteTaskQuery) Map() map[string]string {
	if p.PageSize == nil && p.PageToken == nil {
		return nil
	}

	m := make(map[string]string)
	if p.PageSize != nil {
		m["PageSize"] = strconv.Itoa(*p.PageSize)
	}

	if p.PageToken != nil {
		m["PageToken"] = *p.PageToken
	}

	return m
}

func (p GetExecuteTaskQuery) Validate() error {
	if p.PageSize != nil {
		if *p.PageSize < 5 || *p.PageSize < 100 {
			return ValidateErrorNotMatch{"PageSize", " 5 <= x <= 100"}
		}
	}

	return nil
}

type ExecuteTaskResponse struct {
	Meta     MetaData        `json:"Meta"`
	Data     ExecuteTaskData `json:"Data"`
	Warnings []Warning       `json:"Warnings,omitempty"`
}

type ExecuteTaskData struct {
	Type    string             `json:"Type"`
	Results ExecuteTaskResults `json:"Results"`
}

type ExecuteTaskResults struct {
	NextPageToken string             `json:"NextPageToken,omitempty"`
	Devices       []DeviceTaskResult `json:"Devices"`
}

type DeviceTaskResult struct {
	Status         ExecuteCommandStatus  `json:"Status"`
	SerialNumber   string                `json:"SerialNumber"`
	CommandResults []CommandResultDetail `json:"CommandResults,omitempty"`
}

type CommandResultDetail struct {
	Output   []string `json:"Output"`
	Command  string   `json:"Command"`
	ExitCode ExitCode `json:"ExitCode"`
}

func (c *YNOClient) CreateTask(ctx context.Context, requestBody *CreateTaskRequest) (*CreateTaskResponse, error) {
	if err := requestBody.Validate(); err != nil {
		return nil, err
	}

	var responseBody CreateTaskResponse
	err := c.client.Post(ctx, "tasks", requestBody, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (c *YNOClient) GetExecuteTask(ctx context.Context, taskID string, requestQuery *GetExecuteTaskQuery) (*ExecuteTaskResponse, error) {
	if err := requestQuery.Validate(); err != nil {
		return nil, err
	}

	var responseBody ExecuteTaskResponse
	err := c.client.Get(ctx, fmt.Sprintf("tasks/%s", taskID), requestQuery.Map(), &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
