package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupSetting provides operations to manage the collection of group entities.
type GroupSetting struct {
    Entity
    // Display name of this group of settings, which comes from the associated template.
    displayName *string;
    // Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
    templateId *string;
    // Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
    values []SettingValueable;
}
// NewGroupSetting instantiates a new groupSetting and sets the default values.
func NewGroupSetting()(*GroupSetting) {
    m := &GroupSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupSettingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGroupSetting(), nil
}
// GetDisplayName gets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(SettingValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
func (m *GroupSetting) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetValues gets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
func (m *GroupSetting) GetValues()([]SettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *GroupSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTemplateId sets the templateId property value. Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
func (m *GroupSetting) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetValues sets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
func (m *GroupSetting) SetValues(value []SettingValueable)() {
    if m != nil {
        m.values = value
    }
}
