package header

import "github.com/use-go/gosips/core"

/**
* A list of supported headers.
*@see Supported
 */

type SupportedList struct {
	SIPHeaderList
}

/** Default Constructor
 */
func NewSupportedList() *SupportedList {
	this := &SupportedList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_SUPPORTED)
	return this
}
