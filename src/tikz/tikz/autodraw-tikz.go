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
package tikz

import (
	"fmt"
	"image/color"
	"compiler/operation"
	"compiler/instruction"
)

type Tikz struct {
	scale float64
	offsetx int16
	offsety int16

	instlist []*instruction.Instruction
}

type TikzError struct {
	reason string
}

func NewTikzError(reason string) *TikzError {
	return &TikzError{reason}
}

func (e *TikzError) Error() string {
	return e.reason
}

func NewTikz() *Tikz{
	tz := new(Tikz)
	tz.scale = 1.0
	return tz
}

func (tz *Tikz) Update(inst *instruction.Instruction) error {
	tz.instlist = append(tz.instlist,inst)
	return nil
}

func (tz *Tikz) GenerateTikzCode() (string,error) {
	code := ""
	defs := ""
	defmap := make(map[string]string)

	for _,inst := range tz.instlist {
		tikzCode,err := InstToTikz(inst,tz.scale)
		if err != nil {
			return "",err
		}
		if tikzCode != "" {
			code += fmt.Sprintf("  %s\n",tikzCode)
		}
		def := InstToDef(inst,defmap)
		if def != "" {
			defs += fmt.Sprintf("  %s\n",def)
		}
	}

	return fmt.Sprintf("\\begin{tikzpicture}\n%s%s\\end{tikzpicture}\n",defs,code),nil
}

func InstToTikz(inst *instruction.Instruction, scale float64) (string,error) {
	switch inst.Command {
	case instruction.LINE_STRIP:
		return fmt.Sprintf("\\draw %s;",
		  GenerateFloatPairs(IntsToScaledFloats(inst.Args,scale))),nil
	case instruction.CURVE:
		return fmt.Sprintf("\\draw %s;",
		  GenerateFloatPairsCurve(IntsToScaledFloats(inst.Args,scale))),nil
	case instruction.TEXT:
		return fmt.Sprintf("\\node[scale=%g] at (%g,%g) {%s};",
		  scale*float64(inst.Args[2])/256.0,float64(inst.Args[0])/100.0,float64(inst.Args[1])/100.0,GenerateString(inst.Args[3:])),nil
	case instruction.NODE:
		c := operation.GetColor(inst.Args[2])
		colorString := GetColorString(c)
		code := fmt.Sprintf("\\node[scale=%g,draw,fill=%s] at (%g,%g) {%s};",
		  scale,colorString,float64(inst.Args[0])/100.0,float64(inst.Args[1])/100.0,GenerateString(inst.Args[3:]))
		return code,nil
	default:
		return "",NewTikzError("invalid instruction: "+inst.ToString())
	}
}

func InstToDef(inst *instruction.Instruction,defmap map[string]string) string {
	switch inst.Command {
	case instruction.NODE:
		c := operation.GetColor(inst.Args[2])
		colorCode := GetColorCode(c)
		def := GetColorDef(c,colorCode,defmap)
		return def
	default:
		return ""
	}
}

func GetColorDef(c color.Color, colorString string, defmap map[string]string) string {
	_,ok := defmap[colorString]
	if ok {
		return ""
	}
	r,g,b,_ := c.RGBA()
	def := fmt.Sprintf("\\definecolor{%s}{RGB}{%d,%d,%d}",colorString,r>>8,g>>8,b>>8)
	defmap[colorString] = def
	return def
}

func GetColorCode(c color.Color) string {
	r,g,b,_ := c.RGBA()
	return fmt.Sprintf("r%dg%db%d",r>>8,g>>8,b>>8)
}

func GetColorString(c color.Color) string {
	r,g,b,a := c.RGBA()
	return fmt.Sprintf("r%dg%db%d!%d",r>>8,g>>8,b>>8,a*100/65535)
}

func GenerateString(args []int16) string {
	ret := make([]byte,len(args))
	for i,c := range args {
		ret[i] = byte(c)
	}
	return string(ret)
}

func IntsToFloats(args []int16) []float64 {
	ret := make([]float64,len(args))
	for i,v := range args {
		ret[i] = float64(v)
	}
	return ret
}

func IntsToScaledFloats(args []int16, scale float64) []float64 {
	ret := make([]float64,len(args))
	for i,v := range args {
		ret[i] = float64(v)/100.0 * scale
	}
	return ret
}

func GenerateFloatPairs(args []float64) string {
	if len(args) == 0 || len(args)%2 == 1{
		return ""
	}
	connect := "--"
	num := len(args)/2
	ret := fmt.Sprintf("(%g,%g)",args[0],args[1])
	for i := 1; i < num; i++ {
		ret += fmt.Sprintf(fmt.Sprintf(" %s %s",connect,"(%g,%g)"),
			args[2*i],args[2*i+1])
	}
	return ret
}

func GenerateFloatPairsCurve(args []float64) string {
	if len(args) < 8 || len(args)%6 != 2{
		return ""
	}
	args = append(args,args[0],args[1])
	num := (len(args)/2-1)/3+1
	ret := fmt.Sprintf("(%g,%g)",args[0],args[1])
	for i := 0; i < num-1; i++ {
		x0,y0 := args[i*6+2],args[i*6+3]
		x1,y1 := args[i*6+4],args[i*6+5]
		x2,y2 := args[i*6+6],args[i*6+7]
		ret += fmt.Sprintf(" .. controls (%g,%g) and (%g,%g) .. (%g,%g)",
			x0,y0,x1,y1,x2,y2)
	}
	return ret
}
