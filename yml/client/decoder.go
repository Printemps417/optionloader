package client

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpctimeout"
	"strings"
	"time"
)

// EndpointBasicInfo should be immutable after created.
type EndpointBasicInfo struct {
	ServiceName string            `yaml:"ServiceName"`
	Method      string            `yaml:"Method"`
	Tags        map[string]string `yaml:"Tags"`
}

type port string

type protocol string
type DestService string

// IdleConfig contains idle configuration for long-connection pool.
type IdleConfig struct {
	MinIdlePerAddress int           `yaml:"MinIdlePerAddress"`
	MaxIdlePerAddress int           `yaml:"MaxIdlePerAddress"`
	MaxIdleGlobal     int           `yaml:"MaxIdleGlobal"`
	MaxIdleTimeout    time.Duration `yaml:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `yaml:"connNum"`
}

type Connection struct {
	Method         string        `yaml:"method"`
	LongConnection IdleConfig    `yaml:"LongConnection"`
	MuxConnection  MuxConnection `yaml:"MuxConnection"`
}

// RPCTimeout
type Timeout rpctimeout.RPCTimeout

// RetryPolicy
type RetryPolicy retry.Policy

type Config interface {
	String() string
}

// CircuitBreaker
type Circuitbreaker circuitbreak.CBConfig
type YMLConfig struct {
	ClientBasicInfo *EndpointBasicInfo `yaml:"ClientBasicInfo"`
	HostPorts       []string           `yaml:"HostPorts"`
	DestService     *string            `yaml:"DestService"`
	Protocol        *string            `yaml:"Protocol"`
	Connection      *Connection        `yaml:"Connection"`
	CustomConfig    Config             `yaml:"CustomConfigConfig"`
}

func (c *YMLConfig) String() string {
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

	if c.CustomConfig != nil {
		builder.WriteString(c.CustomConfig.String())
	}

	return builder.String()
}
