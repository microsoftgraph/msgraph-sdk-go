package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Participant struct {
    Entity
    info *ParticipantInfo;
    isInLobby *bool;
    isMuted *bool;
    mediaStreams []MediaStream;
    recordingInfo *RecordingInfo;
}
func NewParticipant()(*Participant) {
    m := &Participant{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Participant) GetInfo()(*ParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.info
    }
}
func (m *Participant) GetIsInLobby()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInLobby
    }
}
func (m *Participant) GetIsMuted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMuted
    }
}
func (m *Participant) GetMediaStreams()([]MediaStream) {
    if m == nil {
        return nil
    } else {
        return m.mediaStreams
    }
}
func (m *Participant) GetRecordingInfo()(*RecordingInfo) {
    if m == nil {
        return nil
    } else {
        return m.recordingInfo
    }
}
func (m *Participant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["info"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipantInfo() })
        if err != nil {
            return err
        }
        m.SetInfo(val.(*ParticipantInfo))
        return nil
    }
    res["isInLobby"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInLobby(val)
        return nil
    }
    res["isMuted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMuted(val)
        return nil
    }
    res["mediaStreams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaStream() })
        if err != nil {
            return err
        }
        res := make([]MediaStream, len(val))
        for i, v := range val {
            res[i] = *(v.(*MediaStream))
        }
        m.SetMediaStreams(res)
        return nil
    }
    res["recordingInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecordingInfo() })
        if err != nil {
            return err
        }
        m.SetRecordingInfo(val.(*RecordingInfo))
        return nil
    }
    return res
}
func (m *Participant) IsNil()(bool) {
    return m == nil
}
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
        err = writer.WriteObjectValue("recordingInfo", m.GetRecordingInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Participant) SetInfo(value *ParticipantInfo)() {
    m.info = value
}
func (m *Participant) SetIsInLobby(value *bool)() {
    m.isInLobby = value
}
func (m *Participant) SetIsMuted(value *bool)() {
    m.isMuted = value
}
func (m *Participant) SetMediaStreams(value []MediaStream)() {
    m.mediaStreams = value
}
func (m *Participant) SetRecordingInfo(value *RecordingInfo)() {
    m.recordingInfo = value
}
