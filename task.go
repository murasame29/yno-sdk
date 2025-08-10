package yno

import (
	"context"
	"fmt"
)

type CreateTaskRequest struct {
	Type       string        `json:"Type"`
	Timeout    int           `json:"Timeout,omitempty"`
	Parameters TaskParameter `json:"Parameters"`
}

type TaskParameter struct {
	SerialNumbers []string `json:"SerialNumbers"`
	Commands      []string `json:"Commands"`
}

type CreateTaskResponse struct {
	Meta MetaData                `json:"Meta"`
	Data *CreateTaskResponseData `json:"Data,omitempty"`
}

type CreateTaskResponseData struct {
	TaskId string `json:"TaskId"`
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

func (c *ynoClient) CreateTask(ctx context.Context, requestBody *CreateTaskRequest) (*CreateTaskResponse, error) {
	var responseBody CreateTaskResponse
	err := c.client.Post(ctx, "routers/_search", requestBody, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (c *ynoClient) GetExecuteTask(ctx context.Context, taskID string) (*ExecuteTaskResponse, error) {
	var responseBody ExecuteTaskResponse
	err := c.client.Get(ctx, fmt.Sprintf("tasks/%s", taskID), &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
