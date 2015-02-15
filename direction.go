// 2d tile space direction
package direction

import (
	"github.com/kasworld/go-abs"
	"github.com/kasworld/rand"
)

const (
	Dir_stop = iota
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

var vt2Dir = [3][3]uint8{}

func init() {
	for i, v := range Dir2Info {
		vt2Dir[v.Vt[0]+1][v.Vt[1]+1] = uint8(i)
	}
}

func Vt2Dir(x, y int) uint8 { // -1 ~ 1
	return vt2Dir[x+1][y+1]
}
func VtValidate(x, y int) bool {
	return x >= -1 && x <= 1 && y >= -1 && y <= 1
}

func TurnDir(dir uint8, turn int8) uint8 {
	if dir == Dir_stop {
		return dir
	}
	turn %= 8
	return uint8((int8(dir)-1+turn+8)%8 + 1)
}
func ReverseDir(dir uint8) uint8 {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(-vt[0], -vt[1])
}
func InverseX(dir uint8) uint8 {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(-vt[0], vt[1])
}
func InverseY(dir uint8) uint8 {
	vt := Dir2Info[dir].Vt
	return Vt2Dir(vt[0], -vt[1])
}

func RandDir(rnd *rand.Rand) uint8 {
	return uint8(rnd.IntRange(1, 9))
}

/////

func DxDy2Dir8(dx, dy int) uint8 {
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

func DxDy2Dir4(dx, dy int) uint8 {
	dxs, dxv := abs.SignAbsi(dx)
	dys, dyv := abs.SignAbsi(dy)
	if dxv >= dyv {
		dys = 0
	} else {
		dxs = 0
	}
	return Vt2Dir(dxs, dys)
}
func DxDy2Dir(dx, dy int) uint8 {
	dxs, _ := abs.SignAbsi(dx)
	dys, _ := abs.SignAbsi(dy)
	return Vt2Dir(dxs, dys)
}
