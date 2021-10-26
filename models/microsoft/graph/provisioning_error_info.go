package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ProvisioningErrorInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Additional details in case of error.
    additionalDetails *string;
    // Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
    errorCategory *ProvisioningStatusErrorCategory;
    // Unique error code if any occurred. Learn more
    errorCode *string;
    // Summarizes the status and describes why the status happened.
    reason *string;
    // Provides the resolution for the corresponding error.
    recommendedAction *string;
}
// Instantiates a new provisioningErrorInfo and sets the default values.
func NewProvisioningErrorInfo()(*ProvisioningErrorInfo) {
    m := &ProvisioningErrorInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningErrorInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the additionalDetails property value. Additional details in case of error.
func (m *ProvisioningErrorInfo) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// Gets the errorCategory property value. Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
func (m *ProvisioningErrorInfo) GetErrorCategory()(*ProvisioningStatusErrorCategory) {
    if m == nil {
        return nil
    } else {
        return m.errorCategory
    }
}
// Gets the errorCode property value. Unique error code if any occurred. Learn more
func (m *ProvisioningErrorInfo) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the reason property value. Summarizes the status and describes why the status happened.
func (m *ProvisioningErrorInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Gets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *ProvisioningErrorInfo) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
// The deserialization information for the current model
func (m *ProvisioningErrorInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalDetails(val)
        return nil
    }
    res["errorCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStatusErrorCategory)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningStatusErrorCategory)
        m.SetErrorCategory(&cast)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["recommendedAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecommendedAction(val)
        return nil
    }
    return res
}
func (m *ProvisioningErrorInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ProvisioningErrorInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    if m.GetErrorCategory() != nil {
        cast := m.GetErrorCategory().String()
        err := writer.WriteStringValue("errorCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
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
func (m *ProvisioningErrorInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the additionalDetails property value. Additional details in case of error.
// Parameters:
//  - value : Value to set for the additionalDetails property.
func (m *ProvisioningErrorInfo) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
// Sets the errorCategory property value. Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
// Parameters:
//  - value : Value to set for the errorCategory property.
func (m *ProvisioningErrorInfo) SetErrorCategory(value *ProvisioningStatusErrorCategory)() {
    m.errorCategory = value
}
// Sets the errorCode property value. Unique error code if any occurred. Learn more
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *ProvisioningErrorInfo) SetErrorCode(value *string)() {
    m.errorCode = value
}
// Sets the reason property value. Summarizes the status and describes why the status happened.
// Parameters:
//  - value : Value to set for the reason property.
func (m *ProvisioningErrorInfo) SetReason(value *string)() {
    m.reason = value
}
// Sets the recommendedAction property value. Provides the resolution for the corresponding error.
// Parameters:
//  - value : Value to set for the recommendedAction property.
func (m *ProvisioningErrorInfo) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
