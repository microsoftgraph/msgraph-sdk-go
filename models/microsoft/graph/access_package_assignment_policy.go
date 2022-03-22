package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentPolicy 
type AccessPackageAssignmentPolicy struct {
    Entity
    // The access package with this policy. Read-only. Nullable. Supports $expand.
    accessPackage AccessPackageable;
    // 
    allowedTargetScope *AllowedTargetScope;
    // 
    catalog AccessPackageCatalogable;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the policy.
    description *string;
    // The display name of the policy. Supports $filter (eq).
    displayName *string;
    // 
    expiration ExpirationPatternable;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Who must approve requests for access package in this policy.
    requestApprovalSettings AccessPackageAssignmentApprovalSettingsable;
    // Who can request this access package from this policy.
    requestorSettings AccessPackageAssignmentRequestorSettingsable;
    // 
    reviewSettings AccessPackageAssignmentReviewSettingsable;
    // 
    specificAllowedTargets []SubjectSetable;
}
// NewAccessPackageAssignmentPolicy instantiates a new accessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentPolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageAssignmentPolicy(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(AccessPackageable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAllowedTargetScope gets the allowedTargetScope property value. 
func (m *AccessPackageAssignmentPolicy) GetAllowedTargetScope()(*AllowedTargetScope) {
    if m == nil {
        return nil
    } else {
        return m.allowedTargetScope
    }
}
// GetCatalog gets the catalog property value. 
func (m *AccessPackageAssignmentPolicy) GetCatalog()(AccessPackageCatalogable) {
    if m == nil {
        return nil
    } else {
        return m.catalog
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExpiration gets the expiration property value. 
func (m *AccessPackageAssignmentPolicy) GetExpiration()(ExpirationPatternable) {
    if m == nil {
        return nil
    } else {
        return m.expiration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["allowedTargetScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedTargetScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedTargetScope(val.(*AllowedTargetScope))
        }
        return nil
    }
    res["catalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(AccessPackageCatalogable))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["expiration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpirationPatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val.(ExpirationPatternable))
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["requestApprovalSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestApprovalSettings(val.(AccessPackageAssignmentApprovalSettingsable))
        }
        return nil
    }
    res["requestorSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorSettings(val.(AccessPackageAssignmentRequestorSettingsable))
        }
        return nil
    }
    res["reviewSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSettings(val.(AccessPackageAssignmentReviewSettingsable))
        }
        return nil
    }
    res["specificAllowedTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetSpecificAllowedTargets(res)
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetRequestApprovalSettings gets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(AccessPackageAssignmentApprovalSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.requestApprovalSettings
    }
}
// GetRequestorSettings gets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(AccessPackageAssignmentRequestorSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.requestorSettings
    }
}
// GetReviewSettings gets the reviewSettings property value. 
func (m *AccessPackageAssignmentPolicy) GetReviewSettings()(AccessPackageAssignmentReviewSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.reviewSettings
    }
}
// GetSpecificAllowedTargets gets the specificAllowedTargets property value. 
func (m *AccessPackageAssignmentPolicy) GetSpecificAllowedTargets()([]SubjectSetable) {
    if m == nil {
        return nil
    } else {
        return m.specificAllowedTargets
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedTargetScope() != nil {
        cast := (*m.GetAllowedTargetScope()).String()
        err = writer.WriteStringValue("allowedTargetScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("catalog", m.GetCatalog())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestApprovalSettings", m.GetRequestApprovalSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestorSettings", m.GetRequestorSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSettings", m.GetReviewSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSpecificAllowedTargets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSpecificAllowedTargets()))
        for i, v := range m.GetSpecificAllowedTargets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("specificAllowedTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package with this policy. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value AccessPackageable)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAllowedTargetScope sets the allowedTargetScope property value. 
func (m *AccessPackageAssignmentPolicy) SetAllowedTargetScope(value *AllowedTargetScope)() {
    if m != nil {
        m.allowedTargetScope = value
    }
}
// SetCatalog sets the catalog property value. 
func (m *AccessPackageAssignmentPolicy) SetCatalog(value AccessPackageCatalogable)() {
    if m != nil {
        m.catalog = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the policy. Supports $filter (eq).
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExpiration sets the expiration property value. 
func (m *AccessPackageAssignmentPolicy) SetExpiration(value ExpirationPatternable)() {
    if m != nil {
        m.expiration = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
// SetRequestApprovalSettings sets the requestApprovalSettings property value. Who must approve requests for access package in this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value AccessPackageAssignmentApprovalSettingsable)() {
    if m != nil {
        m.requestApprovalSettings = value
    }
}
// SetRequestorSettings sets the requestorSettings property value. Who can request this access package from this policy.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value AccessPackageAssignmentRequestorSettingsable)() {
    if m != nil {
        m.requestorSettings = value
    }
}
// SetReviewSettings sets the reviewSettings property value. 
func (m *AccessPackageAssignmentPolicy) SetReviewSettings(value AccessPackageAssignmentReviewSettingsable)() {
    if m != nil {
        m.reviewSettings = value
    }
}
// SetSpecificAllowedTargets sets the specificAllowedTargets property value. 
func (m *AccessPackageAssignmentPolicy) SetSpecificAllowedTargets(value []SubjectSetable)() {
    if m != nil {
        m.specificAllowedTargets = value
    }
}
