package sftp

type Config struct {
	Username  string   `envconfig:"SFTP_USERNAME"`
	Password  Password `envconfig:"SFTP_PASSWORD"`
	Hostname  string   `envconfig:"SFTP_HOSTNAME"`
	Directory string   `envconfig:"SFTP_DIRECTORY"`
}

type Password string

func (Password) String() string {
	return "***"
}
