package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageSecurityStateable 
type MessageSecurityStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetConnectingIP()(*string)
    GetDeliveryAction()(*string)
    GetDeliveryLocation()(*string)
    GetDirectionality()(*string)
    GetInternetMessageId()(*string)
    GetMessageFingerprint()(*string)
    GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMessageSubject()(*string)
    GetNetworkMessageId()(*string)
    SetConnectingIP(value *string)()
    SetDeliveryAction(value *string)()
    SetDeliveryLocation(value *string)()
    SetDirectionality(value *string)()
    SetInternetMessageId(value *string)()
    SetMessageFingerprint(value *string)()
    SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMessageSubject(value *string)()
    SetNetworkMessageId(value *string)()
}
