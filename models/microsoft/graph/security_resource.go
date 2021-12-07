package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityResource 
type SecurityResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the resource that is related to current alert. Required.
    resource *string;
    // Represents type of security resources related to an alert. Possible values are: attacked, related.
    resourceType *SecurityResourceType;
}
// NewSecurityResource instantiates a new securityResource and sets the default values.
func NewSecurityResource()(*SecurityResource) {
    m := &SecurityResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetResource gets the resource property value. Name of the resource that is related to current alert. Required.
func (m *SecurityResource) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceType gets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
func (m *SecurityResource) GetResourceType()(*SecurityResourceType) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["resourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityResourceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SecurityResourceType)
            m.SetResourceType(&cast)
        }
        return nil
    }
    return res
}
func (m *SecurityResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResource sets the resource property value. Name of the resource that is related to current alert. Required.
func (m *SecurityResource) SetResource(value *string)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceType sets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
func (m *SecurityResource) SetResourceType(value *SecurityResourceType)() {
    if m != nil {
        m.resourceType = value
    }
}
