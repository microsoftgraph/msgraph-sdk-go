package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecordOperation provides operations to call the recordResponse method.
type RecordOperation struct {
    CommsOperation
    // The access token required to retrieve the recording.
    recordingAccessToken *string;
    // The location where the recording is located.
    recordingLocation *string;
}
// NewRecordOperation instantiates a new recordOperation and sets the default values.
func NewRecordOperation()(*RecordOperation) {
    m := &RecordOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateRecordOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecordOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecordOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["recordingAccessToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingAccessToken(val)
        }
        return nil
    }
    res["recordingLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingLocation(val)
        }
        return nil
    }
    return res
}
// GetRecordingAccessToken gets the recordingAccessToken property value. The access token required to retrieve the recording.
func (m *RecordOperation) GetRecordingAccessToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingAccessToken
    }
}
// GetRecordingLocation gets the recordingLocation property value. The location where the recording is located.
func (m *RecordOperation) GetRecordingLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordingLocation
    }
}
func (m *RecordOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetRecordingAccessToken sets the recordingAccessToken property value. The access token required to retrieve the recording.
func (m *RecordOperation) SetRecordingAccessToken(value *string)() {
    if m != nil {
        m.recordingAccessToken = value
    }
}
// SetRecordingLocation sets the recordingLocation property value. The location where the recording is located.
func (m *RecordOperation) SetRecordingLocation(value *string)() {
    if m != nil {
        m.recordingLocation = value
    }
}
