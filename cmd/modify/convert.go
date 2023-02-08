// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package modify

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
	"storj.io/storj-up/cmd"
	"storj.io/storj-up/pkg/common"
	"strings"
)

func convertCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "convert <service>",
		Aliases: []string{"config"},
		Short:   "Print out available configuration for specific service",
		RunE: func(cmd *cobra.Command, args []string) error {
			selector, _, err := common.ParseArgumentsWithSelector(args, 0)
			if err != nil {
				return err
			}
			return convertConfig(selector)
		},
	}
}

func init() {
	cmd.RootCmd.AddCommand(convertCmd())
}

func convertConfig(services []string) error {
	fileName := "pkg/recipe/minimal.yaml"
	raw, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	y := map[string]interface{}{}
	err = yaml.Unmarshal(raw, &y)
	if err != nil {
		return err
	}

	for i, a := range y["add"].([]interface{}) {
		name := a.(map[string]interface{})["name"].(string)
		configs, err := readConfigs("/tmp/" + name + ".yaml")
		if err != nil {
			return err
		}
		cfg := y["add"].([]interface{})[i].(map[string]interface{})["config"].(map[string]interface{})
		replacement := map[string]interface{}{}
		for k, v := range cfg {
			option := findReplacement(configs, k)
			if option == "" {
				fmt.Println("missing", name, k)
				replacement[k] = v
				continue
			}
			replacement[option] = v
			//fmt.Println(configToEnvName(option.Name))
		}
		y["add"].([]interface{})[i].(map[string]interface{})["config"] = replacement
	}

	raw, err = yaml.Marshal(y)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, raw, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readConfigs(s string) (map[string]string, error) {
	res := map[string]string{}
	raw, err := ioutil.ReadFile(s)
	if err != nil {
		return res, err
	}
	desc := ""
	for _, l := range strings.Split(string(raw), "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		l = strings.TrimSpace(strings.TrimPrefix(l, "#"))
		if desc != "" {
			res[strings.SplitN(l, ":", 2)[0]] = desc
			desc = ""
		} else {
			desc = l
		}
	}
	return res, nil
}

func findReplacement(configs map[string]string, k string) string {
	k = strings.TrimPrefix(k, "STORJ_")
	for c, _ := range configs {
		//fmt.Println(configToEnvName(c.Name), k)
		if configToEnvName(c) == k {
			return c
		}
	}
	return ""
}

func configToEnvName(name string) string {
	smallCapital := regexp.MustCompile("([a-z])([A-Z])")
	name = smallCapital.ReplaceAllString(name, "${1}_$2")
	name = strings.ToUpper(name)
	name = strings.ReplaceAll(name, ".", "_")
	name = strings.ReplaceAll(name, "-", "_")
	return name

}
