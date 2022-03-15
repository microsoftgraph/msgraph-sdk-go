package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CurrencyColumn provides operations to manage the collection of drive entities.
type CurrencyColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the locale from which to infer the currency symbol.
    locale *string;
}
// NewCurrencyColumn instantiates a new currencyColumn and sets the default values.
func NewCurrencyColumn()(*CurrencyColumn) {
    m := &CurrencyColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCurrencyColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCurrencyColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCurrencyColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CurrencyColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CurrencyColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["locale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    return res
}
// GetLocale gets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
func (m *CurrencyColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CurrencyColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
func (m *CurrencyColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLocale sets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) SetLocale(value *string)() {
    if m != nil {
        m.locale = value
    }
}
