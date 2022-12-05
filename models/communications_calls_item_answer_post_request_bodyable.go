package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemAnswerPostRequestBodyable 
type CommunicationsCallsItemAnswerPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedModalities()([]Modality)
    GetCallbackUri()(*string)
    GetCallOptions()(IncomingCallOptionsable)
    GetMediaConfig()(MediaConfigable)
    GetParticipantCapacity()(*int32)
    SetAcceptedModalities(value []Modality)()
    SetCallbackUri(value *string)()
    SetCallOptions(value IncomingCallOptionsable)()
    SetMediaConfig(value MediaConfigable)()
    SetParticipantCapacity(value *int32)()
}
