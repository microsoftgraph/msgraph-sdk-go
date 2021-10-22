package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Domain struct {
    Entity
    authenticationType *string;
    availabilityStatus *string;
    domainNameReferences []DirectoryObject;
    isAdminManaged *bool;
    isDefault *bool;
    isInitial *bool;
    isRoot *bool;
    isVerified *bool;
    manufacturer *string;
    model *string;
    passwordNotificationWindowInDays *int32;
    passwordValidityPeriodInDays *int32;
    serviceConfigurationRecords []DomainDnsRecord;
    state *DomainState;
    supportedServices []string;
    verificationDnsRecords []DomainDnsRecord;
}
func NewDomain()(*Domain) {
    m := &Domain{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Domain) GetAuthenticationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationType
    }
}
func (m *Domain) GetAvailabilityStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availabilityStatus
    }
}
func (m *Domain) GetDomainNameReferences()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.domainNameReferences
    }
}
func (m *Domain) GetIsAdminManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAdminManaged
    }
}
func (m *Domain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *Domain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
func (m *Domain) GetIsRoot()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoot
    }
}
func (m *Domain) GetIsVerified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVerified
    }
}
func (m *Domain) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *Domain) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *Domain) GetPasswordNotificationWindowInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordNotificationWindowInDays
    }
}
func (m *Domain) GetPasswordValidityPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordValidityPeriodInDays
    }
}
func (m *Domain) GetServiceConfigurationRecords()([]DomainDnsRecord) {
    if m == nil {
        return nil
    } else {
        return m.serviceConfigurationRecords
    }
}
func (m *Domain) GetState()(*DomainState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *Domain) GetSupportedServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedServices
    }
}
func (m *Domain) GetVerificationDnsRecords()([]DomainDnsRecord) {
    if m == nil {
        return nil
    } else {
        return m.verificationDnsRecords
    }
}
func (m *Domain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationType(val)
        return nil
    }
    res["availabilityStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAvailabilityStatus(val)
        return nil
    }
    res["domainNameReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDomainNameReferences(res)
        return nil
    }
    res["isAdminManaged"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAdminManaged(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isInitial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInitial(val)
        return nil
    }
    res["isRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRoot(val)
        return nil
    }
    res["isVerified"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVerified(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["passwordNotificationWindowInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasswordNotificationWindowInDays(val)
        return nil
    }
    res["passwordValidityPeriodInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasswordValidityPeriodInDays(val)
        return nil
    }
    res["serviceConfigurationRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainDnsRecord() })
        if err != nil {
            return err
        }
        res := make([]DomainDnsRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*DomainDnsRecord))
        }
        m.SetServiceConfigurationRecords(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainState() })
        if err != nil {
            return err
        }
        m.SetState(val.(*DomainState))
        return nil
    }
    res["supportedServices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSupportedServices(res)
        return nil
    }
    res["verificationDnsRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainDnsRecord() })
        if err != nil {
            return err
        }
        res := make([]DomainDnsRecord, len(val))
        for i, v := range val {
            res[i] = *(v.(*DomainDnsRecord))
        }
        m.SetVerificationDnsRecords(res)
        return nil
    }
    return res
}
func (m *Domain) IsNil()(bool) {
    return m == nil
}
func (m *Domain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationType", m.GetAuthenticationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("availabilityStatus", m.GetAvailabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDomainNameReferences()))
        for i, v := range m.GetDomainNameReferences() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("domainNameReferences", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAdminManaged", m.GetIsAdminManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRoot", m.GetIsRoot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isVerified", m.GetIsVerified())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordNotificationWindowInDays", m.GetPasswordNotificationWindowInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordValidityPeriodInDays", m.GetPasswordValidityPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServiceConfigurationRecords()))
        for i, v := range m.GetServiceConfigurationRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("serviceConfigurationRecords", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedServices", m.GetSupportedServices())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVerificationDnsRecords()))
        for i, v := range m.GetVerificationDnsRecords() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("verificationDnsRecords", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Domain) SetAuthenticationType(value *string)() {
    m.authenticationType = value
}
func (m *Domain) SetAvailabilityStatus(value *string)() {
    m.availabilityStatus = value
}
func (m *Domain) SetDomainNameReferences(value []DirectoryObject)() {
    m.domainNameReferences = value
}
func (m *Domain) SetIsAdminManaged(value *bool)() {
    m.isAdminManaged = value
}
func (m *Domain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *Domain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
func (m *Domain) SetIsRoot(value *bool)() {
    m.isRoot = value
}
func (m *Domain) SetIsVerified(value *bool)() {
    m.isVerified = value
}
func (m *Domain) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *Domain) SetModel(value *string)() {
    m.model = value
}
func (m *Domain) SetPasswordNotificationWindowInDays(value *int32)() {
    m.passwordNotificationWindowInDays = value
}
func (m *Domain) SetPasswordValidityPeriodInDays(value *int32)() {
    m.passwordValidityPeriodInDays = value
}
func (m *Domain) SetServiceConfigurationRecords(value []DomainDnsRecord)() {
    m.serviceConfigurationRecords = value
}
func (m *Domain) SetState(value *DomainState)() {
    m.state = value
}
func (m *Domain) SetSupportedServices(value []string)() {
    m.supportedServices = value
}
func (m *Domain) SetVerificationDnsRecords(value []DomainDnsRecord)() {
    m.verificationDnsRecords = value
}
