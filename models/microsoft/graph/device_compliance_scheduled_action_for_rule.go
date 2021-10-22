package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScheduledActionForRule struct {
    Entity
    ruleName *string;
    scheduledActionConfigurations []DeviceComplianceActionItem;
}
func NewDeviceComplianceScheduledActionForRule()(*DeviceComplianceScheduledActionForRule) {
    m := &DeviceComplianceScheduledActionForRule{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceComplianceScheduledActionForRule) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
func (m *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations()([]DeviceComplianceActionItem) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionConfigurations
    }
}
func (m *DeviceComplianceScheduledActionForRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleName(val)
        return nil
    }
    res["scheduledActionConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceActionItem() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceActionItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceActionItem))
        }
        m.SetScheduledActionConfigurations(res)
        return nil
    }
    return res
}
func (m *DeviceComplianceScheduledActionForRule) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScheduledActionForRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduledActionConfigurations()))
        for i, v := range m.GetScheduledActionConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceComplianceScheduledActionForRule) SetRuleName(value *string)() {
    m.ruleName = value
}
func (m *DeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(value []DeviceComplianceActionItem)() {
    m.scheduledActionConfigurations = value
}
