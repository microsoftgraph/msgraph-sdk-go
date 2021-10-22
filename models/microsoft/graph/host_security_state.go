package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type HostSecurityState struct {
    additionalData map[string]interface{};
    fqdn *string;
    isAzureAdJoined *bool;
    isAzureAdRegistered *bool;
    isHybridAzureDomainJoined *bool;
    netBiosName *string;
    os *string;
    privateIpAddress *string;
    publicIpAddress *string;
    riskScore *string;
}
func NewHostSecurityState()(*HostSecurityState) {
    m := &HostSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *HostSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *HostSecurityState) GetFqdn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fqdn
    }
}
func (m *HostSecurityState) GetIsAzureAdJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdJoined
    }
}
func (m *HostSecurityState) GetIsAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdRegistered
    }
}
func (m *HostSecurityState) GetIsHybridAzureDomainJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureDomainJoined
    }
}
func (m *HostSecurityState) GetNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.netBiosName
    }
}
func (m *HostSecurityState) GetOs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.os
    }
}
func (m *HostSecurityState) GetPrivateIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privateIpAddress
    }
}
func (m *HostSecurityState) GetPublicIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicIpAddress
    }
}
func (m *HostSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *HostSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fqdn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFqdn(val)
        return nil
    }
    res["isAzureAdJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAzureAdJoined(val)
        return nil
    }
    res["isAzureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAzureAdRegistered(val)
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsHybridAzureDomainJoined(val)
        return nil
    }
    res["netBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetBiosName(val)
        return nil
    }
    res["os"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOs(val)
        return nil
    }
    res["privateIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrivateIpAddress(val)
        return nil
    }
    res["publicIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublicIpAddress(val)
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
func (m *HostSecurityState) IsNil()(bool) {
    return m == nil
}
func (m *HostSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fqdn", m.GetFqdn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAzureAdJoined", m.GetIsAzureAdJoined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAzureAdRegistered", m.GetIsAzureAdRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHybridAzureDomainJoined", m.GetIsHybridAzureDomainJoined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("netBiosName", m.GetNetBiosName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privateIpAddress", m.GetPrivateIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publicIpAddress", m.GetPublicIpAddress())
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
func (m *HostSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *HostSecurityState) SetFqdn(value *string)() {
    m.fqdn = value
}
func (m *HostSecurityState) SetIsAzureAdJoined(value *bool)() {
    m.isAzureAdJoined = value
}
func (m *HostSecurityState) SetIsAzureAdRegistered(value *bool)() {
    m.isAzureAdRegistered = value
}
func (m *HostSecurityState) SetIsHybridAzureDomainJoined(value *bool)() {
    m.isHybridAzureDomainJoined = value
}
func (m *HostSecurityState) SetNetBiosName(value *string)() {
    m.netBiosName = value
}
func (m *HostSecurityState) SetOs(value *string)() {
    m.os = value
}
func (m *HostSecurityState) SetPrivateIpAddress(value *string)() {
    m.privateIpAddress = value
}
func (m *HostSecurityState) SetPublicIpAddress(value *string)() {
    m.publicIpAddress = value
}
func (m *HostSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
