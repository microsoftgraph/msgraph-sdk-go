package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisionedPlan 
type ProvisionedPlan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For example, 'Enabled'.
    capabilityStatus *string;
    // For example, 'Success'.
    provisioningStatus *string;
    // The name of the service; for example, 'AccessControlS2S'
    service *string;
}
// NewProvisionedPlan instantiates a new provisionedPlan and sets the default values.
func NewProvisionedPlan()(*ProvisionedPlan) {
    m := &ProvisionedPlan{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionedPlan) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCapabilityStatus gets the capabilityStatus property value. For example, 'Enabled'.
func (m *ProvisionedPlan) GetCapabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilityStatus
    }
}
// GetProvisioningStatus gets the provisioningStatus property value. For example, 'Success'.
func (m *ProvisionedPlan) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
// GetService gets the service property value. The name of the service; for example, 'AccessControlS2S'
func (m *ProvisionedPlan) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionedPlan) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["capabilityStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
        }
        return nil
    }
    res["provisioningStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    return res
}
func (m *ProvisionedPlan) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionedPlan) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCapabilityStatus sets the capabilityStatus property value. For example, 'Enabled'.
func (m *ProvisionedPlan) SetCapabilityStatus(value *string)() {
    if m != nil {
        m.capabilityStatus = value
    }
}
// SetProvisioningStatus sets the provisioningStatus property value. For example, 'Success'.
func (m *ProvisionedPlan) SetProvisioningStatus(value *string)() {
    if m != nil {
        m.provisioningStatus = value
    }
}
// SetService sets the service property value. The name of the service; for example, 'AccessControlS2S'
func (m *ProvisionedPlan) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
