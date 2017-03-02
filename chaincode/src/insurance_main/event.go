package main

//==============================================================================================================================
//	ClaimSettledEvent - Defines the structure for a claim settled event.
//==============================================================================================================================
type ClaimSettledEvent struct {
	Type		string				`json:"eventType"`
	ClaimId		string				`json:"claimId"`
	PolicyId	string				`json:"policyId"`
}

//==============================================================================================================================
//	InsurerPaymentEvent - Defines the structure for an insurer payment event.
//==============================================================================================================================
type InsurerPaymentEvent struct {
	Type		string				`json:"eventType"`
	ClaimId		string				`json:"claimId"`
	PolicyId	string				`json:"policyId"`
}

//==============================================================================================================================
//	 Event type codes
//==============================================================================================================================
const EVENT_TYPE_CLAIM_SETTLED = "ClaimSettled";
const EVENT_TYPE_INSURER_PAYMENT_PAID = "InsurerPaymentPaid";
const EVENT_TYPE_INSURER_PAYMENT_ADDED = "InsurerPaymentAdded";

//=================================================================================================================================
//	 NewClaimSettledEvent	-	Constructs a new ClaimSettledEvent
//=================================================================================================================================
func NewClaimSettledEvent(claimId string, policyId string) (ClaimSettledEvent) {
	var event ClaimSettledEvent

	event.Type = EVENT_TYPE_CLAIM_SETTLED
	event.ClaimId = claimId
	event.PolicyId = policyId

	return event
}

//=================================================================================================================================
//	 NewInsurerPaymentPaidEvent	-	Constructs a new InsurerPaymentPaidEvent
//=================================================================================================================================
func NewInsurerPaymentPaidEvent(claimId string, policyId string) (InsurerPaymentEvent) {
	var event InsurerPaymentEvent

	event.Type = EVENT_TYPE_INSURER_PAYMENT_PAID
	event.ClaimId = claimId
	event.PolicyId = policyId

	return event
}

//=================================================================================================================================
//	 NewInsurerPaymentAddedEvent	-	Constructs a new InsurerPaymentAddedEvent
//=================================================================================================================================
func NewInsurerPaymentAddedEvent(claimId string, policyId string) (InsurerPaymentEvent) {
	var event InsurerPaymentEvent

	event.Type = EVENT_TYPE_INSURER_PAYMENT_ADDED
	event.ClaimId = claimId
	event.PolicyId = policyId

	return event
}
