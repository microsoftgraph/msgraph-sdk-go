package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAppConfiguration struct {
    ManagedAppPolicy
    // A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
    customSettings []KeyValuePair;
}
// Instantiates a new managedAppConfiguration and sets the default values.
func NewManagedAppConfiguration()(*ManagedAppConfiguration) {
    m := &ManagedAppConfiguration{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// Gets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
func (m *ManagedAppConfiguration) GetCustomSettings()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.customSettings
    }
}
// The deserialization information for the current model
func (m *ManagedAppConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["customSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetCustomSettings(res)
        return nil
    }
    return res
}
func (m *ManagedAppConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedAppConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomSettings()))
        for i, v := range m.GetCustomSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
// Parameters:
//  - value : Value to set for the customSettings property.
func (m *ManagedAppConfiguration) SetCustomSettings(value []KeyValuePair)() {
    m.customSettings = value
}
