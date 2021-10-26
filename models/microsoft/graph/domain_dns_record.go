package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DomainDnsRecord struct {
    Entity
    // If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
    isOptional *bool;
    // Value used when configuring the name of the DNS record at the DNS host.
    label *string;
    // Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
    recordType *string;
    // Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
    supportedService *string;
    // Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
    ttl *int32;
}
// Instantiates a new domainDnsRecord and sets the default values.
func NewDomainDnsRecord()(*DomainDnsRecord) {
    m := &DomainDnsRecord{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
func (m *DomainDnsRecord) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// Gets the label property value. Value used when configuring the name of the DNS record at the DNS host.
func (m *DomainDnsRecord) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// Gets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
func (m *DomainDnsRecord) GetRecordType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordType
    }
}
// Gets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
func (m *DomainDnsRecord) GetSupportedService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedService
    }
}
// Gets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
func (m *DomainDnsRecord) GetTtl()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ttl
    }
}
// The deserialization information for the current model
func (m *DomainDnsRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOptional(val)
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLabel(val)
        return nil
    }
    res["recordType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecordType(val)
        return nil
    }
    res["supportedService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSupportedService(val)
        return nil
    }
    res["ttl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTtl(val)
        return nil
    }
    return res
}
func (m *DomainDnsRecord) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DomainDnsRecord) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordType", m.GetRecordType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportedService", m.GetSupportedService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ttl", m.GetTtl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
// Parameters:
//  - value : Value to set for the isOptional property.
func (m *DomainDnsRecord) SetIsOptional(value *bool)() {
    m.isOptional = value
}
// Sets the label property value. Value used when configuring the name of the DNS record at the DNS host.
// Parameters:
//  - value : Value to set for the label property.
func (m *DomainDnsRecord) SetLabel(value *string)() {
    m.label = value
}
// Sets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
// Parameters:
//  - value : Value to set for the recordType property.
func (m *DomainDnsRecord) SetRecordType(value *string)() {
    m.recordType = value
}
// Sets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
// Parameters:
//  - value : Value to set for the supportedService property.
func (m *DomainDnsRecord) SetSupportedService(value *string)() {
    m.supportedService = value
}
// Sets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
// Parameters:
//  - value : Value to set for the ttl property.
func (m *DomainDnsRecord) SetTtl(value *int32)() {
    m.ttl = value
}
