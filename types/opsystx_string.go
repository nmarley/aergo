// Code generated by "stringer -type=OpSysTx"; DO NOT EDIT.

package types

import "strconv"

const _OpSysTx_name = "OpvoteBPOpvoteProposalOpstakeOpunstakeOpcreateProposalOpSysTxMax"

var _OpSysTx_index = [...]uint8{0, 8, 22, 29, 38, 54, 64}

func (i OpSysTx) String() string {
	if i < 0 || i >= OpSysTx(len(_OpSysTx_index)-1) {
		return "OpSysTx(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpSysTx_name[_OpSysTx_index[i]:_OpSysTx_index[i+1]]
}