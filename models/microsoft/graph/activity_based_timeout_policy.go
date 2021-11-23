package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// activityBasedTimeoutPolicy 
type ActivityBasedTimeoutPolicy struct {
    StsPolicy
}
// NewActivityBasedTimeoutPolicy instantiates a new activityBasedTimeoutPolicy and sets the default values.
func NewActivityBasedTimeoutPolicy()(*ActivityBasedTimeoutPolicy) {
    m := &ActivityBasedTimeoutPolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivityBasedTimeoutPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
func (m *ActivityBasedTimeoutPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ActivityBasedTimeoutPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
