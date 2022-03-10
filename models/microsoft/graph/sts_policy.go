package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StsPolicy provides operations to manage the collection of application entities.
type StsPolicy struct {
    PolicyBase
    // 
    appliesTo []DirectoryObjectable;
    // A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
    definition []string;
    // If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
    isOrganizationDefault *bool;
}
// NewStsPolicy instantiates a new stsPolicy and sets the default values.
func NewStsPolicy()(*StsPolicy) {
    m := &StsPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateStsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStsPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStsPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. 
func (m *StsPolicy) GetAppliesTo()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetDefinition gets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
func (m *StsPolicy) GetDefinition()([]string) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
        }
        return nil
    }
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefinition(res)
        }
        return nil
    }
    res["isOrganizationDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationDefault(val)
        }
        return nil
    }
    return res
}
// GetIsOrganizationDefault gets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
func (m *StsPolicy) GetIsOrganizationDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizationDefault
    }
}
func (m *StsPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StsPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefinition() != nil {
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
// SetAppliesTo sets the appliesTo property value. 
func (m *StsPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetDefinition sets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
func (m *StsPolicy) SetDefinition(value []string)() {
    if m != nil {
        m.definition = value
    }
}
// SetIsOrganizationDefault sets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
func (m *StsPolicy) SetIsOrganizationDefault(value *bool)() {
    if m != nil {
        m.isOrganizationDefault = value
    }
}
