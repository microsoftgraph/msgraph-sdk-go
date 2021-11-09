package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type HostSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
    fqdn *string;
    // 
    isAzureAdJoined *bool;
    // 
    isAzureAdRegistered *bool;
    // True if the host is domain joined to an on-premises Active Directory domain.
    isHybridAzureDomainJoined *bool;
    // The local host name, without the DNS domain name.
    netBiosName *string;
    // Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
    os *string;
    // Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
    privateIpAddress *string;
    // Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
    publicIpAddress *string;
    // Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
}
// Instantiates a new hostSecurityState and sets the default values.
func NewHostSecurityState()(*HostSecurityState) {
    m := &HostSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HostSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the fqdn property value. Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
func (m *HostSecurityState) GetFqdn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fqdn
    }
}
// Gets the isAzureAdJoined property value. 
func (m *HostSecurityState) GetIsAzureAdJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdJoined
    }
}
// Gets the isAzureAdRegistered property value. 
func (m *HostSecurityState) GetIsAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdRegistered
    }
}
// Gets the isHybridAzureDomainJoined property value. True if the host is domain joined to an on-premises Active Directory domain.
func (m *HostSecurityState) GetIsHybridAzureDomainJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureDomainJoined
    }
}
// Gets the netBiosName property value. The local host name, without the DNS domain name.
func (m *HostSecurityState) GetNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.netBiosName
    }
}
// Gets the os property value. Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
func (m *HostSecurityState) GetOs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.os
    }
}
// Gets the privateIpAddress property value. Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
func (m *HostSecurityState) GetPrivateIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privateIpAddress
    }
}
// Gets the publicIpAddress property value. Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
func (m *HostSecurityState) GetPublicIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicIpAddress
    }
}
// Gets the riskScore property value. Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
func (m *HostSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// The deserialization information for the current model
func (m *HostSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fqdn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["isAzureAdJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdJoined(val)
        }
        return nil
    }
    res["isAzureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdRegistered(val)
        }
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureDomainJoined(val)
        }
        return nil
    }
    res["netBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetBiosName(val)
        }
        return nil
    }
    res["os"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["privateIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateIpAddress(val)
        }
        return nil
    }
    res["publicIpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIpAddress(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    return res
}
func (m *HostSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *HostSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the fqdn property value. Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
// Parameters:
//  - value : Value to set for the fqdn property.
func (m *HostSecurityState) SetFqdn(value *string)() {
    m.fqdn = value
}
// Sets the isAzureAdJoined property value. 
// Parameters:
//  - value : Value to set for the isAzureAdJoined property.
func (m *HostSecurityState) SetIsAzureAdJoined(value *bool)() {
    m.isAzureAdJoined = value
}
// Sets the isAzureAdRegistered property value. 
// Parameters:
//  - value : Value to set for the isAzureAdRegistered property.
func (m *HostSecurityState) SetIsAzureAdRegistered(value *bool)() {
    m.isAzureAdRegistered = value
}
// Sets the isHybridAzureDomainJoined property value. True if the host is domain joined to an on-premises Active Directory domain.
// Parameters:
//  - value : Value to set for the isHybridAzureDomainJoined property.
func (m *HostSecurityState) SetIsHybridAzureDomainJoined(value *bool)() {
    m.isHybridAzureDomainJoined = value
}
// Sets the netBiosName property value. The local host name, without the DNS domain name.
// Parameters:
//  - value : Value to set for the netBiosName property.
func (m *HostSecurityState) SetNetBiosName(value *string)() {
    m.netBiosName = value
}
// Sets the os property value. Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
// Parameters:
//  - value : Value to set for the os property.
func (m *HostSecurityState) SetOs(value *string)() {
    m.os = value
}
// Sets the privateIpAddress property value. Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
// Parameters:
//  - value : Value to set for the privateIpAddress property.
func (m *HostSecurityState) SetPrivateIpAddress(value *string)() {
    m.privateIpAddress = value
}
// Sets the publicIpAddress property value. Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
// Parameters:
//  - value : Value to set for the publicIpAddress property.
func (m *HostSecurityState) SetPublicIpAddress(value *string)() {
    m.publicIpAddress = value
}
// Sets the riskScore property value. Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *HostSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
