package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackage provides operations to manage the admin singleton.
type AccessPackage struct {
    Entity
    // The assignmentPolicies property
    assignmentPolicies []AccessPackageAssignmentPolicyable
    // The catalog property
    catalog AccessPackageCatalogable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the access package.
    description *string
    // The display name of the access package. Supports $filter (eq, contains).
    displayName *string
    // Whether the access package is hidden from the requestor.
    isHidden *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAccessPackage instantiates a new accessPackage and sets the default values.
func NewAccessPackage()(*AccessPackage) {
    m := &AccessPackage{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackage(), nil
}
// GetAssignmentPolicies gets the assignmentPolicies property value. The assignmentPolicies property
func (m *AccessPackage) GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable) {
    return m.assignmentPolicies
}
// GetCatalog gets the catalog property value. The catalog property
func (m *AccessPackage) GetCatalog()(AccessPackageCatalogable) {
    return m.catalog
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the access package.
func (m *AccessPackage) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
func (m *AccessPackage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentPolicyable)
            }
            m.SetAssignmentPolicies(res)
        }
        return nil
    }
    res["catalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(AccessPackageCatalogable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetIsHidden gets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) GetIsHidden()(*bool) {
    return m.isHidden
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// Serialize serializes information the current object
func (m *AccessPackage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentPolicies()))
        for i, v := range m.GetAssignmentPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentPolicies", cast)
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
        err = writer.WriteBoolValue("isHidden", m.GetIsHidden())
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
    return nil
}
// SetAssignmentPolicies sets the assignmentPolicies property value. The assignmentPolicies property
func (m *AccessPackage) SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)() {
    m.assignmentPolicies = value
}
// SetCatalog sets the catalog property value. The catalog property
func (m *AccessPackage) SetCatalog(value AccessPackageCatalogable)() {
    m.catalog = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the access package.
func (m *AccessPackage) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the access package. Supports $filter (eq, contains).
func (m *AccessPackage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsHidden sets the isHidden property value. Whether the access package is hidden from the requestor.
func (m *AccessPackage) SetIsHidden(value *bool)() {
    m.isHidden = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackage) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
