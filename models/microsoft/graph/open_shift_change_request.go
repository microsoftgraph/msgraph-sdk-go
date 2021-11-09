package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OpenShiftChangeRequest struct {
    ScheduleChangeRequest
    // ID for the open shift.
    openShiftId *string;
}
// Instantiates a new openShiftChangeRequest and sets the default values.
func NewOpenShiftChangeRequest()(*OpenShiftChangeRequest) {
    m := &OpenShiftChangeRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
// Gets the openShiftId property value. ID for the open shift.
func (m *OpenShiftChangeRequest) GetOpenShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.openShiftId
    }
}
// The deserialization information for the current model
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
func (m *OpenShiftChangeRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the openShiftId property value. ID for the open shift.
// Parameters:
//  - value : Value to set for the openShiftId property.
func (m *OpenShiftChangeRequest) SetOpenShiftId(value *string)() {
    m.openShiftId = value
}
