// Copyright 2023 Gabriele Iannetti <g.iannetti@gsi.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
)

// Reads a file and panics on error
func MustReadFile(file *string) string {
	data, err := os.ReadFile(*file)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func ExecuteCommandWithSudo(command string, args ...string) (*string, error) {
	cmdWithArgs := append([]string{command}, args...)

	cmd := exec.Command("sudo", cmdWithArgs...)

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	out, err := ioutil.ReadAll(pipe)
	if err != nil {
		return nil, err
	}

	// TODO: Timeout handling?
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	// TrimSpace on []bytes is more efficient than calling TrimSpace on a string since it creates a copy
	content := string(bytes.TrimSpace(out))

	return &content, nil
}
