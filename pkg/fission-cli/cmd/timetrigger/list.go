/*
Copyright 2019 The Fission Authors.

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

package timetrigger

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fission/fission/pkg/controller/client"
	"github.com/fission/fission/pkg/fission-cli/cliwrapper/cli"
	flagkey "github.com/fission/fission/pkg/fission-cli/flag/key"
	"github.com/fission/fission/pkg/fission-cli/util"
	"github.com/pkg/errors"
)

type ListSubCommand struct {
	client *client.Client
}

func List(input cli.Input) error {
	c, err := util.GetServer(input)
	if err != nil {
		return err
	}
	opts := ListSubCommand{
		client: c,
	}
	return opts.do(input)
}

func (opts *ListSubCommand) do(input cli.Input) error {
	ttNs := input.String(flagkey.NamespaceTrigger)
	tts, err := opts.client.TimeTriggerList(ttNs)
	if err != nil {
		return errors.Wrap(err, "list Time triggers")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "CRON", "FUNCTION_NAME")
	for _, tt := range tts {
		fmt.Fprintf(w, "%v\t%v\t%v\n",
			tt.Metadata.Name, tt.Spec.Cron, tt.Spec.FunctionReference.Name)
	}
	w.Flush()

	return nil
}
