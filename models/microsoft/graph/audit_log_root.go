package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuditLogRoot struct {
    Entity
    directoryAudits []DirectoryAudit;
    provisioning []ProvisioningObjectSummary;
    restrictedSignIns []RestrictedSignIn;
    signIns []SignIn;
}
func NewAuditLogRoot()(*AuditLogRoot) {
    m := &AuditLogRoot{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AuditLogRoot) GetDirectoryAudits()([]DirectoryAudit) {
    if m == nil {
        return nil
    } else {
        return m.directoryAudits
    }
}
func (m *AuditLogRoot) GetProvisioning()([]ProvisioningObjectSummary) {
    if m == nil {
        return nil
    } else {
        return m.provisioning
    }
}
func (m *AuditLogRoot) GetRestrictedSignIns()([]RestrictedSignIn) {
    if m == nil {
        return nil
    } else {
        return m.restrictedSignIns
    }
}
func (m *AuditLogRoot) GetSignIns()([]SignIn) {
    if m == nil {
        return nil
    } else {
        return m.signIns
    }
}
func (m *AuditLogRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["directoryAudits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryAudit() })
        if err != nil {
            return err
        }
        res := make([]DirectoryAudit, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryAudit))
        }
        m.SetDirectoryAudits(res)
        return nil
    }
    res["provisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningObjectSummary() })
        if err != nil {
            return err
        }
        res := make([]ProvisioningObjectSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProvisioningObjectSummary))
        }
        m.SetProvisioning(res)
        return nil
    }
    res["restrictedSignIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRestrictedSignIn() })
        if err != nil {
            return err
        }
        res := make([]RestrictedSignIn, len(val))
        for i, v := range val {
            res[i] = *(v.(*RestrictedSignIn))
        }
        m.SetRestrictedSignIns(res)
        return nil
    }
    res["signIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignIn() })
        if err != nil {
            return err
        }
        res := make([]SignIn, len(val))
        for i, v := range val {
            res[i] = *(v.(*SignIn))
        }
        m.SetSignIns(res)
        return nil
    }
    return res
}
func (m *AuditLogRoot) IsNil()(bool) {
    return m == nil
}
func (m *AuditLogRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectoryAudits()))
        for i, v := range m.GetDirectoryAudits() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("directoryAudits", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisioning()))
        for i, v := range m.GetProvisioning() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("provisioning", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRestrictedSignIns()))
        for i, v := range m.GetRestrictedSignIns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("restrictedSignIns", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSignIns()))
        for i, v := range m.GetSignIns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("signIns", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AuditLogRoot) SetDirectoryAudits(value []DirectoryAudit)() {
    m.directoryAudits = value
}
func (m *AuditLogRoot) SetProvisioning(value []ProvisioningObjectSummary)() {
    m.provisioning = value
}
func (m *AuditLogRoot) SetRestrictedSignIns(value []RestrictedSignIn)() {
    m.restrictedSignIns = value
}
func (m *AuditLogRoot) SetSignIns(value []SignIn)() {
    m.signIns = value
}
