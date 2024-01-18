// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
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
//
// SPDX-License-Identifier: Apache-2.0

package golang

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"path"

	"github.com/specterops/bloodhound/packages/go/stbernard/analyzers/codeclimate"
	"github.com/specterops/bloodhound/slices"
)

var (
	ErrNonZeroExit = errors.New("non-zero exit status")
)

func Run(cwd string, modPaths []string, env []string) ([]codeclimate.Entry, error) {
	var (
		result []codeclimate.Entry
		args   = []string{"run", "--out-format", "code-climate", "--config", ".golangci.json", "--"}
		outb   bytes.Buffer
	)

	args = append(args, slices.Map(modPaths, func(modPath string) string {
		return path.Join(modPath, "...")
	})...)

	cmd := exec.Command("golangci-lint")
	cmd.Env = env
	cmd.Dir = cwd
	cmd.Stdout = &outb
	cmd.Args = append(cmd.Args, args...)

	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		err = ErrNonZeroExit
	} else {
		return result, fmt.Errorf("unexpected failure: %w", err)
	}

	if err := json.NewDecoder(&outb).Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode output: %w", err)
	}

	return result, err
}
