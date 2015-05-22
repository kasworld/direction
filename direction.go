// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// 2d tile space direction
package direction

import (
	"github.com/kasworld/go-abs"
	"github.com/kasworld/rand"
)

type Dir_Type uint8

const (
	Dir_stop = Dir_Type(iota)
	Dir_n
	Dir_ne
	Dir_e
	Dir_se
	Dir_s
	Dir_sw
	Dir_w
	Dir_nw
)

var Dir2Info = []struct {
	Name string
	Vt   [2]int
	Len  float64
}{
	Dir_stop: {".", [2]int{0, 0}, 0.0},
	Dir_n:    {"N", [2]int{0, -1}, 1.0},
	Dir_ne:   {"NE", [2]int{1, -1}, 1.4},
	Dir_e:    {"E", [2]int{1, 0}, 1.0},
	Dir_se:   {"SE", [2]int{1, 1}, 1.4},
	Dir_s:    {"S", [2]int{0, 1}, 1.0},
	Dir_sw:   {"SW", [2]int{-1, 1}, 1.4},
	Dir_w:    {"W", [2]int{-1, 0}, 1.0},
	Dir_nw:   {"NW", [2]int{-1, -1}, 1.4},
}

var vt2Dir = [3][3]Dir_Type{}

func init() {
	for i, v := range Dir2Info {
		vt2Dir[v.Vt[0]+1][v.Vt[1]+1] = Dir_Type(i)
	}
}

func Vt2Dir(x, y int) Dir_Type { // -1 ~ 1
	return vt2Dir[x+1][y+1]
}
func VtValidate(x, y int) bool {
	return x >= -1 && x <= 1 && y >= -1 && y <= 1
}

func TurnDir(dir Dir_Type, turn int8) Dir_Type {
	if dir == Dir_stop {
		return dir
	}
	turn %= 8
	return Dir_Type((int8(dir)-1+turn+8)%8 + 1)
}
func ReverseDir(dir Dir_Type) Dir_Type {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(-vt[0], -vt[1])
}
func InverseX(dir Dir_Type) Dir_Type {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(-vt[0], vt[1])
}
func InverseY(dir Dir_Type) Dir_Type {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(vt[0], -vt[1])
}

func RandDir(rnd *rand.Rand) Dir_Type {
	return Dir_Type(rnd.IntRange(1, 9))
}

/////

// find remote pos direction 8way
func DxDy2Dir8(dx, dy int) Dir_Type {
	dxs, dxv := abs.SignAbsi(dx)
	dys, dyv := abs.SignAbsi(dy)
	if dxv > dyv*2 {
		dys = 0
	}
	if dyv > dxv*2 {
		dxs = 0
	}
	return Vt2Dir(dxs, dys)
}

// find remote pos direction 4way
func DxDy2Dir4(dx, dy int) Dir_Type {
	dxs, dxv := abs.SignAbsi(dx)
	dys, dyv := abs.SignAbsi(dy)
	if dxv >= dyv {
		dys = 0
	} else {
		dxs = 0
	}
	return Vt2Dir(dxs, dys)
}

// contact only
func DxDy2Dir(dx, dy int) Dir_Type {
	dxs, _ := abs.SignAbsi(dx)
	dys, _ := abs.SignAbsi(dy)
	return Vt2Dir(dxs, dys)
}
