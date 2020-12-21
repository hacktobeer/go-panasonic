package types

// Exported constants
const (
	URLServer       = "https://accsmart.panasonic.com"
	URLLogin        = "/auth/login"
	URLGroups       = "/device/group"
	URLDeviceStatus = "/deviceStatus/now/"
	URLHistory      = "/deviceHistoryData"
	URLControl      = "/deviceStatus/control"
	URLValidate1    = "/auth/agreement/status/1"
	SuccessResponse = `{"result":0}`
	FailureResponse = `{"result":1}`
)

// HistoryDataMode maps out the time intervals to fetch history data
var HistoryDataMode = map[string]int{
	"day":   0,
	"week":  1,
	"month": 2,
	"year":  3,
}

// Modes define the different AC modes the device can be in
var Modes = map[string]int{
	"auto": 0,
	"dry":  1,
	"cool": 2,
	"heat": 3,
	"fan":  4,
}

// ModesReverse define the different AC modes the device can be in
var ModesReverse = map[int]string{
	0: "auto",
	1: "dry",
	2: "cool",
	3: "heat",
	4: "fan",
}

// Operate defines if the AC is on or off
var Operate = map[int]string{
	0: "Off",
	1: "On",
}

// Session is a login session structure
type Session struct {
	Utoken   string `json:"uToken"`
	Result   int    `json:"result"`
	Language int    `json:"language"`
}

// Groups is a set of grouped devices
type Groups struct {
	GroupCount int     `json:"groupCount"`
	Groups     []Group `json:"groupList"`
}

// Group defines a control group with devices
type Group struct {
	GroupID   int      `json:"groupId"`
	GroupName string   `json:"groupName"`
	Devices   []Device `json:"deviceList"`
}

// DeviceControlParameters are the device control parameters
// used in Marshalling control commands
// We need to duplicate this with pointers to make sure the
// 'omitempty' parameter will not cancel out eg operate = 0
// when sending control commands to the unit
type DeviceControlParameters struct {
	ActualNanoe             *int     `json:"actualNanoe,omitempty"`
	AirDirection            *int     `json:"airDirection,omitempty"`
	AirQuality              *int     `json:"airQuality,omitempty"`
	AirSwingLR              *int     `json:"airSwingLR,omitempty"`
	AirSwingUD              *int     `json:"airSwingUD,omitempty"`
	Defrosting              *int     `json:"defrosting,omitempty"`
	DevGUID                 *string  `json:"devGuid,omitempty"`
	DevRacCommunicateStatus *int     `json:"devRacCommunicateStatus,omitempty"`
	EcoFunctionData         *int     `json:"ecoFunctionData,omitempty"`
	EcoMode                 *int     `json:"ecoMode,omitempty"`
	EcoNavi                 *int     `json:"ecoNavi,omitempty"`
	Permission              *int     `json:"permission,omitempty"`
	ErrorCode               *int     `json:"errorCode,omitempty"`
	ErrorCodeStr            *string  `json:"errorCodeStr,omitempty"`
	ErrorStatus             *int     `json:"errorStatus,omitempty"`
	ErrorStatusFlg          *bool    `json:"errorStatusFlg,omitempty"`
	FanAutoMode             *int     `json:"fanAutoMode,omitempty"`
	FanSpeed                *int     `json:"fanSpeed,omitempty"`
	HTTPErrorCode           *int     `json:"httpErrorCode,omitempty"`
	Iauto                   *int     `json:"iAuto,omitempty"`
	InsideTemperature       *float64 `json:"insideTemperature,omitempty"`
	Nanoe                   *int     `json:"nanoe,omitempty"`
	Online                  *bool    `json:"online,omitempty"`
	Operate                 *int     `json:"operate,omitempty"`       // Turn on/off
	OperationMode           *int     `json:"operationMode,omitempty"` // Set Mode (heat, dry, etc)
	OutsideTemperature      *float64 `json:"outTemperature,omitempty"`
	PowerfulMode            *bool    `json:"powerfulMode,omitempty"`
	TemperatureSet          *float64 `json:"temperatureSet,omitempty"` // Set Temperature
	UpdateTime              *int     `json:"updateTime,omitempty"`
}

