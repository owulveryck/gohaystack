package tags

import "github.com/owulveryck/gohaystack"

var (
	// AbsorptionLabel https://project-haystack.org/tag/absorption
	AbsorptionLabel = labelDB["absorption"]
	// AcLabel https://project-haystack.org/tag/ac
	AcLabel = labelDB["ac"]
	// ActiveLabel https://project-haystack.org/tag/active
	ActiveLabel = labelDB["active"]
	// AhuLabel https://project-haystack.org/tag/ahu
	AhuLabel = labelDB["ahu"]
	// AhuRefLabel https://project-haystack.org/tag/ahuRef
	AhuRefLabel = labelDB["ahuRef"]
	// AirLabel https://project-haystack.org/tag/air
	AirLabel = labelDB["air"]
	// AirCooledLabel https://project-haystack.org/tag/airCooled
	AirCooledLabel = labelDB["airCooled"]
	// AngleLabel https://project-haystack.org/tag/angle
	AngleLabel = labelDB["angle"]
	// ApparentLabel https://project-haystack.org/tag/apparent
	ApparentLabel = labelDB["apparent"]
	// AreaLabel https://project-haystack.org/tag/area
	AreaLabel = labelDB["area"]
	// AvgLabel https://project-haystack.org/tag/avg
	AvgLabel = labelDB["avg"]
	// BarometricLabel https://project-haystack.org/tag/barometric
	BarometricLabel = labelDB["barometric"]
	// BlowdownLabel https://project-haystack.org/tag/blowdown
	BlowdownLabel = labelDB["blowdown"]
	// BoilerLabel https://project-haystack.org/tag/boiler
	BoilerLabel = labelDB["boiler"]
	// BypassLabel https://project-haystack.org/tag/bypass
	BypassLabel = labelDB["bypass"]
	// CentrifugalLabel https://project-haystack.org/tag/centrifugal
	CentrifugalLabel = labelDB["centrifugal"]
	// ChilledLabel https://project-haystack.org/tag/chilled
	ChilledLabel = labelDB["chilled"]
	// ChilledBeamZoneLabel https://project-haystack.org/tag/chilledBeamZone
	ChilledBeamZoneLabel = labelDB["chilledBeamZone"]
	// ChilledWaterCoolLabel https://project-haystack.org/tag/chilledWaterCool
	ChilledWaterCoolLabel = labelDB["chilledWaterCool"]
	// ChilledWaterPlantLabel https://project-haystack.org/tag/chilledWaterPlant
	ChilledWaterPlantLabel = labelDB["chilledWaterPlant"]
	// ChilledWaterPlantRefLabel https://project-haystack.org/tag/chilledWaterPlantRef
	ChilledWaterPlantRefLabel = labelDB["chilledWaterPlantRef"]
	// ChillerLabel https://project-haystack.org/tag/chiller
	ChillerLabel = labelDB["chiller"]
	// CircLabel https://project-haystack.org/tag/circ
	CircLabel = labelDB["circ"]
	// CircuitLabel https://project-haystack.org/tag/circuit
	CircuitLabel = labelDB["circuit"]
	// ClosedLoopLabel https://project-haystack.org/tag/closedLoop
	ClosedLoopLabel = labelDB["closedLoop"]
	// CloudageLabel https://project-haystack.org/tag/cloudage
	CloudageLabel = labelDB["cloudage"]
	// CmdLabel https://project-haystack.org/tag/cmd
	CmdLabel = labelDB["cmd"]
	// CoLabel https://project-haystack.org/tag/co
	CoLabel = labelDB["co"]
	// Co2Label https://project-haystack.org/tag/co2
	Co2Label = labelDB["co2"]
	// ColdDeckLabel https://project-haystack.org/tag/coldDeck
	ColdDeckLabel = labelDB["coldDeck"]
	// CondensateLabel https://project-haystack.org/tag/condensate
	CondensateLabel = labelDB["condensate"]
	// CondenserLabel https://project-haystack.org/tag/condenser
	CondenserLabel = labelDB["condenser"]
	// ConnectionLabel https://project-haystack.org/tag/connection
	ConnectionLabel = labelDB["connection"]
	// ConstantVolumeLabel https://project-haystack.org/tag/constantVolume
	ConstantVolumeLabel = labelDB["constantVolume"]
	// CoolLabel https://project-haystack.org/tag/cool
	CoolLabel = labelDB["cool"]
	// CoolOnlyLabel https://project-haystack.org/tag/coolOnly
	CoolOnlyLabel = labelDB["coolOnly"]
	// CoolingLabel https://project-haystack.org/tag/cooling
	CoolingLabel = labelDB["cooling"]
	// CoolingCapacityLabel https://project-haystack.org/tag/coolingCapacity
	CoolingCapacityLabel = labelDB["coolingCapacity"]
	// CoolingTowerLabel https://project-haystack.org/tag/coolingTower
	CoolingTowerLabel = labelDB["coolingTower"]
	// CurLabel https://project-haystack.org/tag/cur
	CurLabel = labelDB["cur"]
	// CurErrLabel https://project-haystack.org/tag/curErr
	CurErrLabel = labelDB["curErr"]
	// CurStatusLabel https://project-haystack.org/tag/curStatus
	CurStatusLabel = labelDB["curStatus"]
	// CurValLabel https://project-haystack.org/tag/curVal
	CurValLabel = labelDB["curVal"]
	// CurrentLabel https://project-haystack.org/tag/current
	CurrentLabel = labelDB["current"]
	// DamperLabel https://project-haystack.org/tag/damper
	DamperLabel = labelDB["damper"]
	// DcLabel https://project-haystack.org/tag/dc
	DcLabel = labelDB["dc"]
	// DeltaLabel https://project-haystack.org/tag/delta
	DeltaLabel = labelDB["delta"]
	// DeviceLabel https://project-haystack.org/tag/device
	DeviceLabel = labelDB["device"]
	// Device1RefLabel https://project-haystack.org/tag/device1Ref
	Device1RefLabel = labelDB["device1Ref"]
	// Device2RefLabel https://project-haystack.org/tag/device2Ref
	Device2RefLabel = labelDB["device2Ref"]
	// DewLabel https://project-haystack.org/tag/dew
	DewLabel = labelDB["dew"]
	// DirectZoneLabel https://project-haystack.org/tag/directZone
	DirectZoneLabel = labelDB["directZone"]
	// DirectionLabel https://project-haystack.org/tag/direction
	DirectionLabel = labelDB["direction"]
	// DisLabel https://project-haystack.org/tag/dis
	DisLabel = labelDB["dis"]
	// DischargeLabel https://project-haystack.org/tag/discharge
	DischargeLabel = labelDB["discharge"]
	// DivertingLabel https://project-haystack.org/tag/diverting
	DivertingLabel = labelDB["diverting"]
	// DomesticLabel https://project-haystack.org/tag/domestic
	DomesticLabel = labelDB["domestic"]
	// DualDuctLabel https://project-haystack.org/tag/dualDuct
	DualDuctLabel = labelDB["dualDuct"]
	// DuctAreaLabel https://project-haystack.org/tag/ductArea
	DuctAreaLabel = labelDB["ductArea"]
	// DxCoolLabel https://project-haystack.org/tag/dxCool
	DxCoolLabel = labelDB["dxCool"]
	// EffectiveLabel https://project-haystack.org/tag/effective
	EffectiveLabel = labelDB["effective"]
	// EfficiencyLabel https://project-haystack.org/tag/efficiency
	EfficiencyLabel = labelDB["efficiency"]
	// ElecLabel https://project-haystack.org/tag/elec
	ElecLabel = labelDB["elec"]
	// ElecHeatLabel https://project-haystack.org/tag/elecHeat
	ElecHeatLabel = labelDB["elecHeat"]
	// ElecMeterLoadLabel https://project-haystack.org/tag/elecMeterLoad
	ElecMeterLoadLabel = labelDB["elecMeterLoad"]
	// ElecMeterRefLabel https://project-haystack.org/tag/elecMeterRef
	ElecMeterRefLabel = labelDB["elecMeterRef"]
	// ElecPanelLabel https://project-haystack.org/tag/elecPanel
	ElecPanelLabel = labelDB["elecPanel"]
	// ElecPanelOfLabel https://project-haystack.org/tag/elecPanelOf
	ElecPanelOfLabel = labelDB["elecPanelOf"]
	// ElecReheatLabel https://project-haystack.org/tag/elecReheat
	ElecReheatLabel = labelDB["elecReheat"]
	// EnableLabel https://project-haystack.org/tag/enable
	EnableLabel = labelDB["enable"]
	// EnergyLabel https://project-haystack.org/tag/energy
	EnergyLabel = labelDB["energy"]
	// EnteringLabel https://project-haystack.org/tag/entering
	EnteringLabel = labelDB["entering"]
	// EnumLabel https://project-haystack.org/tag/enum
	EnumLabel = labelDB["enum"]
	// EquipLabel https://project-haystack.org/tag/equip
	EquipLabel = labelDB["equip"]
	// EquipRefLabel https://project-haystack.org/tag/equipRef
	EquipRefLabel = labelDB["equipRef"]
	// EvaporatorLabel https://project-haystack.org/tag/evaporator
	EvaporatorLabel = labelDB["evaporator"]
	// ExhaustLabel https://project-haystack.org/tag/exhaust
	ExhaustLabel = labelDB["exhaust"]
	// ExportLabel https://project-haystack.org/tag/export
	ExportLabel = labelDB["export"]
	// FaceBypassLabel https://project-haystack.org/tag/faceBypass
	FaceBypassLabel = labelDB["faceBypass"]
	// FanLabel https://project-haystack.org/tag/fan
	FanLabel = labelDB["fan"]
	// FanPoweredLabel https://project-haystack.org/tag/fanPowered
	FanPoweredLabel = labelDB["fanPowered"]
	// FcuLabel https://project-haystack.org/tag/fcu
	FcuLabel = labelDB["fcu"]
	// FilterLabel https://project-haystack.org/tag/filter
	FilterLabel = labelDB["filter"]
	// FlowLabel https://project-haystack.org/tag/flow
	FlowLabel = labelDB["flow"]
	// FlueLabel https://project-haystack.org/tag/flue
	FlueLabel = labelDB["flue"]
	// FreezeStatLabel https://project-haystack.org/tag/freezeStat
	FreezeStatLabel = labelDB["freezeStat"]
	// FreqLabel https://project-haystack.org/tag/freq
	FreqLabel = labelDB["freq"]
	// GasLabel https://project-haystack.org/tag/gas
	GasLabel = labelDB["gas"]
	// GasHeatLabel https://project-haystack.org/tag/gasHeat
	GasHeatLabel = labelDB["gasHeat"]
	// GasMeterLoadLabel https://project-haystack.org/tag/gasMeterLoad
	GasMeterLoadLabel = labelDB["gasMeterLoad"]
	// GeoAddrLabel https://project-haystack.org/tag/geoAddr
	GeoAddrLabel = labelDB["geoAddr"]
	// GeoCityLabel https://project-haystack.org/tag/geoCity
	GeoCityLabel = labelDB["geoCity"]
	// GeoCoordLabel https://project-haystack.org/tag/geoCoord
	GeoCoordLabel = labelDB["geoCoord"]
	// GeoCountryLabel https://project-haystack.org/tag/geoCountry
	GeoCountryLabel = labelDB["geoCountry"]
	// GeoCountyLabel https://project-haystack.org/tag/geoCounty
	GeoCountyLabel = labelDB["geoCounty"]
	// GeoPostalCodeLabel https://project-haystack.org/tag/geoPostalCode
	GeoPostalCodeLabel = labelDB["geoPostalCode"]
	// GeoStateLabel https://project-haystack.org/tag/geoState
	GeoStateLabel = labelDB["geoState"]
	// GeoStreetLabel https://project-haystack.org/tag/geoStreet
	GeoStreetLabel = labelDB["geoStreet"]
	// HeaderLabel https://project-haystack.org/tag/header
	HeaderLabel = labelDB["header"]
	// HeatLabel https://project-haystack.org/tag/heat
	HeatLabel = labelDB["heat"]
	// HeatExchangerLabel https://project-haystack.org/tag/heatExchanger
	HeatExchangerLabel = labelDB["heatExchanger"]
	// HeatPumpLabel https://project-haystack.org/tag/heatPump
	HeatPumpLabel = labelDB["heatPump"]
	// HeatWheelLabel https://project-haystack.org/tag/heatWheel
	HeatWheelLabel = labelDB["heatWheel"]
	// HeatingLabel https://project-haystack.org/tag/heating
	HeatingLabel = labelDB["heating"]
	// HisLabel https://project-haystack.org/tag/his
	HisLabel = labelDB["his"]
	// HisErrLabel https://project-haystack.org/tag/hisErr
	HisErrLabel = labelDB["hisErr"]
	// HisInterpolateLabel https://project-haystack.org/tag/hisInterpolate
	HisInterpolateLabel = labelDB["hisInterpolate"]
	// HisStatusLabel https://project-haystack.org/tag/hisStatus
	HisStatusLabel = labelDB["hisStatus"]
	// HisTotalizedLabel https://project-haystack.org/tag/hisTotalized
	HisTotalizedLabel = labelDB["hisTotalized"]
	// HotLabel https://project-haystack.org/tag/hot
	HotLabel = labelDB["hot"]
	// HotDeckLabel https://project-haystack.org/tag/hotDeck
	HotDeckLabel = labelDB["hotDeck"]
	// HotWaterHeatLabel https://project-haystack.org/tag/hotWaterHeat
	HotWaterHeatLabel = labelDB["hotWaterHeat"]
	// HotWaterPlantLabel https://project-haystack.org/tag/hotWaterPlant
	HotWaterPlantLabel = labelDB["hotWaterPlant"]
	// HotWaterPlantRefLabel https://project-haystack.org/tag/hotWaterPlantRef
	HotWaterPlantRefLabel = labelDB["hotWaterPlantRef"]
	// HotWaterReheatLabel https://project-haystack.org/tag/hotWaterReheat
	HotWaterReheatLabel = labelDB["hotWaterReheat"]
	// HumidifierLabel https://project-haystack.org/tag/humidifier
	HumidifierLabel = labelDB["humidifier"]
	// HumidityLabel https://project-haystack.org/tag/humidity
	HumidityLabel = labelDB["humidity"]
	// HvacLabel https://project-haystack.org/tag/hvac
	HvacLabel = labelDB["hvac"]
	// IDLabel https://project-haystack.org/tag/iD
	IDLabel = labelDB["id"]
	// ImbalanceLabel https://project-haystack.org/tag/imbalance
	ImbalanceLabel = labelDB["imbalance"]
	// ImportLabel https://project-haystack.org/tag/import
	ImportLabel = labelDB["import"]
	// IrradianceLabel https://project-haystack.org/tag/irradiance
	IrradianceLabel = labelDB["irradiance"]
	// IsolationLabel https://project-haystack.org/tag/isolation
	IsolationLabel = labelDB["isolation"]
	// KindLabel https://project-haystack.org/tag/kind
	KindLabel = labelDB["kind"]
	// LeavingLabel https://project-haystack.org/tag/leaving
	LeavingLabel = labelDB["leaving"]
	// LevelLabel https://project-haystack.org/tag/level
	LevelLabel = labelDB["level"]
	// LightLevelLabel https://project-haystack.org/tag/lightLevel
	LightLevelLabel = labelDB["lightLevel"]
	// LightingLabel https://project-haystack.org/tag/lighting
	LightingLabel = labelDB["lighting"]
	// LightsLabel https://project-haystack.org/tag/lights
	LightsLabel = labelDB["lights"]
	// LightsGroupLabel https://project-haystack.org/tag/lightsGroup
	LightsGroupLabel = labelDB["lightsGroup"]
	// LoadLabel https://project-haystack.org/tag/load
	LoadLabel = labelDB["load"]
	// MagLabel https://project-haystack.org/tag/mag
	MagLabel = labelDB["mag"]
	// MakeupLabel https://project-haystack.org/tag/makeup
	MakeupLabel = labelDB["makeup"]
	// MauLabel https://project-haystack.org/tag/mau
	MauLabel = labelDB["mau"]
	// MaxLabel https://project-haystack.org/tag/max
	MaxLabel = labelDB["max"]
	// MaxValLabel https://project-haystack.org/tag/maxVal
	MaxValLabel = labelDB["maxVal"]
	// MeterLabel https://project-haystack.org/tag/meter
	MeterLabel = labelDB["meter"]
	// MinLabel https://project-haystack.org/tag/min
	MinLabel = labelDB["min"]
	// MinValLabel https://project-haystack.org/tag/minVal
	MinValLabel = labelDB["minVal"]
	// MixedLabel https://project-haystack.org/tag/mixed
	MixedLabel = labelDB["mixed"]
	// MixingLabel https://project-haystack.org/tag/mixing
	MixingLabel = labelDB["mixing"]
	// MultiZoneLabel https://project-haystack.org/tag/multiZone
	MultiZoneLabel = labelDB["multiZone"]
	// NetLabel https://project-haystack.org/tag/net
	NetLabel = labelDB["net"]
	// NetworkLabel https://project-haystack.org/tag/network
	NetworkLabel = labelDB["network"]
	// NetworkRefLabel https://project-haystack.org/tag/networkRef
	NetworkRefLabel = labelDB["networkRef"]
	// NeutralDeckLabel https://project-haystack.org/tag/neutralDeck
	NeutralDeckLabel = labelDB["neutralDeck"]
	// OccLabel https://project-haystack.org/tag/occ
	OccLabel = labelDB["occ"]
	// OccupancyIndicatorLabel https://project-haystack.org/tag/occupancyIndicator
	OccupancyIndicatorLabel = labelDB["occupancyIndicator"]
	// OccupiedLabel https://project-haystack.org/tag/occupied
	OccupiedLabel = labelDB["occupied"]
	// OilLabel https://project-haystack.org/tag/oil
	OilLabel = labelDB["oil"]
	// OpenLoopLabel https://project-haystack.org/tag/openLoop
	OpenLoopLabel = labelDB["openLoop"]
	// OutsideLabel https://project-haystack.org/tag/outside
	OutsideLabel = labelDB["outside"]
	// ParallelLabel https://project-haystack.org/tag/parallel
	ParallelLabel = labelDB["parallel"]
	// PerimeterHeatLabel https://project-haystack.org/tag/perimeterHeat
	PerimeterHeatLabel = labelDB["perimeterHeat"]
	// PfLabel https://project-haystack.org/tag/pf
	PfLabel = labelDB["pf"]
	// PhaseLabel https://project-haystack.org/tag/phase
	PhaseLabel = labelDB["phase"]
	// PointLabel https://project-haystack.org/tag/point
	PointLabel = labelDB["point"]
	// PowerLabel https://project-haystack.org/tag/power
	PowerLabel = labelDB["power"]
	// PrecipitationLabel https://project-haystack.org/tag/precipitation
	PrecipitationLabel = labelDB["precipitation"]
	// PressureLabel https://project-haystack.org/tag/pressure
	PressureLabel = labelDB["pressure"]
	// PressureDependentLabel https://project-haystack.org/tag/pressureDependent
	PressureDependentLabel = labelDB["pressureDependent"]
	// PressureIndependentLabel https://project-haystack.org/tag/pressureIndependent
	PressureIndependentLabel = labelDB["pressureIndependent"]
	// PrimaryFunctionLabel https://project-haystack.org/tag/primaryFunction
	PrimaryFunctionLabel = labelDB["primaryFunction"]
	// PrimaryLoopLabel https://project-haystack.org/tag/primaryLoop
	PrimaryLoopLabel = labelDB["primaryLoop"]
	// ProtocolLabel https://project-haystack.org/tag/protocol
	ProtocolLabel = labelDB["protocol"]
	// PumpLabel https://project-haystack.org/tag/pump
	PumpLabel = labelDB["pump"]
	// ReactiveLabel https://project-haystack.org/tag/reactive
	ReactiveLabel = labelDB["reactive"]
	// ReciprocalLabel https://project-haystack.org/tag/reciprocal
	ReciprocalLabel = labelDB["reciprocal"]
	// RefrigLabel https://project-haystack.org/tag/refrig
	RefrigLabel = labelDB["refrig"]
	// ReheatLabel https://project-haystack.org/tag/reheat
	ReheatLabel = labelDB["reheat"]
	// ReheatingLabel https://project-haystack.org/tag/reheating
	ReheatingLabel = labelDB["reheating"]
	// ReturnLabel https://project-haystack.org/tag/return
	ReturnLabel = labelDB["return"]
	// RooftopLabel https://project-haystack.org/tag/rooftop
	RooftopLabel = labelDB["rooftop"]
	// RunLabel https://project-haystack.org/tag/run
	RunLabel = labelDB["run"]
	// ScrewLabel https://project-haystack.org/tag/screw
	ScrewLabel = labelDB["screw"]
	// SecondaryLoopLabel https://project-haystack.org/tag/secondaryLoop
	SecondaryLoopLabel = labelDB["secondaryLoop"]
	// SensorLabel https://project-haystack.org/tag/sensor
	SensorLabel = labelDB["sensor"]
	// SeriesLabel https://project-haystack.org/tag/series
	SeriesLabel = labelDB["series"]
	// SingleDuctLabel https://project-haystack.org/tag/singleDuct
	SingleDuctLabel = labelDB["singleDuct"]
	// SiteLabel https://project-haystack.org/tag/site
	SiteLabel = labelDB["site"]
	// SiteMeterLabel https://project-haystack.org/tag/siteMeter
	SiteMeterLabel = labelDB["siteMeter"]
	// SitePanelLabel https://project-haystack.org/tag/sitePanel
	SitePanelLabel = labelDB["sitePanel"]
	// SiteRefLabel https://project-haystack.org/tag/siteRef
	SiteRefLabel = labelDB["siteRef"]
	// SolarLabel https://project-haystack.org/tag/solar
	SolarLabel = labelDB["solar"]
	// SpLabel https://project-haystack.org/tag/sp
	SpLabel = labelDB["sp"]
	// SpeedLabel https://project-haystack.org/tag/speed
	SpeedLabel = labelDB["speed"]
	// StageLabel https://project-haystack.org/tag/stage
	StageLabel = labelDB["stage"]
	// StandbyLabel https://project-haystack.org/tag/standby
	StandbyLabel = labelDB["standby"]
	// SteamLabel https://project-haystack.org/tag/steam
	SteamLabel = labelDB["steam"]
	// SteamHeatLabel https://project-haystack.org/tag/steamHeat
	SteamHeatLabel = labelDB["steamHeat"]
	// SteamMeterLoadLabel https://project-haystack.org/tag/steamMeterLoad
	SteamMeterLoadLabel = labelDB["steamMeterLoad"]
	// SteamPlantLabel https://project-haystack.org/tag/steamPlant
	SteamPlantLabel = labelDB["steamPlant"]
	// SteamPlantRefLabel https://project-haystack.org/tag/steamPlantRef
	SteamPlantRefLabel = labelDB["steamPlantRef"]
	// SubPanelOfLabel https://project-haystack.org/tag/subPanelOf
	SubPanelOfLabel = labelDB["subPanelOf"]
	// SubmeterOfLabel https://project-haystack.org/tag/submeterOf
	SubmeterOfLabel = labelDB["submeterOf"]
	// SunriseLabel https://project-haystack.org/tag/sunrise
	SunriseLabel = labelDB["sunrise"]
	// TankLabel https://project-haystack.org/tag/tank
	TankLabel = labelDB["tank"]
	// TempLabel https://project-haystack.org/tag/temp
	TempLabel = labelDB["temp"]
	// ThdLabel https://project-haystack.org/tag/thd
	ThdLabel = labelDB["thd"]
	// TotalLabel https://project-haystack.org/tag/total
	TotalLabel = labelDB["total"]
	// TripleDuctLabel https://project-haystack.org/tag/tripleDuct
	TripleDuctLabel = labelDB["tripleDuct"]
	// TzLabel https://project-haystack.org/tag/tz
	TzLabel = labelDB["tz"]
	// UnitLabel https://project-haystack.org/tag/unit
	UnitLabel = labelDB["unit"]
	// UnoccLabel https://project-haystack.org/tag/unocc
	UnoccLabel = labelDB["unocc"]
	// UvLabel https://project-haystack.org/tag/uv
	UvLabel = labelDB["uv"]
	// ValveLabel https://project-haystack.org/tag/valve
	ValveLabel = labelDB["valve"]
	// VariableVolumeLabel https://project-haystack.org/tag/variableVolume
	VariableVolumeLabel = labelDB["variableVolume"]
	// VavLabel https://project-haystack.org/tag/vav
	VavLabel = labelDB["vav"]
	// VavModeLabel https://project-haystack.org/tag/vavMode
	VavModeLabel = labelDB["vavMode"]
	// VavZoneLabel https://project-haystack.org/tag/vavZone
	VavZoneLabel = labelDB["vavZone"]
	// VfdLabel https://project-haystack.org/tag/vfd
	VfdLabel = labelDB["vfd"]
	// VisibilityLabel https://project-haystack.org/tag/visibility
	VisibilityLabel = labelDB["visibility"]
	// VoltLabel https://project-haystack.org/tag/volt
	VoltLabel = labelDB["volt"]
	// VolumeLabel https://project-haystack.org/tag/volume
	VolumeLabel = labelDB["volume"]
	// WaterLabel https://project-haystack.org/tag/water
	WaterLabel = labelDB["water"]
	// WaterCooledLabel https://project-haystack.org/tag/waterCooled
	WaterCooledLabel = labelDB["waterCooled"]
	// WaterMeterLoadLabel https://project-haystack.org/tag/waterMeterLoad
	WaterMeterLoadLabel = labelDB["waterMeterLoad"]
	// WeatherLabel https://project-haystack.org/tag/weather
	WeatherLabel = labelDB["weather"]
	// WeatherCondLabel https://project-haystack.org/tag/weatherCond
	WeatherCondLabel = labelDB["weatherCond"]
	// WeatherPointLabel https://project-haystack.org/tag/weatherPoint
	WeatherPointLabel = labelDB["weatherPoint"]
	// WeatherRefLabel https://project-haystack.org/tag/weatherRef
	WeatherRefLabel = labelDB["weatherRef"]
	// WetBulbLabel https://project-haystack.org/tag/wetBulb
	WetBulbLabel = labelDB["wetBulb"]
	// WindLabel https://project-haystack.org/tag/wind
	WindLabel = labelDB["wind"]
	// WritableLabel https://project-haystack.org/tag/writable
	WritableLabel = labelDB["writable"]
	// WriteErrLabel https://project-haystack.org/tag/writeErr
	WriteErrLabel = labelDB["writeErr"]
	// WriteLevelLabel https://project-haystack.org/tag/writeLevel
	WriteLevelLabel = labelDB["writeLevel"]
	// WriteStatusLabel https://project-haystack.org/tag/writeStatus
	WriteStatusLabel = labelDB["writeStatus"]
	// WriteValLabel https://project-haystack.org/tag/writeVal
	WriteValLabel = labelDB["writeVal"]
	// YearBuiltLabel https://project-haystack.org/tag/yearBuilt
	YearBuiltLabel = labelDB["yearBuilt"]
	// ZoneLabel https://project-haystack.org/tag/zone
	ZoneLabel = labelDB["zone"]
)

