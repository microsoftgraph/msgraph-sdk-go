package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisionedPlan struct {
    additionalData map[string]interface{};
    capabilityStatus *string;
    provisioningStatus *string;
    service *string;
}
func NewProvisionedPlan()(*ProvisionedPlan) {
    m := &ProvisionedPlan{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProvisionedPlan) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProvisionedPlan) GetCapabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilityStatus
    }
}
func (m *ProvisionedPlan) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
func (m *ProvisionedPlan) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
func (m *ProvisionedPlan) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["capabilityStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCapabilityStatus(val)
        return nil
    }
    res["provisioningStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvisioningStatus(val)
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetService(val)
        return nil
    }
    return res
}
func (m *ProvisionedPlan) IsNil()(bool) {
    return m == nil
}
func (m *ProvisionedPlan) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
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
func (m *ProvisionedPlan) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProvisionedPlan) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
func (m *ProvisionedPlan) SetProvisioningStatus(value *string)() {
    m.provisioningStatus = value
}
func (m *ProvisionedPlan) SetService(value *string)() {
    m.service = value
}
