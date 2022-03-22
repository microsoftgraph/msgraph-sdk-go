package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditLogRoot provides operations to manage the auditLogRoot singleton.
type AuditLogRoot struct {
    Entity
    // Read-only. Nullable.
    directoryAudits []DirectoryAuditable;
    // 
    provisioning []ProvisioningObjectSummaryable;
    // 
    restrictedSignIns []RestrictedSignInable;
    // Read-only. Nullable.
    signIns []SignInable;
}
// NewAuditLogRoot instantiates a new auditLogRoot and sets the default values.
func NewAuditLogRoot()(*AuditLogRoot) {
    m := &AuditLogRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuditLogRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogRootFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAuditLogRoot(), nil
}
// GetDirectoryAudits gets the directoryAudits property value. Read-only. Nullable.
func (m *AuditLogRoot) GetDirectoryAudits()([]DirectoryAuditable) {
    if m == nil {
        return nil
    } else {
        return m.directoryAudits
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["directoryAudits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryAuditable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryAuditable)
            }
            m.SetDirectoryAudits(res)
        }
        return nil
    }
    res["provisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningObjectSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisioningObjectSummaryable)
            }
            m.SetProvisioning(res)
        }
        return nil
    }
    res["restrictedSignIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRestrictedSignInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RestrictedSignInable, len(val))
            for i, v := range val {
                res[i] = v.(RestrictedSignInable)
            }
            m.SetRestrictedSignIns(res)
        }
        return nil
    }
    res["signIns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignInable, len(val))
            for i, v := range val {
                res[i] = v.(SignInable)
            }
            m.SetSignIns(res)
        }
        return nil
    }
    return res
}
// GetProvisioning gets the provisioning property value. 
func (m *AuditLogRoot) GetProvisioning()([]ProvisioningObjectSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.provisioning
    }
}
// GetRestrictedSignIns gets the restrictedSignIns property value. 
func (m *AuditLogRoot) GetRestrictedSignIns()([]RestrictedSignInable) {
    if m == nil {
        return nil
    } else {
        return m.restrictedSignIns
    }
}
// GetSignIns gets the signIns property value. Read-only. Nullable.
func (m *AuditLogRoot) GetSignIns()([]SignInable) {
    if m == nil {
        return nil
    } else {
        return m.signIns
    }
}
// Serialize serializes information the current object
func (m *AuditLogRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectoryAudits() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectoryAudits()))
        for i, v := range m.GetDirectoryAudits() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("directoryAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioning() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisioning()))
        for i, v := range m.GetProvisioning() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("provisioning", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedSignIns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRestrictedSignIns()))
        for i, v := range m.GetRestrictedSignIns() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("restrictedSignIns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignIns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSignIns()))
        for i, v := range m.GetSignIns() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("signIns", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryAudits sets the directoryAudits property value. Read-only. Nullable.
func (m *AuditLogRoot) SetDirectoryAudits(value []DirectoryAuditable)() {
    if m != nil {
        m.directoryAudits = value
    }
}
// SetProvisioning sets the provisioning property value. 
func (m *AuditLogRoot) SetProvisioning(value []ProvisioningObjectSummaryable)() {
    if m != nil {
        m.provisioning = value
    }
}
// SetRestrictedSignIns sets the restrictedSignIns property value. 
func (m *AuditLogRoot) SetRestrictedSignIns(value []RestrictedSignInable)() {
    if m != nil {
        m.restrictedSignIns = value
    }
}
// SetSignIns sets the signIns property value. Read-only. Nullable.
func (m *AuditLogRoot) SetSignIns(value []SignInable)() {
    if m != nil {
        m.signIns = value
    }
}
