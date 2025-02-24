package config

// Config estructura que almacena la configuración de la aplicación
type Config struct {
	Port string // Puerto en el que se ejecutará el servidor
}

// Load carga la configuración y devuelve un puntero a la estructura Config
func Load() *Config {
	return &Config{
		Port: "8080", // Puerto predeterminado
	}
}
