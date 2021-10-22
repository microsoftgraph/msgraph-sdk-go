package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OpenShiftChangeRequest struct {
    ScheduleChangeRequest
    openShiftId *string;
}
func NewOpenShiftChangeRequest()(*OpenShiftChangeRequest) {
    m := &OpenShiftChangeRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
func (m *OpenShiftChangeRequest) GetOpenShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.openShiftId
    }
}
func (m *OpenShiftChangeRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleChangeRequest.GetFieldDeserializers()
    res["openShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOpenShiftId(val)
        return nil
    }
    return res
}
func (m *OpenShiftChangeRequest) IsNil()(bool) {
    return m == nil
}
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
func (m *OpenShiftChangeRequest) SetOpenShiftId(value *string)() {
    m.openShiftId = value
}
