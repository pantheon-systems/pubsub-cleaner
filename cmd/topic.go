// Copyright © 2016 Jesse Nelson <spheromak@gmail.com>
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

package cmd

import (
	"fmt"

	"github.com/pantheon-systems/pubsub-cleaner/pkg/pubsub"
	"github.com/spf13/cobra"
)

// keep is the keep string passed in from args
var keep string

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runTopic,
}

func runTopic(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Need to supply a topic name to operate on")
	}
	topic := args[0]
	if topic == "" {
		return fmt.Errorf("No topic specified")
	}

	noOP, err := cmd.Flags().GetBool("no-op")
	if err != nil {
		return err
	}

	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	if project == "" {
		return fmt.Errorf("No GCE project specified")
	}

	psConfig := pscleaner.Config{
		NoOP:    noOP,
		Topic:   topic,
		Keep:    keep,
		Project: project,
	}

	psc, err := pscleaner.NewCleaner(psConfig)
	if err != nil {
		return err
	}

	return psc.CleanTopicSubscriptions()
}

func init() {
	RootCmd.AddCommand(topicCmd)

	topicCmd.Flags().StringVarP(&keep, "keep", "k", "", "String match for subscriptions to keep")
}
