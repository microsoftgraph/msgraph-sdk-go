package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceAndAppManagementRoleAssignment struct {
    RoleAssignment
    // The list of ids of role member security groups. These are IDs from Azure Active Directory.
    members []string;
}
// Instantiates a new deviceAndAppManagementRoleAssignment and sets the default values.
func NewDeviceAndAppManagementRoleAssignment()(*DeviceAndAppManagementRoleAssignment) {
    m := &DeviceAndAppManagementRoleAssignment{
        RoleAssignment: *NewRoleAssignment(),
    }
    return m
}
// Gets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
func (m *DeviceAndAppManagementRoleAssignment) GetMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// The deserialization information for the current model
func (m *DeviceAndAppManagementRoleAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.RoleAssignment.GetFieldDeserializers()
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMembers(res)
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementRoleAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceAndAppManagementRoleAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.RoleAssignment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("members", m.GetMembers())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
// Parameters:
//  - value : Value to set for the members property.
func (m *DeviceAndAppManagementRoleAssignment) SetMembers(value []string)() {
    m.members = value
}
