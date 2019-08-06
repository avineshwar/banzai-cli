// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dns

import (
	"github.com/goph/emperror"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func validateSpec(specObj map[string]interface{}) error {
	providers := map[string]bool{
		"route53": true,
	}

	var spec struct {
		AutoDNS *struct {
			Enabled bool `mapstructure:"enabled"`
		} `mapstructure:"autoDns"`
		CustomDNS *struct {
			Enabled       bool     `mapstructure:"enabled"`
			DomainFilters []string `mapstructure:"domainFilters"`
			ClusterDomain string   `mapstructure:"clusterDomain"`
			Provider      *struct {
				Name     string                 `mapstructure:"name"`
				Options  map[string]interface{} `mapstructure:"options"`
				SecretID string                 `mapstructure:"secretId"`
			} `mapstructure:"provider"`
		} `mapstructure:"customDns"`
	}

	if err := mapstructure.Decode(specObj, &spec); err != nil {
		return emperror.Wrap(err, "feature specification does not conform to schema")
	}
	if spec.AutoDNS != nil && spec.AutoDNS.Enabled && spec.CustomDNS != nil && spec.CustomDNS.Enabled {
		return errors.New("Cannot enable 'autoDns' and 'customDns' at the same time")
	}
	if spec.CustomDNS != nil && spec.CustomDNS.Enabled {
		if len(spec.CustomDNS.DomainFilters) == 0 {
			return errors.New("At least one domain filter must be specified")
		}
		for _, df := range spec.CustomDNS.DomainFilters {
			if df == "" {
				return errors.New("Domain filter must not be empty")
			}
		}
		if spec.CustomDNS.ClusterDomain == "" {
			return errors.New("Cluster domain must not be empty")
		}
		if spec.CustomDNS.Provider == nil {
			return errors.New("Provider must be specified")
		}
		if !providers[spec.CustomDNS.Provider.Name] {
			return errors.Errorf("Unsupported provider %q", spec.CustomDNS.Provider.Name)
		}
		if spec.CustomDNS.Provider.SecretID == "" {
			return errors.New("Secret ID must be specified")
		}
	}
	return nil
}