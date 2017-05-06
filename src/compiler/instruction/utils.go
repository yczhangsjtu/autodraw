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
package instruction

func composeBytes(b1,b2 byte) int16 {
	return int16(uint(b1)*256+uint(b2))
}

func getInt16(data []byte, i *int) (int16,error) {
	if (*i)*2+1 >= len(data) {
		return 0,NewInstructionError("index out of range")
	}
	ret := composeBytes(data[(*i)*2],data[(*i)*2+1])
	(*i)++
	return ret,nil
}

func setUint16(data []byte, i *int, v uint16) error {
	if (*i)*2+1 >= len(data) {
		return NewInstructionError("index out of range")
	}
	data[(*i)*2]   = byte(v/256)
	data[(*i)*2+1] = byte(v%256)
	(*i)++
	return nil
}

func getInts16(data []byte, ptr *int, count int) ([]int16,error) {
	ret := make([]int16,count)
	var err error
	for i := 0; i < count; i++ {
		ret[i],err = getInt16(data,ptr)
		if err != nil {
			return ret,err
		}
	}
	return ret,err
}

func addLengthPrefix(args []int16) []int16 {
	return append([]int16{int16(len(args))},args...)
}
