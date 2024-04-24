package Struct

// StationStats - Struct for output yaml
type StationStats struct {
	TotalCars    int `yaml:"total_cars"`
	TotalTime    int `yaml:"total_time"`
	AvgQueueTime int `yaml:"avg_queue_time"`
	MaxQueueTime int `yaml:"max_queue_time"`
}
