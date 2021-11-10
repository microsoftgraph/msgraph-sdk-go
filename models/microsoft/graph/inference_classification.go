package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InferenceClassification struct {
    Entity
    // A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
    overrides []InferenceClassificationOverride;
}
// Instantiates a new inferenceClassification and sets the default values.
func NewInferenceClassification()(*InferenceClassification) {
    m := &InferenceClassification{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
func (m *InferenceClassification) GetOverrides()([]InferenceClassificationOverride) {
    if m == nil {
        return nil
    } else {
        return m.overrides
    }
}
// The deserialization information for the current model
func (m *InferenceClassification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["overrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInferenceClassificationOverride() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InferenceClassificationOverride, len(val))
            for i, v := range val {
                res[i] = *(v.(*InferenceClassificationOverride))
            }
            m.SetOverrides(res)
        }
        return nil
    }
    return res
}
func (m *InferenceClassification) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *InferenceClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOverrides()))
        for i, v := range m.GetOverrides() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("overrides", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the overrides property.
func (m *InferenceClassification) SetOverrides(value []InferenceClassificationOverride)() {
    m.overrides = value
}
