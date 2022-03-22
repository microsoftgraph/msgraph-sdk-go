package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessLocations 
type ConditionalAccessLocations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Location IDs excluded from scope of policy.
    excludeLocations []string;
    // Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
    includeLocations []string;
}
// NewConditionalAccessLocations instantiates a new conditionalAccessLocations and sets the default values.
func NewConditionalAccessLocations()(*ConditionalAccessLocations) {
    m := &ConditionalAccessLocations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessLocationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessLocationsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConditionalAccessLocations(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessLocations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeLocations gets the excludeLocations property value. Location IDs excluded from scope of policy.
func (m *ConditionalAccessLocations) GetExcludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeLocations
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetIncludeLocations gets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
func (m *ConditionalAccessLocations) GetIncludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeLocations
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessLocations) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExcludeLocations() != nil {
        err := writer.WriteCollectionOfStringValues("excludeLocations", m.GetExcludeLocations())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeLocations() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessLocations) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeLocations sets the excludeLocations property value. Location IDs excluded from scope of policy.
func (m *ConditionalAccessLocations) SetExcludeLocations(value []string)() {
    if m != nil {
        m.excludeLocations = value
    }
}
// SetIncludeLocations sets the includeLocations property value. Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted.
func (m *ConditionalAccessLocations) SetIncludeLocations(value []string)() {
    if m != nil {
        m.includeLocations = value
    }
}
