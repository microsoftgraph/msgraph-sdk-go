package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningStatusInfo 
type ProvisioningStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    errorInformation *ProvisioningErrorInfo;
    // Possible values are: success, warning, failure, skipped, unknownFutureValue.
    status *ProvisioningResult;
}
// NewProvisioningStatusInfo instantiates a new provisioningStatusInfo and sets the default values.
func NewProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    m := &ProvisioningStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningStatusInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorInformation gets the errorInformation property value. 
func (m *ProvisioningStatusInfo) GetErrorInformation()(*ProvisioningErrorInfo) {
    if m == nil {
        return nil
    } else {
        return m.errorInformation
    }
}
// GetStatus gets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *ProvisioningStatusInfo) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningStatusInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningErrorInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorInformation(val.(*ProvisioningErrorInfo))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProvisioningResult))
        }
        return nil
    }
    return res
}
func (m *ProvisioningStatusInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProvisioningStatusInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("errorInformation", m.GetErrorInformation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ProvisioningStatusInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorInformation sets the errorInformation property value. 
func (m *ProvisioningStatusInfo) SetErrorInformation(value *ProvisioningErrorInfo)() {
    if m != nil {
        m.errorInformation = value
    }
}
// SetStatus sets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *ProvisioningStatusInfo) SetStatus(value *ProvisioningResult)() {
    if m != nil {
        m.status = value
    }
}
