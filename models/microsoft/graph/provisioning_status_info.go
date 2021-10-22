package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisioningStatusInfo struct {
    additionalData map[string]interface{};
    errorInformation *ProvisioningErrorInfo;
    status *ProvisioningResult;
}
func NewProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    m := &ProvisioningStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProvisioningStatusInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProvisioningStatusInfo) GetErrorInformation()(*ProvisioningErrorInfo) {
    if m == nil {
        return nil
    } else {
        return m.errorInformation
    }
}
func (m *ProvisioningStatusInfo) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
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
func (m *ProvisioningStatusInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProvisioningStatusInfo) SetErrorInformation(value *ProvisioningErrorInfo)() {
    m.errorInformation = value
}
func (m *ProvisioningStatusInfo) SetStatus(value *ProvisioningResult)() {
    m.status = value
}
