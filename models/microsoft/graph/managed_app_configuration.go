package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppConfiguration provides operations to manage the deviceAppManagement singleton.
type ManagedAppConfiguration struct {
    ManagedAppPolicy
    // A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
    customSettings []KeyValuePairable;
}
// NewManagedAppConfiguration instantiates a new managedAppConfiguration and sets the default values.
func NewManagedAppConfiguration()(*ManagedAppConfiguration) {
    m := &ManagedAppConfiguration{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// CreateManagedAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedAppConfiguration(), nil
}
// GetCustomSettings gets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
func (m *ManagedAppConfiguration) GetCustomSettings()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.customSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["customSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetCustomSettings(res)
        }
        return nil
    }
    return res
}
func (m *ManagedAppConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedAppConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomSettings()))
        for i, v := range m.GetCustomSettings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomSettings sets the customSettings property value. A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service
func (m *ManagedAppConfiguration) SetCustomSettings(value []KeyValuePairable)() {
    if m != nil {
        m.customSettings = value
    }
}
