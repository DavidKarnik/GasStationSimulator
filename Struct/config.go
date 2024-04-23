package Struct

type Config struct {
	Cars struct {
		Count          int `yaml:"count"`
		ArrivalTimeMin int `yaml:"arrival_time_min"`
		ArrivalTimeMax int `yaml:"arrival_time_max"`
	} `yaml:"cars"`
	Stations struct {
		Gas struct {
			Count        int `yaml:"count"`
			ServeTimeMin int `yaml:"serve_time_min"`
			ServeTimeMax int `yaml:"serve_time_max"`
		} `yaml:"gas"`
		Diesel struct {
			Count        int `yaml:"count"`
			ServeTimeMin int `yaml:"serve_time_min"`
			ServeTimeMax int `yaml:"serve_time_max"`
		} `yaml:"diesel"`
		Lpg struct {
			Count        int `yaml:"count"`
			ServeTimeMin int `yaml:"serve_time_min"`
			ServeTimeMax int `yaml:"serve_time_max"`
		} `yaml:"lpg"`
		Electric struct {
			Count        int `yaml:"count"`
			ServeTimeMin int `yaml:"serve_time_min"`
			ServeTimeMax int `yaml:"serve_time_max"`
		} `yaml:"electric"`
	} `yaml:"stations"`
	Registers struct {
		Count         int `yaml:"count"`
		HandleTimeMin int `yaml:"handle_time_min"`
		HandleTimeMax int `yaml:"handle_time_max"`
	} `yaml:"registers"`
}
