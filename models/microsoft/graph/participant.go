package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Participant struct {
    Entity
    // 
    info *ParticipantInfo;
    // true if the participant is in lobby.
    isInLobby *bool;
    // true if the participant is muted (client or server muted).
    isMuted *bool;
    // The list of media streams.
    mediaStreams []MediaStream;
    // A blob of data provided by the participant in the roster.
    metadata *string;
    // Information about whether the participant has recording capability.
    recordingInfo *RecordingInfo;
}
// Instantiates a new participant and sets the default values.
func NewParticipant()(*Participant) {
    m := &Participant{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the info property value. 
func (m *Participant) GetInfo()(*ParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
// Gets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) GetIsInLobby()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInLobby
    }
}
// Gets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) GetIsMuted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMuted
    }
}
// Gets the mediaStreams property value. The list of media streams.
func (m *Participant) GetMediaStreams()([]MediaStream) {
    if m == nil {
        return nil
    } else {
        return m.mediaStreams
    }
}
// Gets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) GetMetadata()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the recordingInfo property value. Information about whether the participant has recording capability.
func (m *Participant) GetRecordingInfo()(*RecordingInfo) {
    if m == nil {
        return nil
    } else {
        return m.recordingInfo
    }
}
// The deserialization information for the current model
func (m *Participant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["info"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(*ParticipantInfo))
        }
        return nil
    }
    res["isInLobby"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInLobby(val)
        }
        return nil
    }
    res["isMuted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMuted(val)
        }
        return nil
    }
    res["mediaStreams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaStream() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaStream, len(val))
            for i, v := range val {
                res[i] = *(v.(*MediaStream))
            }
            m.SetMediaStreams(res)
        }
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["recordingInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecordingInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingInfo(val.(*RecordingInfo))
        }
        return nil
    }
    return res
}
func (m *Participant) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Participant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInLobby", m.GetIsInLobby())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMuted", m.GetIsMuted())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMediaStreams()))
        for i, v := range m.GetMediaStreams() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mediaStreams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recordingInfo", m.GetRecordingInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the info property value. 
// Parameters:
//  - value : Value to set for the info property.
func (m *Participant) SetInfo(value *ParticipantInfo)() {
    m.info = value
}
// Sets the isInLobby property value. true if the participant is in lobby.
// Parameters:
//  - value : Value to set for the isInLobby property.
func (m *Participant) SetIsInLobby(value *bool)() {
    m.isInLobby = value
}
// Sets the isMuted property value. true if the participant is muted (client or server muted).
// Parameters:
//  - value : Value to set for the isMuted property.
func (m *Participant) SetIsMuted(value *bool)() {
    m.isMuted = value
}
// Sets the mediaStreams property value. The list of media streams.
// Parameters:
//  - value : Value to set for the mediaStreams property.
func (m *Participant) SetMediaStreams(value []MediaStream)() {
    m.mediaStreams = value
}
// Sets the metadata property value. A blob of data provided by the participant in the roster.
// Parameters:
//  - value : Value to set for the metadata property.
func (m *Participant) SetMetadata(value *string)() {
    m.metadata = value
}
// Sets the recordingInfo property value. Information about whether the participant has recording capability.
// Parameters:
//  - value : Value to set for the recordingInfo property.
func (m *Participant) SetRecordingInfo(value *RecordingInfo)() {
    m.recordingInfo = value
}
