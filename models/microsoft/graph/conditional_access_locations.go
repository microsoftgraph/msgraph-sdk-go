package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessLocations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Location IDs excluded from scope of policy.
    excludeLocations []string;
    // Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
    includeLocations []string;
}
// Instantiates a new conditionalAccessLocations and sets the default values.
func NewConditionalAccessLocations()(*ConditionalAccessLocations) {
    m := &ConditionalAccessLocations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessLocations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeLocations property value. Location IDs excluded from scope of policy.
func (m *ConditionalAccessLocations) GetExcludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeLocations
    }
}
// Gets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
func (m *ConditionalAccessLocations) GetIncludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeLocations
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessLocations) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeLocations(res)
        }
        return nil
    }
    res["includeLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeLocations(res)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessLocations) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessLocations) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeLocations", m.GetExcludeLocations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeLocations", m.GetIncludeLocations())
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
func (m *ConditionalAccessLocations) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeLocations property value. Location IDs excluded from scope of policy.
// Parameters:
//  - value : Value to set for the excludeLocations property.
func (m *ConditionalAccessLocations) SetExcludeLocations(value []string)() {
    m.excludeLocations = value
}
// Sets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
// Parameters:
//  - value : Value to set for the includeLocations property.
func (m *ConditionalAccessLocations) SetIncludeLocations(value []string)() {
    m.includeLocations = value
}
