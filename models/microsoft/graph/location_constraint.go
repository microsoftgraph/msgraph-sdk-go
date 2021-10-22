package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LocationConstraint struct {
    additionalData map[string]interface{};
    isRequired *bool;
    locations []LocationConstraintItem;
    suggestLocation *bool;
}
func NewLocationConstraint()(*LocationConstraint) {
    m := &LocationConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LocationConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LocationConstraint) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
func (m *LocationConstraint) GetLocations()([]LocationConstraintItem) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
func (m *LocationConstraint) GetSuggestLocation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.suggestLocation
    }
}
func (m *LocationConstraint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRequired(val)
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocationConstraintItem() })
        if err != nil {
            return err
        }
        res := make([]LocationConstraintItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*LocationConstraintItem))
        }
        m.SetLocations(res)
        return nil
    }
    res["suggestLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSuggestLocation(val)
        return nil
    }
    return res
}
func (m *LocationConstraint) IsNil()(bool) {
    return m == nil
}
func (m *LocationConstraint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("locations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("suggestLocation", m.GetSuggestLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *LocationConstraint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LocationConstraint) SetIsRequired(value *bool)() {
    m.isRequired = value
}
func (m *LocationConstraint) SetLocations(value []LocationConstraintItem)() {
    m.locations = value
}
func (m *LocationConstraint) SetSuggestLocation(value *bool)() {
    m.suggestLocation = value
}
