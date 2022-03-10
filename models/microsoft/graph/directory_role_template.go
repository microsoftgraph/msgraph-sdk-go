package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectoryRoleTemplate provides operations to manage the collection of directoryRoleTemplate entities.
type DirectoryRoleTemplate struct {
    DirectoryObject
    // The description to set for the directory role. Read-only.
    description *string;
    // The display name to set for the directory role. Read-only.
    displayName *string;
}
// NewDirectoryRoleTemplate instantiates a new directoryRoleTemplate and sets the default values.
func NewDirectoryRoleTemplate()(*DirectoryRoleTemplate) {
    m := &DirectoryRoleTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// CreateDirectoryRoleTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryRoleTemplateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDirectoryRoleTemplate(), nil
}
// GetDescription gets the description property value. The description to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDescription sets the description property value. The description to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
