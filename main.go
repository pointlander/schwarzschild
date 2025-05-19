// Copyright 2025 The Schwarzschild Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
)

func main() {
	square := func(a float64) float64 {
		return a * a
	}
	ds2 := func(r float64) float64 {
		const (
			r_s    = 33.0
			c      = 299792458.0
			dt     = 1
			dr     = 1
			dtheta = 1
			theta  = 0
			dphi   = 0
		)
		return (1-r_s/r)*square(c)*dt + square(dr)/(1-r_s/r) + square(r)*(square(dtheta)+square(math.Sin(theta))*square(dphi))
	}
	for i := 1.0; i > 0; i -= .01 {
		fmt.Println(ds2(i), -1/math.Pow(i, 33))
	}
}
