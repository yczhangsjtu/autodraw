// This file is part of autodraw.
//
// Autodraw is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Autodraw is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with autodraw.  If not, see <http://www.gnu.org/licenses/>.
package operation

import (
	"image/color"
)

func ValidName(name string) bool {
	pattern := GetVariableRegexp()
	return pattern.MatchString(name)
}

func GetColor(c int16) color.Color {
	code := uint16(c)
	r,g,b,a := (code&0xf000)>>12,(code&0xf00)>>8,(code&0xf0)>>4,(code&0xf)
	return color.RGBA{(uint8(r)<<4)|uint8(r),(uint8(g)<<4)|uint8(g),(uint8(b)<<4)|uint8(b),(uint8(a)<<4)|uint8(a)}
}
