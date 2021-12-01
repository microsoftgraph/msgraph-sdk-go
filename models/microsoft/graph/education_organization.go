package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationOrganization 
type EducationOrganization struct {
    Entity
    // Organization description.
    description *string;
    // Organization display name.
    displayName *string;
    // Source where this organization was created from. Possible values are: sis, manual.
    externalSource *EducationExternalSource;
    // The name of the external source this resources was generated from.
    externalSourceDetail *string;
}
// NewEducationOrganization instantiates a new educationOrganization and sets the default values.
func NewEducationOrganization()(*EducationOrganization) {
    m := &EducationOrganization{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Organization description.
func (m *EducationOrganization) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Organization display name.
func (m *EducationOrganization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalSource gets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
func (m *EducationOrganization) GetExternalSource()(*EducationExternalSource) {
    if m == nil {
        return nil
    } else {
        return m.externalSource
    }
}
// GetExternalSourceDetail gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationOrganization) GetExternalSourceDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalSourceDetail
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationOrganization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["externalSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationExternalSource)
            m.SetExternalSource(&cast)
        }
        return nil
    }
    res["externalSourceDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
        }
        return nil
    }
    return res
}
func (m *EducationOrganization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationOrganization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
    if m.GetExternalSource() != nil {
        cast := m.GetExternalSource().String()
        err = writer.WriteStringValue("externalSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalSourceDetail", m.GetExternalSourceDetail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Organization description.
func (m *EducationOrganization) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Organization display name.
func (m *EducationOrganization) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalSource sets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
func (m *EducationOrganization) SetExternalSource(value *EducationExternalSource)() {
    if m != nil {
        m.externalSource = value
    }
}
// SetExternalSourceDetail sets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationOrganization) SetExternalSourceDetail(value *string)() {
    if m != nil {
        m.externalSourceDetail = value
    }
}
