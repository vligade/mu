package cli

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stelligent/mu/common"
)

func TestNewServicesCommand(t *testing.T) {
	assert := assert.New(t)

	config := common.NewConfig()

	command := newServicesCommand(config)

	assert.NotNil(command)
	assert.Equal("service", command.Name, "Name should match")
	assert.Equal(1, len(command.Aliases), "Aliases len should match")
	assert.Equal("svc", command.Aliases[0], "Aliases should match")
	assert.Equal("options for managing services", command.Usage, "Usage should match")
	assert.Equal(4, len(command.Subcommands), "Subcommands len should match")
}

func TestNewServicesShowCommand(t *testing.T) {
	assert := assert.New(t)

	config := common.NewConfig()

	command := newServicesShowCommand(config)

	assert.NotNil(command)
	assert.Equal("show", command.Name, "Name should match")
	assert.Equal(1, len(command.Flags), "Flags length")
	assert.Equal("service, s", command.Flags[0].GetName(), "Flags Name")
	assert.NotNil(command.Action)
}

func TestNewServicesDeployCommand(t *testing.T) {
	assert := assert.New(t)

	config := common.NewConfig()

	command := newServicesDeployCommand(config)

	assert.NotNil(command)
	assert.Equal("deploy", command.Name, "Name should match")
	assert.Equal("<environment>", command.ArgsUsage, "ArgsUsage should match")
	assert.Equal(1, len(command.Flags), "Flags length")
	assert.Equal("service, s", command.Flags[0].GetName(), "Flags Name")
	assert.NotNil(command.Action)
}

func TestNewSetenvCommand(t *testing.T) {
	assert := assert.New(t)

	config := common.NewConfig()

	command := newServicesSetenvCommand(config)

	assert.NotNil(command)
	assert.Equal("setenv", command.Name, "Name should match")
	assert.Equal("<environment> <key1>=<value1>...", command.ArgsUsage, "ArgsUsage should match")
	assert.Equal(1, len(command.Flags), "Flags length")
	assert.Equal("service, s", command.Flags[0].GetName(), "Flags Name")
	assert.NotNil(command.Action)
}

func TestNewUndeployCommand(t *testing.T) {
	assert := assert.New(t)

	config := common.NewConfig()

	command := newServicesUndeployCommand(config)

	assert.NotNil(command)
	assert.Equal("undeploy", command.Name, "Name should match")
	assert.Equal("<environment>", command.ArgsUsage, "ArgsUsage should match")
	assert.Equal(1, len(command.Flags), "Flags length")
	assert.Equal("service, s", command.Flags[0].GetName(), "Flags Name")
	assert.NotNil(command.Action)
}

