package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrganizationalBranding provides operations to manage the organizationalBranding singleton.
type OrganizationalBranding struct {
    OrganizationalBrandingProperties
    // Add different branding based on a locale.
    localizations []OrganizationalBrandingLocalizationable;
}
// NewOrganizationalBranding instantiates a new organizationalBranding and sets the default values.
func NewOrganizationalBranding()(*OrganizationalBranding) {
    m := &OrganizationalBranding{
        OrganizationalBrandingProperties: *NewOrganizationalBrandingProperties(),
    }
    return m
}
// CreateOrganizationalBrandingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalBrandingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOrganizationalBranding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBranding) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OrganizationalBrandingProperties.GetFieldDeserializers()
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalBrandingLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalBrandingLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalBrandingLocalizationable)
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    return res
}
// GetLocalizations gets the localizations property value. Add different branding based on a locale.
func (m *OrganizationalBranding) GetLocalizations()([]OrganizationalBrandingLocalizationable) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
// Serialize serializes information the current object
func (m *OrganizationalBranding) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OrganizationalBrandingProperties.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLocalizations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocalizations sets the localizations property value. Add different branding based on a locale.
func (m *OrganizationalBranding) SetLocalizations(value []OrganizationalBrandingLocalizationable)() {
    if m != nil {
        m.localizations = value
    }
}
