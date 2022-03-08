package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecordingInfo provides operations to manage the cloudCommunications singleton.
type RecordingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identities of the recording initiator.
    initiator IdentitySetable;
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
// CreateRecordingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecordingInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecordingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["initiator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["recordingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecordingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingStatus(val.(*RecordingStatus))
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) GetInitiator()(IdentitySetable) {
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
        cast := (*m.GetRecordingStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetInitiator sets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) SetInitiator(value IdentitySetable)() {
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
