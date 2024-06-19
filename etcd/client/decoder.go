package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ConfigParser interface {
	Decode(data []byte, config *EtcdConfig) error
}

type defaultParser struct {
}

func (p *defaultParser) Decode(data []byte, config *EtcdConfig) error {
	return json.Unmarshal(data, config)
}

type Config interface {
	String() string
}

type EtcdConfig struct {
	ClientBasicInfo *EndpointBasicInfo `mapstructure:"ClientBasicInfo"`
	HostPorts       []string           `mapstructure:"HostPorts"`
	DestService     *string            `mapstructure:"DestService"`
	Protocol        *string            `mapstructure:"Protocol"`
	Connection      *Connection        `mapstructure:"Connection"`
	MyConfig        Config             `mapstructure:"MyConfig"`
}

func (c *EtcdConfig) String() string {
	var builder strings.Builder

	if c.ClientBasicInfo != nil {
		builder.WriteString(fmt.Sprintf("ClientBasicInfo: %v\n", *c.ClientBasicInfo))
	}

	if c.HostPorts != nil {
		builder.WriteString(fmt.Sprintf("HostPorts: %v\n", c.HostPorts))
	}

	if c.DestService != nil {
		builder.WriteString(fmt.Sprintf("DestService: %v\n", *c.DestService))
	}

	if c.Protocol != nil {
		builder.WriteString(fmt.Sprintf("Protocol: %v\n", *c.Protocol))
	}

	if c.Connection != nil {
		builder.WriteString(fmt.Sprintf("Connection: %v\n", *c.Connection))
	}

	if c.MyConfig != nil {
		builder.WriteString(c.MyConfig.String())
	}

	return builder.String()
}

type EndpointBasicInfo struct {
	ServiceName string            `mapstructure:"ServiceName"`
	Method      string            `mapstructure:"Method"`
	Tags        map[string]string `mapstructure:"Tags"`
}

type IdleConfig struct {
	MinIdlePerAddress int    `mapstructure:"MinIdlePerAddress"`
	MaxIdlePerAddress int    `mapstructure:"MaxIdlePerAddress"`
	MaxIdleGlobal     int    `mapstructure:"MaxIdleGlobal"`
	MaxIdleTimeout    string `mapstructure:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `mapstructure:"ConnNum"`
}

type Connection struct {
	Method         string        `mapstructure:"Method"`
	LongConnection IdleConfig    `mapstructure:"LongConnection"`
	MuxConnection  MuxConnection `mapstructure:"MuxConnection"`
}
