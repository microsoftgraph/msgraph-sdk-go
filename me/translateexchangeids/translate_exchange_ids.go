package translateexchangeids

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// TranslateExchangeIds 
type TranslateExchangeIds struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
    errorDetails *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GenericError;
    // The identifier that was converted. This value is the original, un-converted identifier.
    sourceId *string;
    // The converted identifier. This value is not present if the conversion failed.
    targetId *string;
}
// NewTranslateExchangeIds instantiates a new translateExchangeIds and sets the default values.
func NewTranslateExchangeIds()(*TranslateExchangeIds) {
    m := &TranslateExchangeIds{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslateExchangeIds) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorDetails gets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *TranslateExchangeIds) GetErrorDetails()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GenericError) {
    if m == nil {
        return nil
    } else {
        return m.errorDetails
    }
}
// GetSourceId gets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *TranslateExchangeIds) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTargetId gets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *TranslateExchangeIds) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslateExchangeIds) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewGenericError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDetails(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GenericError))
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    return res
}
func (m *TranslateExchangeIds) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TranslateExchangeIds) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("errorDetails", m.GetErrorDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetId", m.GetTargetId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslateExchangeIds) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetErrorDetails sets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *TranslateExchangeIds) SetErrorDetails(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GenericError)() {
    m.errorDetails = value
}
// SetSourceId sets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *TranslateExchangeIds) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTargetId sets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *TranslateExchangeIds) SetTargetId(value *string)() {
    m.targetId = value
}
