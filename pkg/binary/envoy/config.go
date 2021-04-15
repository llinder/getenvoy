// Copyright 2019 Tetrate
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

package envoy

import (
	"fmt"
	"os"
	"path/filepath"
)

func newConfig() *Config {
	return &Config{AdminPort: 15000}
}

// Config store Envoy config information for arg mutators
type Config struct {
	AdminAddress string
	AdminPort    int32
}

// GetAdminAddress returns a host:port formatted address of the Envoy admin listener.
func (c *Config) GetAdminAddress() string {
	if c.AdminPort == 0 {
		return ""
	}
	address := c.AdminAddress
	if address == "" {
		address = "localhost"
	}
	return fmt.Sprintf("%s:%d", address, c.AdminPort)
}

// SaveConfig saves configuration string in getenvoy
// directory.
func (r *Runtime) SaveConfig(name, config string) (string, error) {
	configDir := filepath.Join(r.RootDir, "configs")
	if err := os.MkdirAll(configDir, 0750); err != nil {
		return "", fmt.Errorf("Unable to create directory %q: %v", configDir, err)
	}
	filename := name + ".yaml"
	err := os.WriteFile(filepath.Join(configDir, filename), []byte(config), 0600)
	if err != nil {
		return "", fmt.Errorf("Cannot save config file %s: %s", filepath.Join(configDir, filename), err)
	}
	return filepath.Join(configDir, filename), nil
}
