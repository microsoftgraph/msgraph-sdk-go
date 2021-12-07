package redirect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// RedirectRequestBody 
type RedirectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    configuration *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration;
    // 
    destinationPrinterId *string;
}
// NewRedirectRequestBody instantiates a new redirectRequestBody and sets the default values.
func NewRedirectRequestBody()(*RedirectRequestBody) {
    m := &RedirectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfiguration gets the configuration property value. 
func (m *RedirectRequestBody) GetConfiguration()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetDestinationPrinterId gets the destinationPrinterId property value. 
func (m *RedirectRequestBody) GetDestinationPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationPrinterId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrintJobConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration))
        }
        return nil
    }
    res["destinationPrinterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationPrinterId(val)
        }
        return nil
    }
    return res
}
func (m *RedirectRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfiguration sets the configuration property value. 
func (m *RedirectRequestBody) SetConfiguration(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintJobConfiguration)() {
    if m != nil {
        m.configuration = value
    }
}
// SetDestinationPrinterId sets the destinationPrinterId property value. 
func (m *RedirectRequestBody) SetDestinationPrinterId(value *string)() {
    if m != nil {
        m.destinationPrinterId = value
    }
}
