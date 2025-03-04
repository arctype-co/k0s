/*
Copyright 2022 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"fmt"
)

var _ Validateable = (*KubeProxy)(nil)

const (
	ModeIptables  = "iptables"
	ModeIPVS      = "ipvs"
	ModeUSerspace = "userspace"
)

// KubeProxy defines the configuration for kube-proxy
type KubeProxy struct {
	Disabled           bool   `json:"disabled,omitempty"`
	Mode               string `json:"mode,omitempty"`
	MetricsBindAddress string `json:"metricsBindAddress,omitempty"`
}

// DefaultKubeProxy creates the default config for kube-proxy
func DefaultKubeProxy() *KubeProxy {
	return &KubeProxy{
		Disabled:           false,
		Mode:               "iptables",
		MetricsBindAddress: "0.0.0.0:10249",
	}
}

// Validate validates kube proxy config
func (k *KubeProxy) Validate() []error {
	if k.Disabled {
		return nil
	}
	var errors []error
	if k.Mode != "iptables" && k.Mode != "ipvs" && k.Mode != "userspace" {
		errors = append(errors, fmt.Errorf("unsupported mode %s for kubeProxy config", k.Mode))
	}
	return errors
}
