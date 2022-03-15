package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ColumnValidation provides operations to manage the drive singleton.
type ColumnValidation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Default BCP 47 language tag for the description.
    defaultLanguage *string;
    // Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
    descriptions []DisplayNameLocalizationable;
    // The formula to validate column value. For examples, see Examples of common formulas in lists.
    formula *string;
}
// NewColumnValidation instantiates a new columnValidation and sets the default values.
func NewColumnValidation()(*ColumnValidation) {
    m := &ColumnValidation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateColumnValidationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnValidationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewColumnValidation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultLanguage gets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) GetDefaultLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguage
    }
}
// GetDescriptions gets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) GetDescriptions()([]DisplayNameLocalizationable) {
    if m == nil {
        return nil
    } else {
        return m.descriptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnValidation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguage(val)
        }
        return nil
    }
    res["descriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayNameLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayNameLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayNameLocalizationable)
            }
            m.SetDescriptions(res)
        }
        return nil
    }
    res["formula"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormula(val)
        }
        return nil
    }
    return res
}
// GetFormula gets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
func (m *ColumnValidation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ColumnValidation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    if m.GetDescriptions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultLanguage sets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) SetDefaultLanguage(value *string)() {
    if m != nil {
        m.defaultLanguage = value
    }
}
// SetDescriptions sets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) SetDescriptions(value []DisplayNameLocalizationable)() {
    if m != nil {
        m.descriptions = value
    }
}
// SetFormula sets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) SetFormula(value *string)() {
    if m != nil {
        m.formula = value
    }
}
