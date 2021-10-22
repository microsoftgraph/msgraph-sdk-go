package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RecordOperation struct {
    CommsOperation
    recordingAccessToken *string;
    recordingLocation *string;
}
func NewRecordOperation()(*RecordOperation) {
    m := &RecordOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
func (m *RecordOperation) GetRecordingAccessToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingAccessToken
    }
}
func (m *RecordOperation) GetRecordingLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingLocation
    }
}
func (m *RecordOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["recordingAccessToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecordingAccessToken(val)
        return nil
    }
    res["recordingLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecordingLocation(val)
        return nil
    }
    return res
}
func (m *RecordOperation) IsNil()(bool) {
    return m == nil
}
func (m *RecordOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("recordingAccessToken", m.GetRecordingAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordingLocation", m.GetRecordingLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *RecordOperation) SetRecordingAccessToken(value *string)() {
    m.recordingAccessToken = value
}
func (m *RecordOperation) SetRecordingLocation(value *string)() {
    m.recordingLocation = value
}
