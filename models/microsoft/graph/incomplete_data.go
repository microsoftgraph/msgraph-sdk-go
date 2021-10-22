package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IncompleteData struct {
    additionalData map[string]interface{};
    missingDataBeforeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    wasThrottled *bool;
}
func NewIncompleteData()(*IncompleteData) {
    m := &IncompleteData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IncompleteData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IncompleteData) GetMissingDataBeforeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.missingDataBeforeDateTime
    }
}
func (m *IncompleteData) GetWasThrottled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wasThrottled
    }
}
func (m *IncompleteData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["missingDataBeforeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMissingDataBeforeDateTime(val)
        return nil
    }
    res["wasThrottled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWasThrottled(val)
        return nil
    }
    return res
}
func (m *IncompleteData) IsNil()(bool) {
    return m == nil
}
func (m *IncompleteData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("missingDataBeforeDateTime", m.GetMissingDataBeforeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasThrottled", m.GetWasThrottled())
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
func (m *IncompleteData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IncompleteData) SetMissingDataBeforeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.missingDataBeforeDateTime = value
}
func (m *IncompleteData) SetWasThrottled(value *bool)() {
    m.wasThrottled = value
}
