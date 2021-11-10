package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RecordingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identities of the recording initiator.
    initiator *IdentitySet;
    // Possible values are: unknown, notRecording, recording, or failed.
    recordingStatus *RecordingStatus;
}
// Instantiates a new recordingInfo and sets the default values.
func NewRecordingInfo()(*RecordingInfo) {
    m := &RecordingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) GetInitiator()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.initiator
    }
}
// Gets the recordingStatus property value. Possible values are: unknown, notRecording, recording, or failed.
func (m *RecordingInfo) GetRecordingStatus()(*RecordingStatus) {
    if m == nil {
        return nil
    } else {
        return m.recordingStatus
    }
}
// The deserialization information for the current model
func (m *RecordingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["initiator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(*IdentitySet))
        }
        return nil
    }
    res["recordingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecordingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RecordingStatus)
            m.SetRecordingStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *RecordingInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RecordingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    if m.GetRecordingStatus() != nil {
        cast := m.GetRecordingStatus().String()
        err := writer.WriteStringValue("recordingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RecordingInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the initiator property value. The identities of the recording initiator.
// Parameters:
//  - value : Value to set for the initiator property.
func (m *RecordingInfo) SetInitiator(value *IdentitySet)() {
    m.initiator = value
}
// Sets the recordingStatus property value. Possible values are: unknown, notRecording, recording, or failed.
// Parameters:
//  - value : Value to set for the recordingStatus property.
func (m *RecordingInfo) SetRecordingStatus(value *RecordingStatus)() {
    m.recordingStatus = value
}
