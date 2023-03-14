package ports

type ConfigRepository interface {
	GetConfigFile() ([]byte, error)
}
