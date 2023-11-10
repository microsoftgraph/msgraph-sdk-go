package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallRecording 
type CallRecording struct {
    Entity
}
// NewCallRecording instantiates a new callRecording and sets the default values.
func NewCallRecording()(*CallRecording) {
    m := &CallRecording{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallRecordingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRecordingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRecording(), nil
}
// GetContent gets the content property value. The content property
func (m *CallRecording) GetContent()([]byte) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *CallRecording) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRecording) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["meetingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingId(val)
        }
        return nil
    }
    res["meetingOrganizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingOrganizer(val.(IdentitySetable))
        }
        return nil
    }
    res["recordingContentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingContentUrl(val)
        }
        return nil
    }
    return res
}
// GetMeetingId gets the meetingId property value. The meetingId property
func (m *CallRecording) GetMeetingId()(*string) {
    val, err := m.GetBackingStore().Get("meetingId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMeetingOrganizer gets the meetingOrganizer property value. The meetingOrganizer property
func (m *CallRecording) GetMeetingOrganizer()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("meetingOrganizer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetRecordingContentUrl gets the recordingContentUrl property value. The recordingContentUrl property
func (m *CallRecording) GetRecordingContentUrl()(*string) {
    val, err := m.GetBackingStore().Get("recordingContentUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallRecording) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meetingId", m.GetMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingOrganizer", m.GetMeetingOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordingContentUrl", m.GetRecordingContentUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *CallRecording) SetContent(value []byte)() {
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *CallRecording) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingId sets the meetingId property value. The meetingId property
func (m *CallRecording) SetMeetingId(value *string)() {
    err := m.GetBackingStore().Set("meetingId", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingOrganizer sets the meetingOrganizer property value. The meetingOrganizer property
func (m *CallRecording) SetMeetingOrganizer(value IdentitySetable)() {
    err := m.GetBackingStore().Set("meetingOrganizer", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordingContentUrl sets the recordingContentUrl property value. The recordingContentUrl property
func (m *CallRecording) SetRecordingContentUrl(value *string)() {
    err := m.GetBackingStore().Set("recordingContentUrl", value)
    if err != nil {
        panic(err)
    }
}
// CallRecordingable 
type CallRecordingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMeetingId()(*string)
    GetMeetingOrganizer()(IdentitySetable)
    GetRecordingContentUrl()(*string)
    SetContent(value []byte)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMeetingId(value *string)()
    SetMeetingOrganizer(value IdentitySetable)()
    SetRecordingContentUrl(value *string)()
}
