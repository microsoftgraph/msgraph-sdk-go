package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OutlookUser struct {
    Entity
    // A list of categories defined for the user.
    masterCategories []OutlookCategory;
}
// Instantiates a new outlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) GetMasterCategories()([]OutlookCategory) {
    if m == nil {
        return nil
    } else {
        return m.masterCategories
    }
}
// The deserialization information for the current model
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookCategory() })
        if err != nil {
            return err
        }
        res := make([]OutlookCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookCategory))
        }
        m.SetMasterCategories(res)
        return nil
    }
    return res
}
func (m *OutlookUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OutlookUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMasterCategories()))
        for i, v := range m.GetMasterCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the masterCategories property value. A list of categories defined for the user.
// Parameters:
//  - value : Value to set for the masterCategories property.
func (m *OutlookUser) SetMasterCategories(value []OutlookCategory)() {
    m.masterCategories = value
}
