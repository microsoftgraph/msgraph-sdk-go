package redirect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type RedirectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    configuration *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration;
    // 
    destinationPrinterId *string;
}
// Instantiates a new redirectRequestBody and sets the default values.
func NewRedirectRequestBody()(*RedirectRequestBody) {
    m := &RedirectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the configuration property value. 
func (m *RedirectRequestBody) GetConfiguration()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// Gets the destinationPrinterId property value. 
func (m *RedirectRequestBody) GetDestinationPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationPrinterId
    }
}
// The deserialization information for the current model
func (m *RedirectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrintJobConfiguration() })
        if err != nil {
            return err
        }
        m.SetConfiguration(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration))
        return nil
    }
    res["destinationPrinterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationPrinterId(val)
        return nil
    }
    return res
}
func (m *RedirectRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RedirectRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationPrinterId", m.GetDestinationPrinterId())
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
func (m *RedirectRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the configuration property value. 
// Parameters:
//  - value : Value to set for the configuration property.
func (m *RedirectRequestBody) SetConfiguration(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration)() {
    m.configuration = value
}
// Sets the destinationPrinterId property value. 
// Parameters:
//  - value : Value to set for the destinationPrinterId property.
func (m *RedirectRequestBody) SetDestinationPrinterId(value *string)() {
    m.destinationPrinterId = value
}
