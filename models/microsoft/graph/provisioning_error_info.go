package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisioningErrorInfo struct {
    additionalData map[string]interface{};
    additionalDetails *string;
    errorCategory *ProvisioningStatusErrorCategory;
    errorCode *string;
    reason *string;
    recommendedAction *string;
}
func NewProvisioningErrorInfo()(*ProvisioningErrorInfo) {
    m := &ProvisioningErrorInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProvisioningErrorInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProvisioningErrorInfo) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
func (m *ProvisioningErrorInfo) GetErrorCategory()(*ProvisioningStatusErrorCategory) {
    if m == nil {
        return nil
    } else {
        return m.errorCategory
    }
}
func (m *ProvisioningErrorInfo) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *ProvisioningErrorInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *ProvisioningErrorInfo) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
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
func (m *ProvisioningErrorInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProvisioningErrorInfo) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
func (m *ProvisioningErrorInfo) SetErrorCategory(value *ProvisioningStatusErrorCategory)() {
    m.errorCategory = value
}
func (m *ProvisioningErrorInfo) SetErrorCode(value *string)() {
    m.errorCode = value
}
func (m *ProvisioningErrorInfo) SetReason(value *string)() {
    m.reason = value
}
func (m *ProvisioningErrorInfo) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
