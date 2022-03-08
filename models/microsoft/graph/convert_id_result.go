package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConvertIdResult provides operations to call the translateExchangeIds method.
type ConvertIdResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
    errorDetails GenericErrorable;
    // The identifier that was converted. This value is the original, un-converted identifier.
    sourceId *string;
    // The converted identifier. This value is not present if the conversion failed.
    targetId *string;
}
// NewConvertIdResult instantiates a new convertIdResult and sets the default values.
func NewConvertIdResult()(*ConvertIdResult) {
    m := &ConvertIdResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConvertIdResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConvertIdResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConvertIdResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConvertIdResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorDetails gets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *ConvertIdResult) GetErrorDetails()(GenericErrorable) {
    if m == nil {
        return nil
    } else {
        return m.errorDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConvertIdResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGenericErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDetails(val.(GenericErrorable))
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
// GetSourceId gets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *ConvertIdResult) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTargetId gets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *ConvertIdResult) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
func (m *ConvertIdResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConvertIdResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConvertIdResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorDetails sets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *ConvertIdResult) SetErrorDetails(value GenericErrorable)() {
    if m != nil {
        m.errorDetails = value
    }
}
// SetSourceId sets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *ConvertIdResult) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTargetId sets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *ConvertIdResult) SetTargetId(value *string)() {
    if m != nil {
        m.targetId = value
    }
}
