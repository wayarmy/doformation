package resources

// DoDroplet contain Droplet's components
type DoDroplet struct {
	Size           elementWrapper `yaml:"size"`
	Name           elementWrapper `yaml:"name"`
	BlockStorage   blockStorage   `yaml:"block_storage"`
	Region         elementWrapper `yaml:"region"`
	PrivateNetwork elementWrapper `yaml:"private_network"`
	Backups        elementWrapper `yaml:"backups"`
	IPv6           elementWrapper `yaml:"ipv6"`
	UserData       elementWrapper `yaml:"user_data"`
	Monitoring     elementWrapper `yaml:"monitoring"`
	KeyName        elementWrapper `yaml:"key_name"`
}

type blockStorage struct {
	Size elementWrapper `yaml:"size"`
}
