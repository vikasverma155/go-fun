// Code generated by "stringer -type=Pill"; DO NOT EDIT.

package fun

import "fmt"

const _Pill_name = "PlaceboAspirinIbuprofenParacetamol"

var _Pill_index = [...]uint8{0, 7, 14, 23, 34}

func (i Pill) String() string {
	if i < 0 || i >= Pill(len(_Pill_index)-1) {
		return fmt.Sprintf("Pill(%d)", i)
	}
	return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
