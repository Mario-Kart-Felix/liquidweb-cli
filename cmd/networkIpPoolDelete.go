/*
Copyright © LiquidWeb

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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/liquidweb/liquidweb-cli/types/api"
	"github.com/liquidweb/liquidweb-cli/validate"
)

var networkIpPoolDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an IP Pool",
	Long: `Delete an IP Pool.

An IP Pool is a range of nonintersecting, reusable IP addresses reserved to
your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		uniqIdFlag, _ := cmd.Flags().GetString("uniq_id")

		validateFields := map[interface{}]interface{}{
			uniqIdFlag: "UniqId",
		}
		if err := validate.Validate(validateFields); err != nil {
			lwCliInst.Die(err)
		}

		apiArgs := map[string]interface{}{
			"uniq_id": uniqIdFlag,
		}

		var details apiTypes.NetworkIpPoolDelete
		if err := lwCliInst.CallLwApiInto("bleed/network/pool/delete", apiArgs, &details); err != nil {
			lwCliInst.Die(err)
		}

		fmt.Printf("Deleted IP Pool %t\n", details.Deleted)
	},
}

func init() {
	networkIpPoolCmd.AddCommand(networkIpPoolDeleteCmd)

	networkIpPoolDeleteCmd.Flags().String("uniq_id", "", "uniq_id of IP Pool")

	networkIpPoolDeleteCmd.MarkFlagRequired("uniq_id")
}
