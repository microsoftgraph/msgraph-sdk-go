package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Contract provides operations to manage the collection of contract entities.
type Contract struct {
    DirectoryObject
    // Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below.
    contractType *string;
    // The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant's organization resource.
    customerId *string;
    // A copy of the customer tenant's default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's default domain name changes.
    defaultDomainName *string;
    // A copy of the customer tenant's display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's display name changes.
    displayName *string;
}
// NewContract instantiates a new contract and sets the default values.
func NewContract()(*Contract) {
    m := &Contract{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// CreateContractFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContractFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewContract(), nil
}
// GetContractType gets the contractType property value. Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below.
func (m *Contract) GetContractType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contractType
    }
}
// GetCustomerId gets the customerId property value. The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant's organization resource.
func (m *Contract) GetCustomerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customerId
    }
}
// GetDefaultDomainName gets the defaultDomainName property value. A copy of the customer tenant's default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's default domain name changes.
func (m *Contract) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
// GetDisplayName gets the displayName property value. A copy of the customer tenant's display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's display name changes.
func (m *Contract) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Contract) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["contractType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContractType(val)
        }
        return nil
    }
    res["customerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["defaultDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
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
    return res
}
func (m *Contract) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Contract) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contractType", m.GetContractType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
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
    return nil
}
// SetContractType sets the contractType property value. Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below.
func (m *Contract) SetContractType(value *string)() {
    if m != nil {
        m.contractType = value
    }
}
// SetCustomerId sets the customerId property value. The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant's organization resource.
func (m *Contract) SetCustomerId(value *string)() {
    if m != nil {
        m.customerId = value
    }
}
// SetDefaultDomainName sets the defaultDomainName property value. A copy of the customer tenant's default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's default domain name changes.
func (m *Contract) SetDefaultDomainName(value *string)() {
    if m != nil {
        m.defaultDomainName = value
    }
}
// SetDisplayName sets the displayName property value. A copy of the customer tenant's display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's display name changes.
func (m *Contract) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
