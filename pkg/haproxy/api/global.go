package api

import (
	"fmt"

	"github.com/haproxytech/client-native/v3/models"
	parser "github.com/haproxytech/config-parser/v4"
	"github.com/haproxytech/config-parser/v4/types"

	"github.com/haproxytech/kubernetes-ingress/pkg/utils"
)

func (c *clientNative) DefaultsGetConfiguration() (defaults *models.Defaults, err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return nil, err
	}
	_, defaults, err = configuration.GetDefaultsConfiguration(c.activeTransaction)
	if err != nil {
		return nil, fmt.Errorf("unable to get HAProxy's defaults section: %w", err)
	}
	return
}

func (c *clientNative) DefaultsPushConfiguration(defaults models.Defaults) (err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}
	err = configuration.PushDefaultsConfiguration(&defaults, c.activeTransaction, 0)
	if err != nil {
		return fmt.Errorf("unable to update HAProxy's defaults section: %w", err)
	}
	// Force defaults log directive to "log global"
	_ = configuration.DeleteLogTarget(0, string(parser.Defaults), parser.DefaultSectionName, c.activeTransaction, 0)
	err = configuration.CreateLogTarget(string(parser.Defaults), parser.DefaultSectionName, &models.LogTarget{Index: utils.PtrInt64(0), Global: true}, c.activeTransaction, 0)
	if err != nil {
		return fmt.Errorf("unable to set 'log global' directive in defaults section: %w", err)
	}
	return
}

func (c *clientNative) GlobalCfgSnippet(value []string) (err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}
	var config parser.Parser
	config, err = configuration.GetParser(c.activeTransaction)
	if err != nil {
		return
	}
	if len(value) == 0 {
		err = config.Set(parser.Global, parser.GlobalSectionName, "config-snippet", nil)
	} else {
		err = config.Set(parser.Global, parser.GlobalSectionName, "config-snippet", types.StringSliceC{Value: value})
	}
	if err != nil {
		return fmt.Errorf("unable to update global config snippet: %w", err)
	}
	return
}

func (c *clientNative) GlobalGetLogTargets() (lg models.LogTargets, err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return nil, err
	}
	c.activeTransactionHasChanges = true
	_, lg, err = configuration.GetLogTargets("global", parser.GlobalSectionName, c.activeTransaction)
	if err != nil {
		return lg, fmt.Errorf("unable to get HAProxy's global log targets: %w", err)
	}
	return
}

func (c *clientNative) GlobalPushLogTargets(logTargets models.LogTargets) error {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}
	c.activeTransactionHasChanges = true
	for {
		err = configuration.DeleteLogTarget(0, "global", parser.GlobalSectionName, c.activeTransaction, 0)
		if err != nil {
			break
		}
	}
	for _, log := range logTargets {
		err = configuration.CreateLogTarget(string(parser.Global), parser.GlobalSectionName, log, c.activeTransaction, 0)
		if err != nil {
			return fmt.Errorf("unable to update HAProxy's global log targets: %w", err)
		}
	}
	return nil
}

func (c *clientNative) GlobalGetConfiguration() (*models.Global, error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return nil, err
	}
	_, global, err := configuration.GetGlobalConfiguration(c.activeTransaction)
	if err != nil {
		return nil, fmt.Errorf("unable to get HAProxy's global section: %w", err)
	}
	return global, err
}

func (c *clientNative) GlobalPushConfiguration(global models.Global) (err error) {
	configuration, err := c.nativeAPI.Configuration()
	if err != nil {
		return err
	}
	err = configuration.PushGlobalConfiguration(&global, c.activeTransaction, 0)
	if err != nil {
		return fmt.Errorf("unable to update HAProxy's global section: %w", err)
	}
	return
}
