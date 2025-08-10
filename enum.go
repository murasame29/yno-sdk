package yno

type DeviceStatus string

const (
	DeviceStatusOnline        DeviceStatus = "Online"
	DeviceStatusOffline       DeviceStatus = "Offline"
	DeviceStatusCommunicating DeviceStatus = "Communicating"
	DeviceStatusProcessing    DeviceStatus = "Processing"
	DeviceStatusError         DeviceStatus = "Error"
)

type ExecuteCommandStatus string

const (
	ExecuteCommandStatusSuccess ExecuteCommandStatus = "SUCCESS"
	ExecuteCommandStatusFaileed ExecuteCommandStatus = "FAILED"
	ExecuteCommandStatusTimeout ExecuteCommandStatus = "TIMEOUT"
)

type ExitCode string

const (
	ExitCodeSuccess ExitCode = "SUCCESS"
	ExitCodeError   ExitCode = "ERROR"
	ExitCodeSkipped ExitCode = "SKIPPED"
)
