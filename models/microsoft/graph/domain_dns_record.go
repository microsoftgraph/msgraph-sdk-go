package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DomainDnsRecord struct {
    Entity
    isOptional *bool;
    label *string;
    recordType *string;
    supportedService *string;
    ttl *int32;
}
func NewDomainDnsRecord()(*DomainDnsRecord) {
    m := &DomainDnsRecord{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DomainDnsRecord) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
func (m *DomainDnsRecord) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
func (m *DomainDnsRecord) GetRecordType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordType
    }
}
func (m *DomainDnsRecord) GetSupportedService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedService
    }
}
func (m *DomainDnsRecord) GetTtl()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ttl
    }
}
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
func (m *DomainDnsRecord) SetIsOptional(value *bool)() {
    m.isOptional = value
}
func (m *DomainDnsRecord) SetLabel(value *string)() {
    m.label = value
}
func (m *DomainDnsRecord) SetRecordType(value *string)() {
    m.recordType = value
}
func (m *DomainDnsRecord) SetSupportedService(value *string)() {
    m.supportedService = value
}
func (m *DomainDnsRecord) SetTtl(value *int32)() {
    m.ttl = value
}
