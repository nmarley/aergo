// Code generated by "stringer -type=sysParamIndex"; DO NOT EDIT.

package system

import "strconv"

const _sysParamIndex_name = "bpCountnumBPsysParamMax"

var _sysParamIndex_index = [...]uint8{0, 7, 12, 23}

func (i sysParamIndex) String() string {
	if i < 0 || i >= sysParamIndex(len(_sysParamIndex_index)-1) {
		return "sysParamIndex(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _sysParamIndex_name[_sysParamIndex_index[i]:_sysParamIndex_index[i+1]]
}