var (
	// Absorption Marker Tag
	Absorption = Marker(AbsorptionLabel)
	// Ac Marker Tag
	Ac = Marker(AcLabel)
	// Active Marker Tag
	Active = Marker(ActiveLabel)
	// Ahu Marker Tag
	Ahu = Marker(AhuLabel)
	// Air Marker Tag
	Air = Marker(AirLabel)
	// AirCooled Marker Tag
	AirCooled = Marker(AirCooledLabel)
	// Angle Marker Tag
	Angle = Marker(AngleLabel)
	// Apparent Marker Tag
	Apparent = Marker(ApparentLabel)
	// Avg Marker Tag
	Avg = Marker(AvgLabel)
	// Barometric Marker Tag
	Barometric = Marker(BarometricLabel)
	// Blowdown Marker Tag
	Blowdown = Marker(BlowdownLabel)
	// Boiler Marker Tag
	Boiler = Marker(BoilerLabel)
	// Bypass Marker Tag
	Bypass = Marker(BypassLabel)
	// Centrifugal Marker Tag
	Centrifugal = Marker(CentrifugalLabel)
	// Chilled Marker Tag
	Chilled = Marker(ChilledLabel)
	// ChilledBeamZone Marker Tag
	ChilledBeamZone = Marker(ChilledBeamZoneLabel)
	// ChilledWaterCool Marker Tag
	ChilledWaterCool = Marker(ChilledWaterCoolLabel)
	// ChilledWaterPlant Marker Tag
	ChilledWaterPlant = Marker(ChilledWaterPlantLabel)
	// Chiller Marker Tag
	Chiller = Marker(ChillerLabel)
	// Circ Marker Tag
	Circ = Marker(CircLabel)
	// Circuit Marker Tag
	Circuit = Marker(CircuitLabel)
	// ClosedLoop Marker Tag
	ClosedLoop = Marker(ClosedLoopLabel)
	// Cloudage Marker Tag
	Cloudage = Marker(CloudageLabel)
	// Cmd Marker Tag
	Cmd = Marker(CmdLabel)
	// Co Marker Tag
	Co = Marker(CoLabel)
	// Co2 Marker Tag
	Co2 = Marker(Co2Label)
	// ColdDeck Marker Tag
	ColdDeck = Marker(ColdDeckLabel)
	// Condensate Marker Tag
	Condensate = Marker(CondensateLabel)
	// Condenser Marker Tag
	Condenser = Marker(CondenserLabel)
	// Connection Marker Tag
	Connection = Marker(ConnectionLabel)
	// ConstantVolume Marker Tag
	ConstantVolume = Marker(ConstantVolumeLabel)
	// Cool Marker Tag
	Cool = Marker(CoolLabel)
	// CoolOnly Marker Tag
	CoolOnly = Marker(CoolOnlyLabel)
	// Cooling Marker Tag
	Cooling = Marker(CoolingLabel)
	// CoolingTower Marker Tag
	CoolingTower = Marker(CoolingTowerLabel)
	// Cur Marker Tag
	Cur = Marker(CurLabel)
	// Current Marker Tag
	Current = Marker(CurrentLabel)
	// Damper Marker Tag
	Damper = Marker(DamperLabel)
	// Dc Marker Tag
	Dc = Marker(DcLabel)
	// Delta Marker Tag
	Delta = Marker(DeltaLabel)
	// Device Marker Tag
	Device = Marker(DeviceLabel)
	// Dew Marker Tag
	Dew = Marker(DewLabel)
	// DirectZone Marker Tag
	DirectZone = Marker(DirectZoneLabel)
	// Direction Marker Tag
	Direction = Marker(DirectionLabel)
	// Discharge Marker Tag
	Discharge = Marker(DischargeLabel)
	// Diverting Marker Tag
	Diverting = Marker(DivertingLabel)
	// Domestic Marker Tag
	Domestic = Marker(DomesticLabel)
	// DualDuct Marker Tag
	DualDuct = Marker(DualDuctLabel)
	// DuctArea Marker Tag
	DuctArea = Marker(DuctAreaLabel)
	// DxCool Marker Tag
	DxCool = Marker(DxCoolLabel)
	// Effective Marker Tag
	Effective = Marker(EffectiveLabel)
	// Efficiency Marker Tag
	Efficiency = Marker(EfficiencyLabel)
	// Elec Marker Tag
	Elec = Marker(ElecLabel)
	// ElecHeat Marker Tag
	ElecHeat = Marker(ElecHeatLabel)
	// ElecPanel Marker Tag
	ElecPanel = Marker(ElecPanelLabel)
	// ElecReheat Marker Tag
	ElecReheat = Marker(ElecReheatLabel)
	// Enable Marker Tag
	Enable = Marker(EnableLabel)
	// Energy Marker Tag
	Energy = Marker(EnergyLabel)
	// Entering Marker Tag
	Entering = Marker(EnteringLabel)
	// Equip Marker Tag
	Equip = Marker(EquipLabel)
	// Evaporator Marker Tag
	Evaporator = Marker(EvaporatorLabel)
	// Exhaust Marker Tag
	Exhaust = Marker(ExhaustLabel)
	// Export Marker Tag
	Export = Marker(ExportLabel)
	// FaceBypass Marker Tag
	FaceBypass = Marker(FaceBypassLabel)
	// Fan Marker Tag
	Fan = Marker(FanLabel)
	// FanPowered Marker Tag
	FanPowered = Marker(FanPoweredLabel)
	// Fcu Marker Tag
	Fcu = Marker(FcuLabel)
	// Filter Marker Tag
	Filter = Marker(FilterLabel)
	// Flow Marker Tag
	Flow = Marker(FlowLabel)
	// Flue Marker Tag
	Flue = Marker(FlueLabel)
	// FreezeStat Marker Tag
	FreezeStat = Marker(FreezeStatLabel)
	// Freq Marker Tag
	Freq = Marker(FreqLabel)
	// Gas Marker Tag
	Gas = Marker(GasLabel)
	// GasHeat Marker Tag
	GasHeat = Marker(GasHeatLabel)
	// Header Marker Tag
	Header = Marker(HeaderLabel)
	// Heat Marker Tag
	Heat = Marker(HeatLabel)
	// HeatExchanger Marker Tag
	HeatExchanger = Marker(HeatExchangerLabel)
	// HeatPump Marker Tag
	HeatPump = Marker(HeatPumpLabel)
	// HeatWheel Marker Tag
	HeatWheel = Marker(HeatWheelLabel)
	// Heating Marker Tag
	Heating = Marker(HeatingLabel)
	// His Marker Tag
	His = Marker(HisLabel)
	// HisTotalized Marker Tag
	HisTotalized = Marker(HisTotalizedLabel)
	// Hot Marker Tag
	Hot = Marker(HotLabel)
	// HotDeck Marker Tag
	HotDeck = Marker(HotDeckLabel)
	// HotWaterHeat Marker Tag
	HotWaterHeat = Marker(HotWaterHeatLabel)
	// HotWaterPlant Marker Tag
	HotWaterPlant = Marker(HotWaterPlantLabel)
	// HotWaterReheat Marker Tag
	HotWaterReheat = Marker(HotWaterReheatLabel)
	// Humidifier Marker Tag
	Humidifier = Marker(HumidifierLabel)
	// Humidity Marker Tag
	Humidity = Marker(HumidityLabel)
	// Hvac Marker Tag
	Hvac = Marker(HvacLabel)
	// Imbalance Marker Tag
	Imbalance = Marker(ImbalanceLabel)
	// Import Marker Tag
	Import = Marker(ImportLabel)
	// Irradiance Marker Tag
	Irradiance = Marker(IrradianceLabel)
	// Isolation Marker Tag
	Isolation = Marker(IsolationLabel)
	// Leaving Marker Tag
	Leaving = Marker(LeavingLabel)
	// Level Marker Tag
	Level = Marker(LevelLabel)
	// LightLevel Marker Tag
	LightLevel = Marker(LightLevelLabel)
	// Lighting Marker Tag
	Lighting = Marker(LightingLabel)
	// Lights Marker Tag
	Lights = Marker(LightsLabel)
	// LightsGroup Marker Tag
	LightsGroup = Marker(LightsGroupLabel)
	// Load Marker Tag
	Load = Marker(LoadLabel)
	// Mag Marker Tag
	Mag = Marker(MagLabel)
	// Makeup Marker Tag
	Makeup = Marker(MakeupLabel)
	// Mau Marker Tag
	Mau = Marker(MauLabel)
	// Max Marker Tag
	Max = Marker(MaxLabel)
	// Meter Marker Tag
	Meter = Marker(MeterLabel)
	// Min Marker Tag
	Min = Marker(MinLabel)
	// Mixed Marker Tag
	Mixed = Marker(MixedLabel)
	// Mixing Marker Tag
	Mixing = Marker(MixingLabel)
	// MultiZone Marker Tag
	MultiZone = Marker(MultiZoneLabel)
	// Net Marker Tag
	Net = Marker(NetLabel)
	// Network Marker Tag
	Network = Marker(NetworkLabel)
	// NeutralDeck Marker Tag
	NeutralDeck = Marker(NeutralDeckLabel)
	// Occ Marker Tag
	Occ = Marker(OccLabel)
	// OccupancyIndicator Marker Tag
	OccupancyIndicator = Marker(OccupancyIndicatorLabel)
	// Occupied Marker Tag
	Occupied = Marker(OccupiedLabel)
	// Oil Marker Tag
	Oil = Marker(OilLabel)
	// OpenLoop Marker Tag
	OpenLoop = Marker(OpenLoopLabel)
	// Outside Marker Tag
	Outside = Marker(OutsideLabel)
	// Parallel Marker Tag
	Parallel = Marker(ParallelLabel)
	// PerimeterHeat Marker Tag
	PerimeterHeat = Marker(PerimeterHeatLabel)
	// Pf Marker Tag
	Pf = Marker(PfLabel)
	// Point Marker Tag
	Point = Marker(PointLabel)
	// Power Marker Tag
	Power = Marker(PowerLabel)
	// Precipitation Marker Tag
	Precipitation = Marker(PrecipitationLabel)
	// Pressure Marker Tag
	Pressure = Marker(PressureLabel)
	// PressureDependent Marker Tag
	PressureDependent = Marker(PressureDependentLabel)
	// PressureIndependent Marker Tag
	PressureIndependent = Marker(PressureIndependentLabel)
	// PrimaryLoop Marker Tag
	PrimaryLoop = Marker(PrimaryLoopLabel)
	// Pump Marker Tag
	Pump = Marker(PumpLabel)
	// Reactive Marker Tag
	Reactive = Marker(ReactiveLabel)
	// Reciprocal Marker Tag
	Reciprocal = Marker(ReciprocalLabel)
	// Refrig Marker Tag
	Refrig = Marker(RefrigLabel)
	// Reheat Marker Tag
	Reheat = Marker(ReheatLabel)
	// Reheating Marker Tag
	Reheating = Marker(ReheatingLabel)
	// Return Marker Tag
	Return = Marker(ReturnLabel)
	// Rooftop Marker Tag
	Rooftop = Marker(RooftopLabel)
	// Run Marker Tag
	Run = Marker(RunLabel)
	// Screw Marker Tag
	Screw = Marker(ScrewLabel)
	// SecondaryLoop Marker Tag
	SecondaryLoop = Marker(SecondaryLoopLabel)
	// Sensor Marker Tag
	Sensor = Marker(SensorLabel)
	// Series Marker Tag
	Series = Marker(SeriesLabel)
	// SingleDuct Marker Tag
	SingleDuct = Marker(SingleDuctLabel)
	// Site Marker Tag
	Site = Marker(SiteLabel)
	// SiteMeter Marker Tag
	SiteMeter = Marker(SiteMeterLabel)
	// SitePanel Marker Tag
	SitePanel = Marker(SitePanelLabel)
	// Solar Marker Tag
	Solar = Marker(SolarLabel)
	// Sp Marker Tag
	Sp = Marker(SpLabel)
	// Speed Marker Tag
	Speed = Marker(SpeedLabel)
	// Standby Marker Tag
	Standby = Marker(StandbyLabel)
	// Steam Marker Tag
	Steam = Marker(SteamLabel)
	// SteamHeat Marker Tag
	SteamHeat = Marker(SteamHeatLabel)
	// SteamPlant Marker Tag
	SteamPlant = Marker(SteamPlantLabel)
	// Sunrise Marker Tag
	Sunrise = Marker(SunriseLabel)
	// Tank Marker Tag
	Tank = Marker(TankLabel)
	// Temp Marker Tag
	Temp = Marker(TempLabel)
	// Thd Marker Tag
	Thd = Marker(ThdLabel)
	// Total Marker Tag
	Total = Marker(TotalLabel)
	// TripleDuct Marker Tag
	TripleDuct = Marker(TripleDuctLabel)
	// Unocc Marker Tag
	Unocc = Marker(UnoccLabel)
	// Uv Marker Tag
	Uv = Marker(UvLabel)
	// Valve Marker Tag
	Valve = Marker(ValveLabel)
	// VariableVolume Marker Tag
	VariableVolume = Marker(VariableVolumeLabel)
	// Vav Marker Tag
	Vav = Marker(VavLabel)
	// VavMode Marker Tag
	VavMode = Marker(VavModeLabel)
	// VavZone Marker Tag
	VavZone = Marker(VavZoneLabel)
	// Vfd Marker Tag
	Vfd = Marker(VfdLabel)
	// Visibility Marker Tag
	Visibility = Marker(VisibilityLabel)
	// Volt Marker Tag
	Volt = Marker(VoltLabel)
	// Volume Marker Tag
	Volume = Marker(VolumeLabel)
	// Water Marker Tag
	Water = Marker(WaterLabel)
	// WaterCooled Marker Tag
	WaterCooled = Marker(WaterCooledLabel)
	// Weather Marker Tag
	Weather = Marker(WeatherLabel)
	// WeatherCond Marker Tag
	WeatherCond = Marker(WeatherCondLabel)
	// WeatherPoint Marker Tag
	WeatherPoint = Marker(WeatherPointLabel)
	// WetBulb Marker Tag
	WetBulb = Marker(WetBulbLabel)
	// Wind Marker Tag
	Wind = Marker(WindLabel)
	// Writable Marker Tag
	Writable = Marker(WritableLabel)
	// Zone Marker Tag
	Zone = Marker(ZoneLabel)
)

// Marker is a helper function to generate a https://project-haystack.org/tag/ Marker is a helper function to generate a
func Marker(l *gohaystack.Label) func() (*gohaystack.Label, *gohaystack.Value) {
	return func() (*gohaystack.Label, *gohaystack.Value) {
		return l, gohaystack.MarkerValue
	}
}
