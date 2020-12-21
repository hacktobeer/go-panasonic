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
	Utoken   string `njson:"uToken"`
	Result   int    `njson:"result"`
	Language int    `njson:"language"`
}

// Groups is a set of grouped devices
type Groups struct {
	GroupCount int     `njson:"groupCount"`
	Groups     []Group `njson:"groupList"`
}

// Group defines a control group with devices
type Group struct {
	GroupID   int      `njson:"groupId"`
	GroupName string   `njson:"groupName"`
	Devices   []Device `njson:"deviceList"`
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
	ActualNanoe             int     `njson:"actualNanoe"`
	AirDirection            int     `njson:"airDirection"`
	AirQuality              int     `njson:"airQuality"`
	AirSwingLR              int     `njson:"airSwingLR"`
	AirSwingUD              int     `njson:"airSwingUD"`
	Defrosting              int     `njson:"defrosting"`
	DevGUID                 string  `njson:"devGuid"`
	DevRacCommunicateStatus int     `njson:"devRacCommunicateStatus"`
	EcoFunctionData         int     `njson:"ecoFunctionData"`
	EcoMode                 int     `njson:"ecoMode"`
	EcoNavi                 int     `njson:"ecoNavi"`
	Permission              int     `njson:"permission"`
	ErrorCode               int     `njson:"errorCode"`
	ErrorCodeStr            string  `njson:"errorCodeStr"`
	ErrorStatus             int     `njson:"errorStatus"`
	ErrorStatusFlg          bool    `njson:"errorStatusFlg"`
	FanAutoMode             int     `njson:"fanAutoMode"`
	FanSpeed                int     `njson:"fanSpeed"`
	HTTPErrorCode           int     `njson:"httpErrorCode"`
	Iauto                   int     `njson:"iAuto"`
	InsideTemperature       float64 `njson:"insideTemperature"`
	Nanoe                   int     `njson:"nanoe"`
	Online                  bool    `njson:"online"`
	Operate                 int     `njson:"operate"`       // on/off
	OperationMode           int     `njson:"operationMode"` // Mode (heat, dry, etc)
	OutsideTemperature      float64 `njson:"outTemperature"`
	PowerfulMode            bool    `njson:"powerfulMode"`
	TemperatureSet          float64 `njson:"temperatureSet"` // Temperature
	UpdateTime              int     `njson:"updateTime"`
}

//Device is Panasonic device
type Device struct {
	AirSwingLR         bool             `njson:"airSwingLR"`
	AutoMode           bool             `njson:"autoMode"`
	AutoTempMax        int              `njson:"autoTempMax"`
	AutoTempMin        int              `njson:"autoTempMin"`
	CoolMode           bool             `njson:"coolMode"`
	CoolTempMax        int              `njson:"coolTeampMax"`
	CoolTempMin        int              `njson:"coolTempMin"`
	DeviceGUID         string           `njson:"deviceGuid"`
	DeviceHashGUID     string           `njson:"deviceHashGuid"`
	DeviceModuleNumber string           `njson:"deviceModuleNumber"`
	DeviceName         string           `njson:"deviceName"`
	DeviceType         string           `njson:"deviceType"`
	DryMode            bool             `njson:"dryMode"`
	DryTempMax         int              `njson:"dryTempMax"`
	DryTempMin         int              `njson:"dryTempMin"`
	EcoFunction        bool             `njson:"ecoFunction"`
	EcoNavi            bool             `njson:"ecoNavi"`
	FanDirectionMode   int              `njson:"fanDirectionMode"`
	FanMode            bool             `njson:"fanMode"`
	FanSpeedMode       int              `njson:"fanSpeedMode"`
	HeatMode           bool             `njson:"heatMode"`
	HeatTempMax        int              `njson:"heatTempMax"`
	HeatTempMin        int              `njson:"heatTeampMin"`
	IautoX             bool             `njson:"iAutoX"`
	ModeAvlAutoMode    bool             `njson:"modeAvlList.autoMode"`
	ModeAvlFanMode     bool             `njson:"modeAvlList.fanMode"`
	Nanoe              bool             `njson:"nanoe"`
	QuietMode          bool             `njson:"quietMode"`
	SummerHouse        int              `njson:"summerHouse"`
	TemperatureUnit    bool             `njson:"temperatureUnit"`
	TimeStamp          int              `njson:"timestamp"`
	Parameters         DeviceParameters `njson:"parameters"`
}

// History is a list of HistoryEntry points with measurements
type History struct {
	EnergyConsumption  float64        `njson:"energyConsumption"`
	EstimatedCost      float64        `njson:"estimatedCost"`
	DeviceRegisterTime int            `njson:"deviceRegisterTime"`
	CurrencyUnit       string         `njson:"currencyUnit"`
	HistoryEntries     []HistoryEntry `njson:"historyDataList"`
}

// HistoryEntry is detailed data for a given day,week,month,year
type HistoryEntry struct {
	DataNumber         int     `njson:"dataNumber"`
	Consumption        float64 `njson:"consumption"`
	Cost               float64 `njson:"cost"`
	AverageSettingTemp float64 `njson:"averageSettingTemp"`
	AverageInsideTemp  float64 `njson:"averageInsideTemp"`
	AverageOutsideTemp float64 `njson:"averageOutsideTemp"`
}

// Command is basic command control structure
type Command struct {
	DeviceGUID string                  `json:"deviceGuid"`
	Parameters DeviceControlParameters `json:"parameters"`
}
