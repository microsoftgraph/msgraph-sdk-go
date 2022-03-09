package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OpenShiftChangeRequest provides operations to manage the educationRoot singleton.
type OpenShiftChangeRequest struct {
    ScheduleChangeRequest
    // ID for the open shift.
    openShiftId *string;
}
// NewOpenShiftChangeRequest instantiates a new openShiftChangeRequest and sets the default values.
func NewOpenShiftChangeRequest()(*OpenShiftChangeRequest) {
    m := &OpenShiftChangeRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
// CreateOpenShiftChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenShiftChangeRequestFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOpenShiftChangeRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShiftChangeRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleChangeRequest.GetFieldDeserializers()
    res["openShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenShiftId(val)
        }
        return nil
    }
    return res
}
// GetOpenShiftId gets the openShiftId property value. ID for the open shift.
func (m *OpenShiftChangeRequest) GetOpenShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.openShiftId
    }
}
func (m *OpenShiftChangeRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OpenShiftChangeRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ScheduleChangeRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("openShiftId", m.GetOpenShiftId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOpenShiftId sets the openShiftId property value. ID for the open shift.
func (m *OpenShiftChangeRequest) SetOpenShiftId(value *string)() {
    if m != nil {
        m.openShiftId = value
    }
}
