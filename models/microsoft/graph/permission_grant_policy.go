package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PermissionGrantPolicy provides operations to manage the policyRoot singleton.
type PermissionGrantPolicy struct {
    PolicyBase
    // Condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
    excludes []PermissionGrantConditionSetable;
    // Condition sets which are included in this permission grant policy. Automatically expanded on GET.
    includes []PermissionGrantConditionSetable;
}
// NewPermissionGrantPolicy instantiates a new permissionGrantPolicy and sets the default values.
func NewPermissionGrantPolicy()(*PermissionGrantPolicy) {
    m := &PermissionGrantPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreatePermissionGrantPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionGrantPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPermissionGrantPolicy(), nil
}
// GetExcludes gets the excludes property value. Condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPolicy) GetExcludes()([]PermissionGrantConditionSetable) {
    if m == nil {
        return nil
    } else {
        return m.excludes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionGrantPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantConditionSetable)
            }
            m.SetExcludes(res)
        }
        return nil
    }
    res["includes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionGrantConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionGrantConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(PermissionGrantConditionSetable)
            }
            m.SetIncludes(res)
        }
        return nil
    }
    return res
}
// GetIncludes gets the includes property value. Condition sets which are included in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPolicy) GetIncludes()([]PermissionGrantConditionSetable) {
    if m == nil {
        return nil
    } else {
        return m.includes
    }
}
func (m *PermissionGrantPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PermissionGrantPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludes()))
        for i, v := range m.GetExcludes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludes()))
        for i, v := range m.GetIncludes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcludes sets the excludes property value. Condition sets which are excluded in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPolicy) SetExcludes(value []PermissionGrantConditionSetable)() {
    if m != nil {
        m.excludes = value
    }
}
// SetIncludes sets the includes property value. Condition sets which are included in this permission grant policy. Automatically expanded on GET.
func (m *PermissionGrantPolicy) SetIncludes(value []PermissionGrantConditionSetable)() {
    if m != nil {
        m.includes = value
    }
}
