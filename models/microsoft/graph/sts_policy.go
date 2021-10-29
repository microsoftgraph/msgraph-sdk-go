package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type StsPolicy struct {
    PolicyBase
    // 
    appliesTo []DirectoryObject;
    // A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
    definition []string;
    // If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
    isOrganizationDefault *bool;
}
// Instantiates a new stsPolicy and sets the default values.
func NewStsPolicy()(*StsPolicy) {
    m := &StsPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// Gets the appliesTo property value. 
func (m *StsPolicy) GetAppliesTo()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
func (m *StsPolicy) GetDefinition()([]string) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// Gets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
func (m *StsPolicy) GetIsOrganizationDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizationDefault
    }
}
// The deserialization information for the current model
func (m *StsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetAppliesTo(res)
        return nil
    }
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDefinition(res)
        return nil
    }
    res["isOrganizationDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOrganizationDefault(val)
        return nil
    }
    return res
}
func (m *StsPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *StsPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationDefault", m.GetIsOrganizationDefault())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appliesTo property value. 
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *StsPolicy) SetAppliesTo(value []DirectoryObject)() {
    m.appliesTo = value
}
// Sets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
// Parameters:
//  - value : Value to set for the definition property.
func (m *StsPolicy) SetDefinition(value []string)() {
    m.definition = value
}
// Sets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
// Parameters:
//  - value : Value to set for the isOrganizationDefault property.
func (m *StsPolicy) SetIsOrganizationDefault(value *bool)() {
    m.isOrganizationDefault = value
}
