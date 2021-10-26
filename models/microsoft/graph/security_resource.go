package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecurityResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the resource that is related to current alert. Required.
    resource *string;
    // Represents type of security resources related to an alert. Possible values are: attacked, related.
    resourceType *SecurityResourceType;
}
// Instantiates a new securityResource and sets the default values.
func NewSecurityResource()(*SecurityResource) {
    m := &SecurityResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the resource property value. Name of the resource that is related to current alert. Required.
func (m *SecurityResource) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Gets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
func (m *SecurityResource) GetResourceType()(*SecurityResourceType) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SecurityResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the resource property value. Name of the resource that is related to current alert. Required.
// Parameters:
//  - value : Value to set for the resource property.
func (m *SecurityResource) SetResource(value *string)() {
    m.resource = value
}
// Sets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
// Parameters:
//  - value : Value to set for the resourceType property.
func (m *SecurityResource) SetResourceType(value *SecurityResourceType)() {
    m.resourceType = value
}
