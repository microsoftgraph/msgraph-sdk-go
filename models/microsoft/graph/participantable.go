package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Participantable 
type Participantable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetInfo()(ParticipantInfoable)
    GetIsInLobby()(*bool)
    GetIsMuted()(*bool)
    GetMediaStreams()([]MediaStreamable)
    GetMetadata()(*string)
    GetRecordingInfo()(RecordingInfoable)
    SetInfo(value ParticipantInfoable)()
    SetIsInLobby(value *bool)()
    SetIsMuted(value *bool)()
    SetMediaStreams(value []MediaStreamable)()
    SetMetadata(value *string)()
    SetRecordingInfo(value RecordingInfoable)()
}
