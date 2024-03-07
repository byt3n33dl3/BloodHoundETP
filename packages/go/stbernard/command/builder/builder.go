// Copyright 2024 Specter Ops, Inc.
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

package builder

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/specterops/bloodhound/packages/go/stbernard/environment"
	"github.com/specterops/bloodhound/packages/go/stbernard/workspace"
	"github.com/specterops/bloodhound/packages/go/stbernard/yarn"
)

const (
	Name  = "build"
	Usage = "Build commands in current workspace"
)

type Config struct {
	Environment environment.Environment
}

type command struct {
	config Config
}

func (s command) Usage() string {
	return Usage
}

func (s command) Name() string {
	return Name
}

func (s command) Run() error {
	if cwd, err := workspace.FindRoot(); err != nil {
		return fmt.Errorf("could not find workspace root: %w", err)
	} else if err := s.runJSBuild(cwd); err != nil {
		return fmt.Errorf("could not build JS artifacts: %w", err)
	} else if err := s.runGoBuild(cwd); err != nil {
		return fmt.Errorf("could not build Go artifacts: %w", err)
	} else {
		return nil
	}
}

func Create(config Config) (command, error) {
	cmd := flag.NewFlagSet(Name, flag.ExitOnError)

	cmd.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "%s\n\nUsage: %s %s [OPTIONS]\n\nOptions:\n", Usage, filepath.Base(os.Args[0]), Name)
		cmd.PrintDefaults()
	}

	if err := cmd.Parse(os.Args[2:]); err != nil {
		cmd.Usage()
		return command{}, fmt.Errorf("failed to parse build command: %w", err)
	} else {
		return command{config: config}, nil
	}
}

func (s command) runJSBuild(cwd string) error {
	var env = s.config.Environment

	if _, ok := env["BUILD_PATH"]; !ok {
		env["BUILD_PATH"] = filepath.Join(cwd, "dist", "bh-ui")
	}

	if jsPaths, err := workspace.ParseJSAbsPaths(cwd); err != nil {
		return fmt.Errorf("could not retrieve JS paths: %w", err)
	} else if err := yarn.InstallWorkspaceDeps(jsPaths, env.Slice()); err != nil {
		return fmt.Errorf("could not install JS deps: %w", err)
	} else if err := yarn.BuildWorkspace(cwd, env.Slice()); err != nil {
		return fmt.Errorf("could not build JS workspace: %w", err)
	} else {
		return nil
	}
}

func (s command) runGoBuild(cwd string) error {
	if modPaths, err := workspace.ParseModulesAbsPaths(cwd); err != nil {
		return fmt.Errorf("could not parse module absolute paths: %w", err)
	} else if err := workspace.BuildGoMainPackages(cwd, modPaths); err != nil {
		return fmt.Errorf("could not build main packages: %w", err)
	} else {
		return nil
	}
}
