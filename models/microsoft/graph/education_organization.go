package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new educationOrganization and sets the default values.
func NewEducationOrganization()(*EducationOrganization) {
    m := &EducationOrganization{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Organization description.
func (m *EducationOrganization) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Organization display name.
func (m *EducationOrganization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
func (m *EducationOrganization) GetExternalSource()(*EducationExternalSource) {
    if m == nil {
        return nil
    } else {
        return m.externalSource
    }
}
// Gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationOrganization) GetExternalSourceDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalSourceDetail
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the description property value. Organization description.
// Parameters:
//  - value : Value to set for the description property.
func (m *EducationOrganization) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Organization display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationOrganization) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
// Parameters:
//  - value : Value to set for the externalSource property.
func (m *EducationOrganization) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// Sets the externalSourceDetail property value. The name of the external source this resources was generated from.
// Parameters:
//  - value : Value to set for the externalSourceDetail property.
func (m *EducationOrganization) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
