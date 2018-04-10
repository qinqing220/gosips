/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ClientTransaction.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package sip

import (
	"github.com/use-go/gosips/sip/message"
)

/**
 * A client transaction is used by a User Agent Client application to send
 * Request messages to a User Agent Server application.
 * The client transaction is also used to match Responses from the User Agent
 * Server to fire Response events to the SipListener for a specific client
 * transaction. This interfaces enables an application to send a
 * {@link javax.sip.message.Request}'s statefully. A new client transaction
 * is generated by the application calling the
 * {@link SipProvider#getNewClientTransaction(Request)} method.
 * <p>
 * A client transaction of the transaction layer is represented by a finite
 * state machine that is constructed to process a particular request under
 * the covers of a stateful SipProvider. The transaction layer handles
 * application-layer retransmissions, matching of responses to requests, and
 * application-layer timeouts. Any task that a User Agent Client
 * accomplishes takes place using a series of transactions.
 * <p>
 * The client transaction must be unique within the underlying
 * implementation. A common way to create this value is to compute a
 * cryptographic hash of the To tag, From tag, Call-ID header field, the
 * Request-URI of the request received (before translation), the topmost Via
 * header, and the sequence number from the CSeq header field, in addition to
 * any Proxy-Require and Proxy-Authorization header fields that may be present.
 * The algorithm used to compute the hash is implementation-dependent.
 * <p>
 * For the detailed client transaction state machines refer to Chapter
 * 17 of <a href="http://www.ietf.org/rfc/rfc3261.txt">RFC 3261</a>, the
 * allowable transitions are summarized below:
 * <p>
 * <b>Invite Transaction:</b><br>
 * Calling --> Proceeding --> Completed --> Terminated
 * <p>
 * <b>Non-Invite Transaction:</b><br>
 * Trying --> Proceeding --> Completed --> Terminated
 *
 * @author  Rain Liu
 */
type ClientTransaction interface {
	Transaction

	/**
	 * Sends the Request which created this ClientTransaction. When an
	 * application wishes to send a Request message, it creates a Request from
	 * the {@link javax.sip.message.MessageFactory} and then creates a new
	 * ClientTransaction from
	 * {@link SipProvider#getNewClientTransaction(Request)}. Calling this method
	 * on the ClientTransaction sends the Request onto the network. The Request
	 * message gets sent via the ListeningPoint information of the SipProvider
	 * that is associated to this ClientTransaction.
	 * <p>
	 * This method assumes that the Request is sent out of Dialog. It uses
	 * the Router to determine the next hop. If the Router returns a empty
	 * iterator, and a Dialog is associated with the outgoing request of the
	 * Transaction then the Dialog route set is used to send the outgoing
	 * request.
	 * <p>
	 * This method implies that the application is functioning as either a UAC
	 * or a stateful proxy, hence the underlying implementation acts statefully.
	 *
	 * @throws SipException if the SipProvider cannot send the Request for any
	 * reason.
	 * @see Request
	 */
	SendRequest() (SipException error)

	/**
	 * Creates a new Cancel message from the Request associated with this client
	 * transaction. The CANCEL request, is used to cancel the previous request
	 * sent by this client transaction. Specifically, it asks the UAS to cease
	 * processing the request and to generate an error response to that request.
	 * A CANCEL request constitutes its own transaction, but also references
	 * the transaction to be cancelled. CANCEL has no effect on a request to
	 * which a UAS has already given a final response.
	 * <p>
	 * Note that both the transaction corresponding to the original request and
	 * the CANCEL transaction will complete independently.  However, a UAC
	 * canceling a request cannot rely on receiving a 487 (Request Terminated)
	 * response for the original request, as an RFC 2543 compliant UAS will
	 * not generate such a response. Therefore if there is no final response for
	 * the original request the application will receieve a TimeoutEvent with
	 * {@link javax.sip.Timeout#TRANSACTION} and the client should then consider the
	 * original transaction cancelled.
	 * <ul>
	 * <li> UAC - A UAC should not send a CANCEL request to any request explicitly
	 * supported by this specification other than INVITE request. The reason
	 * being requests other than INVITE are responded to immediately and sending
	 * a CANCEL for a non-INVITE request would always create a race condition.
	 * CANCELs are useful as a UAC can not send a BYE request on a dialog
	 * until receipt of 2xx final response to the INVITE request. The CANCEL
	 * attempts to force a non-2xx response to the INVITE, therefore if a UAC
	 * wishes to give up on its call attempt entirely it can send a CANCEL.
	 * <li>Stateful proxies - A stateful proxy may generate CANCEL requests
	 * for:
	 * <ul>
	 * <li>INVITE Requests - A CANCEL can be sent on pending INVITE client
	 * transactions based on the period specified in the INVITE's Expires
	 * header field elapsing.  However, this is generally unnecessary since
	 * the endpoints involved will take care of signaling the end of the
	 * transaction.
	 * <li> Other Requests - A CANCEL can be sent on any other request the proxy
	 * has generated at any time, subject to receiving a provisional response
	 * to that request.
	 * </ul>
	 * </ul>
	 *
	 * @return the new cancel Request specific to the Request of this client
	 * transaction.
	 * @throws SipException if this method is called to cancel a request that
	 * can't be cancelled i.e. ACK.
	 */
	CreateCancel() (r message.Request, SipException error)

	/**
	 * Creates a new Ack message from the Request associated with this client
	 * transaction. This ACK can be used to acknowledge the response to the
	 * request sent by this transaction. It is recommended that a
	 * ClientTransaction be created to send the ACK.
	 *
	 * @return the new ACK Request specific to the Request of this client
	 * transaction.
	 * @throws SipException if this method is called before a final response
	 * is received for the transaction.
	 */
	CreateAck() (r message.Request, SipException error)
}
