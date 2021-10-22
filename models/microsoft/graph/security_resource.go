package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityResource struct {
    additionalData map[string]interface{};
    resource *string;
    resourceType *SecurityResourceType;
}
func NewSecurityResource()(*SecurityResource) {
    m := &SecurityResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SecurityResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SecurityResource) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *SecurityResource) GetResourceType()(*SecurityResourceType) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
func (m *SecurityResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResource(val)
        return nil
    }
    res["resourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityResourceType)
        if err != nil {
            return err
        }
        cast := val.(SecurityResourceType)
        m.SetResourceType(&cast)
        return nil
    }
    return res
}
func (m *SecurityResource) IsNil()(bool) {
    return m == nil
}
func (m *SecurityResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    if m.GetResourceType() != nil {
        cast := m.GetResourceType().String()
        err := writer.WriteStringValue("resourceType", &cast)
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
func (m *SecurityResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SecurityResource) SetResource(value *string)() {
    m.resource = value
}
func (m *SecurityResource) SetResourceType(value *SecurityResourceType)() {
    m.resourceType = value
}
