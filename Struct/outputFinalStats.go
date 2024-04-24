package Struct

// FinalStats - Struct for output yaml
type FinalStats struct {
	Gas       StationStats `yaml:"Gas"`
	Diesel    StationStats `yaml:"Diesel"`
	LPG       StationStats `yaml:"LPG"`
	Electric  StationStats `yaml:"Electric"`
	Registers StationStats `yaml:"Registers"`
}
