package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DelegatedPermissionClassification struct {
    Entity
    // The classification value being given. Possible value: low. Does not support $filter.
    classification *PermissionClassificationType;
    // The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
    permissionId *string;
    // The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
    permissionName *string;
}
// Instantiates a new delegatedPermissionClassification and sets the default values.
func NewDelegatedPermissionClassification()(*DelegatedPermissionClassification) {
    m := &DelegatedPermissionClassification{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the classification property value. The classification value being given. Possible value: low. Does not support $filter.
func (m *DelegatedPermissionClassification) GetClassification()(*PermissionClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// Gets the permissionId property value. The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
func (m *DelegatedPermissionClassification) GetPermissionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionId
    }
}
// Gets the permissionName property value. The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
func (m *DelegatedPermissionClassification) GetPermissionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionName
    }
}
// The deserialization information for the current model
func (m *DelegatedPermissionClassification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PermissionClassificationType)
            m.SetClassification(&cast)
        }
        return nil
    }
    res["permissionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionId(val)
        }
        return nil
    }
    res["permissionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionName(val)
        }
        return nil
    }
    return res
}
func (m *DelegatedPermissionClassification) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DelegatedPermissionClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := m.GetClassification().String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionId", m.GetPermissionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionName", m.GetPermissionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the classification property value. The classification value being given. Possible value: low. Does not support $filter.
// Parameters:
//  - value : Value to set for the classification property.
func (m *DelegatedPermissionClassification) SetClassification(value *PermissionClassificationType)() {
    m.classification = value
}
// Sets the permissionId property value. The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
// Parameters:
//  - value : Value to set for the permissionId property.
func (m *DelegatedPermissionClassification) SetPermissionId(value *string)() {
    m.permissionId = value
}
// Sets the permissionName property value. The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
// Parameters:
//  - value : Value to set for the permissionName property.
func (m *DelegatedPermissionClassification) SetPermissionName(value *string)() {
    m.permissionName = value
}
