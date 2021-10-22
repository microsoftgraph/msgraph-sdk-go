package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OrganizationalBranding struct {
    OrganizationalBrandingProperties
    localizations []OrganizationalBrandingLocalization;
}
func NewOrganizationalBranding()(*OrganizationalBranding) {
    m := &OrganizationalBranding{
        OrganizationalBrandingProperties: *NewOrganizationalBrandingProperties(),
    }
    return m
}
func (m *OrganizationalBranding) GetLocalizations()([]OrganizationalBrandingLocalization) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
func (m *OrganizationalBranding) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OrganizationalBrandingProperties.GetFieldDeserializers()
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOrganizationalBrandingLocalization() })
        if err != nil {
            return err
        }
        res := make([]OrganizationalBrandingLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*OrganizationalBrandingLocalization))
        }
        m.SetLocalizations(res)
        return nil
    }
    return res
}
func (m *OrganizationalBranding) IsNil()(bool) {
    return m == nil
}
func (m *OrganizationalBranding) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OrganizationalBrandingProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OrganizationalBranding) SetLocalizations(value []OrganizationalBrandingLocalization)() {
    m.localizations = value
}
