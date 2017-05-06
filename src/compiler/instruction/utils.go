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

func expandRect(args []int16) []int16 {
	ret := make([]int16,8)
	ret[0],ret[1] = args[0],args[1]
	ret[2],ret[3] = args[0],args[3]
	ret[4],ret[5] = args[2],args[3]
	ret[6],ret[7] = args[2],args[1]
	return ret
}

func expandOval(args []int16) []int16 {
	ret := make([]int16,16)
	x,y,a,b := args[0],args[1],args[2],args[3]
	ret[0],ret[1] = x+a,y
	ret[2],ret[3] = x+a,y+b
	ret[4],ret[5] = x,y+b
	ret[6],ret[7] = x-a,y+b
	ret[8],ret[9] = x-a,y
	ret[10],ret[11] = x-a,y-b
	ret[12],ret[13] = x,y-b
	ret[14],ret[15] = x+a,y-b
	return ret
}
