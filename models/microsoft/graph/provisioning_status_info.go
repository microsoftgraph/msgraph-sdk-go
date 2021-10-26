package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ProvisioningStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    errorInformation *ProvisioningErrorInfo;
    // Possible values are: success, warning, failure, skipped, unknownFutureValue.
    status *ProvisioningResult;
}
// Instantiates a new provisioningStatusInfo and sets the default values.
func NewProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    m := &ProvisioningStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningStatusInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the errorInformation property value. 
func (m *ProvisioningStatusInfo) GetErrorInformation()(*ProvisioningErrorInfo) {
    if m == nil {
        return nil
    } else {
        return m.errorInformation
    }
}
// Gets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *ProvisioningStatusInfo) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *ProvisioningStatusInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["errorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningErrorInfo() })
        if err != nil {
            return err
        }
        m.SetErrorInformation(val.(*ProvisioningErrorInfo))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningResult)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *ProvisioningStatusInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ProvisioningStatusInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("errorInformation", m.GetErrorInformation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ProvisioningStatusInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the errorInformation property value. 
// Parameters:
//  - value : Value to set for the errorInformation property.
func (m *ProvisioningStatusInfo) SetErrorInformation(value *ProvisioningErrorInfo)() {
    m.errorInformation = value
}
// Sets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *ProvisioningStatusInfo) SetStatus(value *ProvisioningResult)() {
    m.status = value
}
