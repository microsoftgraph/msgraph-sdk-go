package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ParentalControlSettings struct {
    additionalData map[string]interface{};
    countriesBlockedForMinors []string;
    legalAgeGroupRule *string;
}
func NewParentalControlSettings()(*ParentalControlSettings) {
    m := &ParentalControlSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ParentalControlSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ParentalControlSettings) GetCountriesBlockedForMinors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.countriesBlockedForMinors
    }
}
func (m *ParentalControlSettings) GetLegalAgeGroupRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legalAgeGroupRule
    }
}
func (m *ParentalControlSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["countriesBlockedForMinors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCountriesBlockedForMinors(res)
        return nil
    }
    res["legalAgeGroupRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLegalAgeGroupRule(val)
        return nil
    }
    return res
}
func (m *ParentalControlSettings) IsNil()(bool) {
    return m == nil
}
func (m *ParentalControlSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("countriesBlockedForMinors", m.GetCountriesBlockedForMinors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("legalAgeGroupRule", m.GetLegalAgeGroupRule())
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
func (m *ParentalControlSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ParentalControlSettings) SetCountriesBlockedForMinors(value []string)() {
    m.countriesBlockedForMinors = value
}
func (m *ParentalControlSettings) SetLegalAgeGroupRule(value *string)() {
    m.legalAgeGroupRule = value
}
