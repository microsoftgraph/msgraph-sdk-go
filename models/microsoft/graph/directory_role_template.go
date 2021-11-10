package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DirectoryRoleTemplate struct {
    DirectoryObject
    // The description to set for the directory role. Read-only.
    description *string;
    // The display name to set for the directory role. Read-only.
    displayName *string;
}
// Instantiates a new directoryRoleTemplate and sets the default values.
func NewDirectoryRoleTemplate()(*DirectoryRoleTemplate) {
    m := &DirectoryRoleTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the description property value. The description to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *DirectoryRoleTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *DirectoryRoleTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DirectoryRoleTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. The description to set for the directory role. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *DirectoryRoleTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name to set for the directory role. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DirectoryRoleTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
