package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// optionalClaim 
type OptionalClaim struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
    additionalProperties []string;
    // If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
    essential *bool;
    // The name of the optional claim.
    name *string;
    // The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
    source *string;
}
// NewOptionalClaim instantiates a new optionalClaim and sets the default values.
func NewOptionalClaim()(*OptionalClaim) {
    m := &OptionalClaim{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaim) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalProperties gets the additionalProperties property value. Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
func (m *OptionalClaim) GetAdditionalProperties()([]string) {
    if m == nil {
        return nil
    } else {
        return m.additionalProperties
    }
}
// GetEssential gets the essential property value. If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
func (m *OptionalClaim) GetEssential()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.essential
    }
}
// GetName gets the name property value. The name of the optional claim.
func (m *OptionalClaim) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSource gets the source property value. The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
func (m *OptionalClaim) GetSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaim) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAdditionalProperties(res)
        }
        return nil
    }
    res["essential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEssential(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    return res
}
func (m *OptionalClaim) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OptionalClaim) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("additionalProperties", m.GetAdditionalProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("essential", m.GetEssential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaim) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalProperties sets the additionalProperties property value. Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
func (m *OptionalClaim) SetAdditionalProperties(value []string)() {
    m.additionalProperties = value
}
// SetEssential sets the essential property value. If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
func (m *OptionalClaim) SetEssential(value *bool)() {
    m.essential = value
}
// SetName sets the name property value. The name of the optional claim.
func (m *OptionalClaim) SetName(value *string)() {
    m.name = value
}
// SetSource sets the source property value. The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
func (m *OptionalClaim) SetSource(value *string)() {
    m.source = value
}
