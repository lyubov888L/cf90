package main

import (
	"encoding/json"
	"errors"
	"github.com/igm/cf"
	"os"
)

type Config struct {
	filename string
	trace    bool
	data     struct {
		ActiveTarget int       `json:"selected"`
		Targets      []*Target `json:"targets"`
	}
}

func NewConfig(filename string) (ret *Config, err error) {
	ret = new(Config)
	ret.filename = filename
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ret.data)
	return
}

func (c *Config) Save() (err error) {
	file, err := os.Create(c.filename)
	if err != nil {
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(&c.data)
	return
}

func NewTarget(host, alias string) *Target {
	return &Target{Target: cf.NewTarget(host), Alias: alias}
}

func (c *Config) AddTarget(target *Target) error {
	for _, existing := range c.data.Targets {
		if existing.Alias == target.Alias {
			return errors.New("Duplicate alias.")
		}
	}
	c.data.Targets = append(c.data.Targets, target)
	return nil
}

func (c *Config) RemoveTarget(target *Target) {
	for i, existing := range c.data.Targets {
		if existing.TargetUrl == target.TargetUrl {
			c.data.Targets = append(c.data.Targets[:i], c.data.Targets[i+1:]...)
		}
	}
}

func (c *Config) SelectedTarget() (*Target, error) {
	if c.data.ActiveTarget >= len(c.data.Targets) {
		return nil, errors.New("Not target selected.")
	}
	ret := c.data.Targets[c.data.ActiveTarget]
	return ret, nil
}

func (c *Config) Select(host string) error {
	for i, existing := range c.data.Targets {
		if existing.TargetUrl == host {
			c.data.ActiveTarget = i
			return nil
		}
	}
	return errors.New("Target does not  exist")
}

func (c *Config) GetTarget(host string) (*Target, error) {
	for _, existing := range c.data.Targets {
		if existing.TargetUrl == host {
			return existing, nil
		}
	}
	return nil, errors.New("Target does not  exist")
}

func (c *Config) GetTargetByAlias(alias string) (*Target, error) {
	for _, existing := range c.data.Targets {
		if existing.Alias == alias {
			return existing, nil
		}
	}
	return nil, errors.New("Target does not  exist")
}
