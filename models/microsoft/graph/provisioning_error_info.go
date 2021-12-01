package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningErrorInfo 
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
// NewProvisioningErrorInfo instantiates a new provisioningErrorInfo and sets the default values.
func NewProvisioningErrorInfo()(*ProvisioningErrorInfo) {
    m := &ProvisioningErrorInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningErrorInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalDetails gets the additionalDetails property value. Additional details in case of error.
func (m *ProvisioningErrorInfo) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// GetErrorCategory gets the errorCategory property value. Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
func (m *ProvisioningErrorInfo) GetErrorCategory()(*ProvisioningStatusErrorCategory) {
    if m == nil {
        return nil
    } else {
        return m.errorCategory
    }
}
// GetErrorCode gets the errorCode property value. Unique error code if any occurred. Learn more
func (m *ProvisioningErrorInfo) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetReason gets the reason property value. Summarizes the status and describes why the status happened.
func (m *ProvisioningErrorInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetRecommendedAction gets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *ProvisioningErrorInfo) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningErrorInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val)
        }
        return nil
    }
    res["errorCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStatusErrorCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ProvisioningStatusErrorCategory)
            m.SetErrorCategory(&cast)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["recommendedAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
        }
        return nil
    }
    return res
}
func (m *ProvisioningErrorInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningErrorInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalDetails sets the additionalDetails property value. Additional details in case of error.
func (m *ProvisioningErrorInfo) SetAdditionalDetails(value *string)() {
    if m != nil {
        m.additionalDetails = value
    }
}
// SetErrorCategory sets the errorCategory property value. Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
func (m *ProvisioningErrorInfo) SetErrorCategory(value *ProvisioningStatusErrorCategory)() {
    if m != nil {
        m.errorCategory = value
    }
}
// SetErrorCode sets the errorCode property value. Unique error code if any occurred. Learn more
func (m *ProvisioningErrorInfo) SetErrorCode(value *string)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetReason sets the reason property value. Summarizes the status and describes why the status happened.
func (m *ProvisioningErrorInfo) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetRecommendedAction sets the recommendedAction property value. Provides the resolution for the corresponding error.
func (m *ProvisioningErrorInfo) SetRecommendedAction(value *string)() {
    if m != nil {
        m.recommendedAction = value
    }
}
