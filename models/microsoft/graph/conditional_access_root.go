package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessRoot 
type ConditionalAccessRoot struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified named locations.
    namedLocations []NamedLocation;
    // Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
    policies []ConditionalAccessPolicy;
}
// NewConditionalAccessRoot instantiates a new conditionalAccessRoot and sets the default values.
func NewConditionalAccessRoot()(*ConditionalAccessRoot) {
    m := &ConditionalAccessRoot{
        Entity: *NewEntity(),
    }
    return m
}
// GetNamedLocations gets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) GetNamedLocations()([]NamedLocation) {
    if m == nil {
        return nil
    } else {
        return m.namedLocations
    }
}
// GetPolicies gets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
func (m *ConditionalAccessRoot) GetPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["namedLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNamedLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NamedLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*NamedLocation))
            }
            m.SetNamedLocations(res)
        }
        return nil
    }
    res["policies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessPolicy))
            }
            m.SetPolicies(res)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNamedLocations()))
        for i, v := range m.GetNamedLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("namedLocations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNamedLocations sets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) SetNamedLocations(value []NamedLocation)() {
    if m != nil {
        m.namedLocations = value
    }
}
// SetPolicies sets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
func (m *ConditionalAccessRoot) SetPolicies(value []ConditionalAccessPolicy)() {
    if m != nil {
        m.policies = value
    }
}
