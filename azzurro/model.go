package azzurro

import "time"

// ----------------------------------------------------------------------------
// Requests
// ----------------------------------------------------------------------------

type RealtimeDataRequest struct {
	RealtimeData RealtimeData `json:"realtimeData"`
}

type RealtimeData struct {
	Command string `json:"command"`
	Params  Params `json:"params"`
}

type Params struct {
	ThingKey       string `json:"thingKey"`
	RequiredValues string `json:"requiredValues"`
}

// ----------------------------------------------------------------------------
// Responses
// ----------------------------------------------------------------------------

type RealtimeDataResponse struct {
	RealtimeData struct {
		Params struct {
			Value []map[string]InverterMetrics `json:"value"`
		} `json:"params"`
		Success bool `json:"success"`
	} `json:"realtimeData"`
}

type InverterMetrics struct {
	// Power metrics (instantaneous values in W)
	PowerGenerating    float64 `json:"powerGenerating"`
	PowerConsuming     float64 `json:"powerConsuming"`
	PowerImporting     float64 `json:"powerImporting"`
	PowerExporting     float64 `json:"powerExporting"`
	PowerAutoconsuming float64 `json:"powerAutoconsuming"`
	PowerCharging      float64 `json:"powerCharging"`
	PowerDischarging   float64 `json:"powerDischarging"`

	// Energy metrics - current session (kWh)
	// The current session is the energy consumed or generated since the
	// last reset, probably since midnight.
	EnergyGenerating    float64 `json:"energyGenerating"`
	EnergyConsuming     float64 `json:"energyConsuming"`
	EnergyImporting     float64 `json:"energyImporting"`
	EnergyExporting     float64 `json:"energyExporting"`
	EnergyAutoconsuming float64 `json:"energyAutoconsuming"`
	EnergyCharging      float64 `json:"energyCharging"`
	EnergyDischarging   float64 `json:"energyDischarging"`

	// Energy metrics - cumulative totals (kWh)
	// The cumulative totals are the energy consumed or generated since
	// the installation of the system.
	EnergyGeneratingTotal    float64 `json:"energyGeneratingTotal"`
	EnergyConsumingTotal     float64 `json:"energyConsumingTotal"`
	EnergyImportingTotal     float64 `json:"energyImportingTotal"`
	EnergyExportingTotal     float64 `json:"energyExportingTotal"`
	EnergyAutoconsumingTotal float64 `json:"energyAutoconsumingTotal"`
	EnergyChargingTotal      float64 `json:"energyChargingTotal"`
	EnergyDischargingTotal   float64 `json:"energyDischargingTotal"`

	// Battery metrics
	BatterySoC       int `json:"batterySoC"`
	BatteryCycletime int `json:"batteryCycletime"`

	// System metadata
	LastUpdate time.Time `json:"lastUpdate"`
	ThingFind  string    `json:"thingFind"`
}
