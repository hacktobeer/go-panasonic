package types

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

//Device is Panasonic device
type Device struct {
	AirSwingLR               bool    `njson:"airSwingLR"`
	AutoMode                 bool    `njson:"autoMode"`
	AutoTempMax              int     `njson:"autoTempMax"`
	AutoTempMin              int     `njson:"autoTempMin"`
	CoolMode                 bool    `njson:"coolMode"`
	CoolTempMax              int     `njson:"coolTeampMax"`
	CoolTempMin              int     `njson:"coolTempMin"`
	DeviceGUID               string  `njson:"deviceGuid"`
	DeviceHashGUID           string  `njson:"deviceHashGuid"`
	DeviceModuleNumber       string  `njson:"deviceModuleNumber"`
	DeviceName               string  `njson:"deviceName"`
	DeviceType               string  `njson:"deviceType"`
	DryMode                  bool    `njson:"dryMode"`
	DryTempMax               int     `njson:"dryTempMax"`
	DryTempMin               int     `njson:"dryTempMin"`
	EcoFunction              bool    `njson:"ecoFunction"`
	EcoNavi                  bool    `njson:"ecoNavi"`
	FanDirectionMode         int     `njson:"fanDirectionMode"`
	FanMode                  bool    `njson:"fanMode"`
	FanSpeedMode             int     `njson:"fanSpeedMode"`
	HeatMode                 bool    `njson:"heatMode"`
	HeatTempMax              int     `njson:"heatTempMax"`
	HeatTempMin              int     `njson:"heatTeampMin"`
	IautoX                   bool    `njson:"iAutoX"`
	ModeAvlAutoMode          bool    `njson:"modeAvlList.autoMode"`
	ModeAvlFanMode           bool    `njson:"modeAvlList.fanMode"`
	Nanoe                    bool    `njson:"nanoe"`
	PactualNanoe             int     `njson:"parameters.actualNanoe"`
	PairDirection            int     `njson:"parameters.airDirection"`
	PairQuality              int     `njson:"parameters.airQuality"`
	PairSwingLR              int     `njson:"parameters.airSwingLR"`
	PairSwingUD              int     `njson:"parameters.airSwingUD"`
	Pdefrosting              int     `njson:"parameters.defrosting"`
	PdevGUID                 string  `njson:"parameters.devGuid"`
	PdevRacCommunicateStatus int     `njson:"parameters.devRacCommunicateStatus"`
	PecoFunctionData         int     `njson:"parameters.ecoFunctionData"`
	PecoMode                 int     `njson:"parameters.ecoMode"`
	PecoNavi                 int     `njson:"parameters.ecoNavi"`
	Permission               int     `njson:"permission"`
	PerrorCode               int     `njson:"parameters.errorCode"`
	PerrorCodeStr            string  `njson:"parameters.errorCodeStr"`
	PerrorStatus             int     `njson:"parameters.errorStatus"`
	PerrorStatusFlg          bool    `njson:"parameters.errorStatusFlg"`
	PfanAutoMode             int     `njson:"parameters.fanAutoMode"`
	PfanSpeed                int     `njson:"parameters.fanSpeed"`
	PhttpErrorCode           int     `njson:"parameters.httpErrorCode"`
	PiAuto                   int     `njson:"parameters.iAuto"`
	PinsideTemperature       float64 `njson:"parameters.insideTemperature"`
	Pnanoe                   int     `njson:"parameters.nanoe"`
	Ponline                  bool    `njson:"parameters.online"`
	Poperate                 int     `njson:"parameters.operate"`
	PoperationMode           int     `njson:"parameters.operationMode"`
	PoutsideTemperature      float64 `njson:"parameters.outTemperature"`
	PpowerfulMode            bool    `njson:"powerfulMode"`
	PtemperatureSet          float64 `njson:"parameters.temperatureSet"`
	PupdateTime              int     `njson:"parameters.updateTime"`
	QuietMode                bool    `njson:"quietMode"`
	SummerHouse              int     `njson:"summerHouse"`
	TemperatureUnit          bool    `njson:"temperatureUnit"`
	TimeStamp                int     `njson:"timestamp"`
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

// HistoryDataMode maps out the time intervals to fetch history data
var HistoryDataMode = map[string]int{
	"day":   0,
	"week":  1,
	"month": 2,
	"year":  3,
}

// Exported constants
const (
	ModeAuto = 0
	ModeDry  = 1
	ModeCool = 2
	ModeHeat = 3
	ModeFan  = 4
)

// Modes definesthe different AC modes the device can be in
var Modes = map[string]int{
	"auto": 0,
	"dry":  1,
	"cool": 2,
	"heat": 3,
	"fan":  3,
}

// Command is basic command control structure
type Command struct {
	DeviceGUID string         `json:"deviceGuid"`
	Parameters CommandDetails `json:"parameters"`
}

// CommandDetails stores details of a command
type CommandDetails struct {
	OperationMode  *int     `json:"operationMode,omitempty"`  // AC Mode
	TemperatureSet *float64 `json:"temperatureSet,omitempty"` // Set temperature
	Operate        *int     `json:"operate,omitempty"`        // Turn device on/off (0/1)
}
