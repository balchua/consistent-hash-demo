package config

type Infra struct {
	Nodes []Node `mapstructure:"nodes"`
}
type Node struct {
	Name string `mapstructure:"name"`
}

type ClusterConfig struct {
	Infra Infra `mapstructure:"infrastructure"`
}
