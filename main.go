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
	cube := func(a float64) float64 {
		return a * a * a
	}
	_ = cube
	c2dt2 := func(r float64) float64 {
		const (
			r_s    = 33.0
			r_g    = .01
			c      = 299792458.0
			dt     = 1
			dr     = 1
			dtheta = 1
			theta  = 0
			dphi   = 0
		)
		//return (1-r_s/r)*square(c)*dt + square(dr)/(1-r_s/r) + square(r)*(square(dtheta)+square(math.Sin(theta))*square(dphi))
		//a := -square(3*math.Sqrt(1-r_s/r_g)-math.Sqrt(1-(square(r)*r_s)/cube(r_g))) * square(c) * square(dt) / 4
		//b := square(dr) / (1 - (square(r)*r_s)/cube(r_g))
		//d := square(r) * (square(dtheta) + square(math.Sin(theta))*square(dphi))
		//fmt.Println(a, b, d)
		//return a + b + d
		//return -square(-math.Sqrt(1-square(r))) + 1/(1-square(r))
		// https://physics.stackexchange.com/questions/530075/interpretation-of-the-interior-schwarzschild-solution
		return (r_s/r_g)*(9-(6*r)/r_g)*square(dt) - square(dr) +
			(r_s/r_g)*(square(r/r_g)*square(dr)-square(r_g/r)*square(dt)) -
			square(r)*(square(dtheta)+square(math.Sin(theta))*square(dphi))

	}
	for i := 1.0; i > 0; i -= .1 {
		fmt.Println(i, c2dt2(i), 1/math.Pow(i, 33))
	}
	//fmt.Println(0, c2dt2(0), 1/math.Pow(0, 33))
}
