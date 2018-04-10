package header

import "github.com/use-go/gosips/core"

/**
* List of ProxyAuthenticate headers.
 */
type ProxyAuthenticateList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewProxyAuthenticateList() *ProxyAuthenticateList {
	this := &ProxyAuthenticateList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_PROXY_AUTHENTICATE)
	return this
}
