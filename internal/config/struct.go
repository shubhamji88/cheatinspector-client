package config

// Config details available throughout the system
type ConfigStruct struct {
	Name     string
	Server   string
	Feeder   string
	FileName string
	Debug    bool
	Ignore   []string
}
