// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_test

import (
	"fmt"

	"github.com/tklauser/go-sysconf"
)

func ExampleSysconfClkTck() {
	clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		fmt.Printf("SC_CLK_TCK: %v\n", clktck)
	}
}

func ExampleSysconfUnsupported() {
	_, err := sysconf.Sysconf(-1)
	fmt.Print(err)

	// Output: invalid parameter value
}