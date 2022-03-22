package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageSubject 
type AccessPackageSubject struct {
    Entity
    // The connected organization of the subject. Read-only. Nullable.
    connectedOrganization ConnectedOrganizationable;
    // The display name of the subject.
    displayName *string;
    // The email address of the subject.
    email *string;
    // The object identifier of the subject. null if the subject is not yet a user in the tenant.
    objectId *string;
    // A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
    onPremisesSecurityIdentifier *string;
    // The principal name, if known, of the subject.
    principalName *string;
    // The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
    subjectType *AccessPackageSubjectType;
}
// NewAccessPackageSubject instantiates a new accessPackageSubject and sets the default values.
func NewAccessPackageSubject()(*AccessPackageSubject) {
    m := &AccessPackageSubject{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageSubjectFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessPackageSubject(), nil
}
// GetConnectedOrganization gets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) GetConnectedOrganization()(ConnectedOrganizationable) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganization
    }
}
// GetDisplayName gets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmail gets the email property value. The email address of the subject.
func (m *AccessPackageSubject) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectedOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganization(val.(ConnectedOrganizationable))
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
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["objectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["principalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["subjectType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageSubjectType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectType(val.(*AccessPackageSubjectType))
        }
        return nil
    }
    return res
}
// GetObjectId gets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// GetPrincipalName gets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
// GetSubjectType gets the subjectType property value. The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
func (m *AccessPackageSubject) GetSubjectType()(*AccessPackageSubjectType) {
    if m == nil {
        return nil
    } else {
        return m.subjectType
    }
}
// Serialize serializes information the current object
func (m *AccessPackageSubject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connectedOrganization", m.GetConnectedOrganization())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectType() != nil {
        cast := (*m.GetSubjectType()).String()
        err = writer.WriteStringValue("subjectType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectedOrganization sets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) SetConnectedOrganization(value ConnectedOrganizationable)() {
    if m != nil {
        m.connectedOrganization = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmail sets the email property value. The email address of the subject.
func (m *AccessPackageSubject) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetObjectId sets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) SetObjectId(value *string)() {
    if m != nil {
        m.objectId = value
    }
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) SetOnPremisesSecurityIdentifier(value *string)() {
    if m != nil {
        m.onPremisesSecurityIdentifier = value
    }
}
// SetPrincipalName sets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) SetPrincipalName(value *string)() {
    if m != nil {
        m.principalName = value
    }
}
// SetSubjectType sets the subjectType property value. The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
func (m *AccessPackageSubject) SetSubjectType(value *AccessPackageSubjectType)() {
    if m != nil {
        m.subjectType = value
    }
}