// DeviceParameters are the current device parameters
// Used when UnMarshalling current device status
type DeviceParameters struct {
	ActualNanoe             int     `json:"actualNanoe"`
	AirDirection            int     `json:"airDirection"`
	AirQuality              int     `json:"airQuality"`
	AirSwingLR              int     `json:"airSwingLR"`
	AirSwingUD              int     `json:"airSwingUD"`
	Defrosting              int     `json:"defrosting"`
	DevGUID                 string  `json:"devGuid"`
	DevRacCommunicateStatus int     `json:"devRacCommunicateStatus"`
	EcoFunctionData         int     `json:"ecoFunctionData"`
	EcoMode                 int     `json:"ecoMode"`
	EcoNavi                 int     `json:"ecoNavi"`
	Permission              int     `json:"permission"`
	ErrorCode               int     `json:"errorCode"`
	ErrorCodeStr            string  `json:"errorCodeStr"`
	ErrorStatus             int     `json:"errorStatus"`
	ErrorStatusFlg          bool    `json:"errorStatusFlg"`
	FanAutoMode             int     `json:"fanAutoMode"`
	FanSpeed                int     `json:"fanSpeed"`
	HTTPErrorCode           int     `json:"httpErrorCode"`
	Iauto                   int     `json:"iAuto"`
	InsideTemperature       float64 `json:"insideTemperature"`
	Nanoe                   int     `json:"nanoe"`
	Online                  bool    `json:"online"`
	Operate                 int     `json:"operate"`       // on/off
	OperationMode           int     `json:"operationMode"` // Mode (heat, dry, etc)
	OutsideTemperature      float64 `json:"outTemperature"`
	PowerfulMode            bool    `json:"powerfulMode"`
	TemperatureSet          float64 `json:"temperatureSet"` // Temperature
	UpdateTime              int     `json:"updateTime"`
}

//Device is Panasonic device
type Device struct {
	AirSwingLR         bool             `json:"airSwingLR"`
	AutoMode           bool             `json:"autoMode"`
	AutoTempMax        int              `json:"autoTempMax"`
	AutoTempMin        int              `json:"autoTempMin"`
	CoolMode           bool             `json:"coolMode"`
	CoolTempMax        int              `json:"coolTeampMax"`
	CoolTempMin        int              `json:"coolTempMin"`
	DeviceGUID         string           `json:"deviceGuid"`
	DeviceHashGUID     string           `json:"deviceHashGuid"`
	DeviceModuleNumber string           `json:"deviceModuleNumber"`
	DeviceName         string           `json:"deviceName"`
	DeviceType         string           `json:"deviceType"`
	DryMode            bool             `json:"dryMode"`
	DryTempMax         int              `json:"dryTempMax"`
	DryTempMin         int              `json:"dryTempMin"`
	EcoFunction        int              `json:"ecoFunction"`
	EcoNavi            bool             `json:"ecoNavi"`
	FanDirectionMode   int              `json:"fanDirectionMode"`
	FanMode            bool             `json:"fanMode"`
	FanSpeedMode       int              `json:"fanSpeedMode"`
	HeatMode           bool             `json:"heatMode"`
	HeatTempMax        int              `json:"heatTempMax"`
	HeatTempMin        int              `json:"heatTeampMin"`
	IautoX             bool             `json:"iAutoX"`
	ModeAvlAutoMode    bool             `json:"modeAvlList.autoMode"`
	ModeAvlFanMode     bool             `json:"modeAvlList.fanMode"`
	Nanoe              bool             `json:"nanoe"`
	QuietMode          bool             `json:"quietMode"`
	SummerHouse        int              `json:"summerHouse"`
	TemperatureUnit    int              `json:"temperatureUnit"`
	TimeStamp          int              `json:"timestamp"`
	Parameters         DeviceParameters `json:"parameters"`
}

// History is a list of HistoryEntry points with measurements
type History struct {
	EnergyConsumption  float64        `json:"energyConsumption"`
	EstimatedCost      float64        `json:"estimatedCost"`
	DeviceRegisterTime string         `json:"deviceRegisterTime"`
	CurrencyUnit       string         `json:"currencyUnit"`
	HistoryEntries     []HistoryEntry `json:"historyDataList"`
}

// HistoryEntry is detailed data for a given day,week,month,year
type HistoryEntry struct {
	DataNumber         int     `json:"dataNumber"`
	Consumption        float64 `json:"consumption"`
	Cost               float64 `json:"cost"`
	AverageSettingTemp float64 `json:"averageSettingTemp"`
	AverageInsideTemp  float64 `json:"averageInsideTemp"`
	AverageOutsideTemp float64 `json:"averageOutsideTemp"`
}

// Command is basic command control structure
type Command struct {
	DeviceGUID string                  `json:"deviceGuid"`
	Parameters DeviceControlParameters `json:"parameters"`
}
