package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudAppSecurityState struct {
    additionalData map[string]interface{};
    destinationServiceIp *string;
    destinationServiceName *string;
    riskScore *string;
}
func NewCloudAppSecurityState()(*CloudAppSecurityState) {
    m := &CloudAppSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CloudAppSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CloudAppSecurityState) GetDestinationServiceIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceIp
    }
}
func (m *CloudAppSecurityState) GetDestinationServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceName
    }
}
func (m *CloudAppSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *CloudAppSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationServiceIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationServiceIp(val)
        return nil
    }
    res["destinationServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationServiceName(val)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    return res
}
func (m *CloudAppSecurityState) IsNil()(bool) {
    return m == nil
}
func (m *CloudAppSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationServiceIp", m.GetDestinationServiceIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationServiceName", m.GetDestinationServiceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
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
func (m *CloudAppSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CloudAppSecurityState) SetDestinationServiceIp(value *string)() {
    m.destinationServiceIp = value
}
func (m *CloudAppSecurityState) SetDestinationServiceName(value *string)() {
    m.destinationServiceName = value
}
func (m *CloudAppSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
