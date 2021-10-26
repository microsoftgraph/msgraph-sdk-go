package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RecordOperation struct {
    CommsOperation
    // The access token required to retrieve the recording.
    recordingAccessToken *string;
    // The location where the recording is located.
    recordingLocation *string;
}
// Instantiates a new recordOperation and sets the default values.
func NewRecordOperation()(*RecordOperation) {
    m := &RecordOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// Gets the recordingAccessToken property value. The access token required to retrieve the recording.
func (m *RecordOperation) GetRecordingAccessToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingAccessToken
    }
}
// Gets the recordingLocation property value. The location where the recording is located.
func (m *RecordOperation) GetRecordingLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingLocation
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the recordingAccessToken property value. The access token required to retrieve the recording.
// Parameters:
//  - value : Value to set for the recordingAccessToken property.
func (m *RecordOperation) SetRecordingAccessToken(value *string)() {
    m.recordingAccessToken = value
}
// Sets the recordingLocation property value. The location where the recording is located.
// Parameters:
//  - value : Value to set for the recordingLocation property.
func (m *RecordOperation) SetRecordingLocation(value *string)() {
    m.recordingLocation = value
}
