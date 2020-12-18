package types

// Exported constants
const (
	URLServer       = "https://accsmart.panasonic.com"
	URLLogin        = "/auth/login"
	URLGroups       = "/device/group"
	URLDeviceStatus = "/deviceStatus/now/"
	URLHistory      = "/deviceHistoryData"
	URLControl      = "/deviceStatus/control"
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
	PactualNanoe             *int     `json:"actualNanoe,omitempty"`
	PairDirection            *int     `json:"airDirection,omitempty"`
	PairQuality              *int     `json:"airQuality,omitempty"`
	PairSwingLR              *int     `json:"airSwingLR,omitempty"`
	PairSwingUD              *int     `json:"airSwingUD,omitempty"`
	Pdefrosting              *int     `json:"defrosting,omitempty"`
	PdevGUID                 *string  `json:"devGuid,omitempty"`
	PdevRacCommunicateStatus *int     `json:"devRacCommunicateStatus,omitempty"`
	PecoFunctionData         *int     `json:"ecoFunctionData,omitempty"`
	PecoMode                 *int     `json:"ecoMode,omitempty"`
	PecoNavi                 *int     `json:"ecoNavi,omitempty"`
	Permission               *int     `json:"permission,omitempty"`
	PerrorCode               *int     `json:"errorCode,omitempty"`
	PerrorCodeStr            *string  `json:"errorCodeStr,omitempty"`
	PerrorStatus             *int     `json:"errorStatus,omitempty"`
	PerrorStatusFlg          *bool    `json:"errorStatusFlg,omitempty"`
	PfanAutoMode             *int     `json:"fanAutoMode,omitempty"`
	PfanSpeed                *int     `json:"fanSpeed,omitempty"`
	PhttpErrorCode           *int     `json:"httpErrorCode,omitempty"`
	PiAuto                   *int     `json:"iAuto,omitempty"`
	PinsideTemperature       *float64 `json:"insideTemperature,omitempty"`
	Pnanoe                   *int     `json:"nanoe,omitempty"`
	Ponline                  *bool    `json:"online,omitempty"`
	Poperate                 *int     `json:"operate,omitempty"`       // Turn on/off
	PoperationMode           *int     `json:"operationMode,omitempty"` // Set Mode (heat, dry, etc)
	PoutsideTemperature      *float64 `json:"outTemperature,omitempty"`
	PpowerfulMode            *bool    `json:"powerfulMode,omitempty"`
	PtemperatureSet          *float64 `json:"temperatureSet,omitempty"` // Set Temperature
	PupdateTime              *int     `json:"updateTime,omitempty"`
}

// DeviceParameters are the current device parameters
// Used when UnMarshalling current device status
type DeviceParameters struct {
	PactualNanoe             int     `njson:"actualNanoe"`
	PairDirection            int     `njson:"airDirection"`
	PairQuality              int     `njson:"airQuality"`
	PairSwingLR              int     `njson:"airSwingLR"`
	PairSwingUD              int     `njson:"airSwingUD"`
	Pdefrosting              int     `njson:"defrosting"`
	PdevGUID                 string  `njson:"devGuid"`
	PdevRacCommunicateStatus int     `njson:"devRacCommunicateStatus"`
	PecoFunctionData         int     `njson:"ecoFunctionData"`
	PecoMode                 int     `njson:"ecoMode"`
	PecoNavi                 int     `njson:"ecoNavi"`
	Permission               int     `njson:"permission"`
	PerrorCode               int     `njson:"errorCode"`
	PerrorCodeStr            string  `njson:"errorCodeStr"`
	PerrorStatus             int     `njson:"errorStatus"`
	PerrorStatusFlg          bool    `njson:"errorStatusFlg"`
	PfanAutoMode             int     `njson:"fanAutoMode"`
	PfanSpeed                int     `njson:"fanSpeed"`
	PhttpErrorCode           int     `njson:"httpErrorCode"`
	PiAuto                   int     `njson:"iAuto"`
	PinsideTemperature       float64 `njson:"insideTemperature"`
	Pnanoe                   int     `njson:"nanoe"`
	Ponline                  bool    `njson:"online"`
	Poperate                 int     `njson:"operate"`       // on/off
	PoperationMode           int     `njson:"operationMode"` // Mode (heat, dry, etc)
	PoutsideTemperature      float64 `njson:"outTemperature"`
	PpowerfulMode            bool    `njson:"powerfulMode"`
	PtemperatureSet          float64 `njson:"temperatureSet"` // Temperature
	PupdateTime              int     `njson:"updateTime"`
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
