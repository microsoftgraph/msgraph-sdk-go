package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ColumnValidation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Default BCP 47 language tag for the description.
    defaultLanguage *string;
    // Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
    descriptions []DisplayNameLocalization;
    // The formula to validate column value. For examples, see Examples of common formulas in lists.
    formula *string;
}
// Instantiates a new columnValidation and sets the default values.
func NewColumnValidation()(*ColumnValidation) {
    m := &ColumnValidation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) GetDefaultLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguage
    }
}
// Gets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) GetDescriptions()([]DisplayNameLocalization) {
    if m == nil {
        return nil
    } else {
        return m.descriptions
    }
}
// Gets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
// The deserialization information for the current model
func (m *ColumnValidation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLanguage(val)
        return nil
    }
    res["descriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDisplayNameLocalization() })
        if err != nil {
            return err
        }
        res := make([]DisplayNameLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*DisplayNameLocalization))
        }
        m.SetDescriptions(res)
        return nil
    }
    res["formula"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormula(val)
        return nil
    }
    return res
}
func (m *ColumnValidation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ColumnValidation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ColumnValidation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the defaultLanguage property value. Default BCP 47 language tag for the description.
// Parameters:
//  - value : Value to set for the defaultLanguage property.
func (m *ColumnValidation) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
// Sets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
// Parameters:
//  - value : Value to set for the descriptions property.
func (m *ColumnValidation) SetDescriptions(value []DisplayNameLocalization)() {
    m.descriptions = value
}
// Sets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
// Parameters:
//  - value : Value to set for the formula property.
func (m *ColumnValidation) SetFormula(value *string)() {
    m.formula = value
}
