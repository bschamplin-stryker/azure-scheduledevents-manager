package drainmanager

import (
	"github.com/webdevops/azure-scheduledevents-manager/azuremetadata"
	"github.com/webdevops/azure-scheduledevents-manager/config"
)

type DrainManagerNoop struct {
	DrainManager
	Conf         config.Opts
	instanceName string
}

func (m *DrainManagerNoop) SetInstanceName(name string) {
	m.instanceName = name
}

func (m *DrainManagerNoop) InstanceName() string {
	return m.instanceName
}

func (m *DrainManagerNoop) Test() {
}

func (m *DrainManagerNoop) Drain(event *azuremetadata.AzureScheduledEvent) bool {
	return true
}

func (m *DrainManagerNoop) Uncordon() bool {
	return true
}
