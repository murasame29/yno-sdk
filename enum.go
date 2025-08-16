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

type DeviceStatType string

const (
	// MemoryUtilization: メモリ使用率
	DeviceStatTypeMemoryUtilization DeviceStatType = "MemoryUtilization"
	// CpuUtilization: CPU使用率
	DeviceStatTypeCpuUtilization DeviceStatType = "CpuUtilization"
	// AmountOfTraffic: トラフィック量
	DeviceStatTypeAmountOfTraffic DeviceStatType = "AmountOfTraffic"
	// NumberOfFastPathFlows: ファストパスフロー数
	DeviceStatTypeNumberOfFastPathFlows DeviceStatType = "NumberOfFastPathFlows"
	// NumberOfNatSessions: NATセッション数
	DeviceStatTypeNumberOfNatSessions DeviceStatType = "NumberOfNatSessions"
	// NumberOfDynamicFilterSessions: 動的フィルターのセッション数
	DeviceStatTypeNumberOfDynamicFilterSessions DeviceStatType = "NumberOfDynamicFilterSessions"
)

type Statistic string

const (
	StatisticTypeAverage Statistic = "Average"
	StatisticTypeMaximum Statistic = "Maximum"
)

type TrafficDirection string

const (
	TrafficDirectionIn  TrafficDirection = "In"
	TrafficDirectionOut TrafficDirection = "Out"
)

type IPVersion string

const (
	IPv4 IPVersion = "4"
	IPv6 IPVersion = "6"
)

type FormatOfAlarmNotificationEmailBody string

const (
	FormatOfAlarmNotificationEmailBodyJson FormatOfAlarmNotificationEmailBody = "JSON"
	FormatOfAlarmNotificationEmailBodyText FormatOfAlarmNotificationEmailBody = "Text"
)
