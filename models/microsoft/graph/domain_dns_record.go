package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DomainDnsRecord 
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
// NewDomainDnsRecord instantiates a new domainDnsRecord and sets the default values.
func NewDomainDnsRecord()(*DomainDnsRecord) {
    m := &DomainDnsRecord{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDomainDnsRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsRecordFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDomainDnsRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["recordType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordType(val)
        }
        return nil
    }
    res["supportedService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedService(val)
        }
        return nil
    }
    res["ttl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTtl(val)
        }
        return nil
    }
    return res
}
// GetIsOptional gets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
func (m *DomainDnsRecord) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// GetLabel gets the label property value. Value used when configuring the name of the DNS record at the DNS host.
func (m *DomainDnsRecord) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetRecordType gets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
func (m *DomainDnsRecord) GetRecordType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordType
    }
}
// GetSupportedService gets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
func (m *DomainDnsRecord) GetSupportedService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedService
    }
}
// GetTtl gets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
func (m *DomainDnsRecord) GetTtl()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ttl
    }
}
// Serialize serializes information the current object
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
// SetIsOptional sets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
func (m *DomainDnsRecord) SetIsOptional(value *bool)() {
    if m != nil {
        m.isOptional = value
    }
}
// SetLabel sets the label property value. Value used when configuring the name of the DNS record at the DNS host.
func (m *DomainDnsRecord) SetLabel(value *string)() {
    if m != nil {
        m.label = value
    }
}
// SetRecordType sets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
func (m *DomainDnsRecord) SetRecordType(value *string)() {
    if m != nil {
        m.recordType = value
    }
}
// SetSupportedService sets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
func (m *DomainDnsRecord) SetSupportedService(value *string)() {
    if m != nil {
        m.supportedService = value
    }
}
// SetTtl sets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
func (m *DomainDnsRecord) SetTtl(value *int32)() {
    if m != nil {
        m.ttl = value
    }
}
