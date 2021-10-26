package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessRoot struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified named locations.
    namedLocations []NamedLocation;
    // Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
    policies []ConditionalAccessPolicy;
}
// Instantiates a new conditionalAccessRoot and sets the default values.
func NewConditionalAccessRoot()(*ConditionalAccessRoot) {
    m := &ConditionalAccessRoot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
func (m *ConditionalAccessRoot) GetNamedLocations()([]NamedLocation) {
    if m == nil {
        return nil
    } else {
        return m.namedLocations
    }
}
// Gets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
func (m *ConditionalAccessRoot) GetPolicies()([]ConditionalAccessPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policies
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["namedLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNamedLocation() })
        if err != nil {
            return err
        }
        res := make([]NamedLocation, len(val))
        for i, v := range val {
            res[i] = *(v.(*NamedLocation))
        }
        m.SetNamedLocations(res)
        return nil
    }
    res["policies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicy() })
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessPolicy))
        }
        m.SetPolicies(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the namedLocations property value. Read-only. Nullable. Returns a collection of the specified named locations.
// Parameters:
//  - value : Value to set for the namedLocations property.
func (m *ConditionalAccessRoot) SetNamedLocations(value []NamedLocation)() {
    m.namedLocations = value
}
// Sets the policies property value. Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies.
// Parameters:
//  - value : Value to set for the policies property.
func (m *ConditionalAccessRoot) SetPolicies(value []ConditionalAccessPolicy)() {
    m.policies = value
}
