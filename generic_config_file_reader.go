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

package main

import (
	"log"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

type GenericConfigFileReader struct{}

func (*GenericConfigFileReader) MustLoadFile(filepath string) {
	config.AddDriver(yamlv3.Driver)

	err := config.LoadFiles(filepath)
	if err != nil {
		panic(err)
	}
}

func (*GenericConfigFileReader) MustHaveString(key string) string {
	value := config.String(key)

	if len(value) == 0 {
		log.Panic("Key not found or has no value in config file: ", key)
	}

	return value
}

func (*GenericConfigFileReader) MustHaveStringList(key string) []string {
	list := config.Strings(key)

	if len(list) == 0 {
		log.Panic("Key not found or has no list items in config file: ", key)
	}

	return list
}

func (*GenericConfigFileReader) MustHaveMap(key string) map[string]string {
	mmap := config.StringMap(key)

	if len(mmap) == 0 {
		log.Panic("Map not found or is empty: ", key)
	}

	return mmap
}
