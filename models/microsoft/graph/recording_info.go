package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecordingInfo 
type RecordingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identities of the recording initiator.
    initiator *IdentitySet;
    // Possible values are: unknown, notRecording, recording, or failed.
    recordingStatus *RecordingStatus;
}
// NewRecordingInfo instantiates a new recordingInfo and sets the default values.
func NewRecordingInfo()(*RecordingInfo) {
    m := &RecordingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetInitiator gets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) GetInitiator()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.initiator
    }
}
// GetRecordingStatus gets the recordingStatus property value. Possible values are: unknown, notRecording, recording, or failed.
func (m *RecordingInfo) GetRecordingStatus()(*RecordingStatus) {
    if m == nil {
        return nil
    } else {
        return m.recordingStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetInitiator sets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) SetInitiator(value *IdentitySet)() {
    if m != nil {
        m.initiator = value
    }
}
// SetRecordingStatus sets the recordingStatus property value. Possible values are: unknown, notRecording, recording, or failed.
func (m *RecordingInfo) SetRecordingStatus(value *RecordingStatus)() {
    if m != nil {
        m.recordingStatus = value
    }
}
