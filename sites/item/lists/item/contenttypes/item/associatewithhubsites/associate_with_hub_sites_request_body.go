package associatewithhubsites

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssociateWithHubSitesRequestBody provides operations to call the associateWithHubSites method.
type AssociateWithHubSitesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hubSiteUrls []string;
    // 
    propagateToExistingLists *bool;
}
// NewAssociateWithHubSitesRequestBody instantiates a new associateWithHubSitesRequestBody and sets the default values.
func NewAssociateWithHubSitesRequestBody()(*AssociateWithHubSitesRequestBody) {
    m := &AssociateWithHubSitesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssociateWithHubSitesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssociateWithHubSitesRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAssociateWithHubSitesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssociateWithHubSitesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssociateWithHubSitesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hubSiteUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHubSiteUrls(res)
        }
        return nil
    }
    res["propagateToExistingLists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateToExistingLists(val)
        }
        return nil
    }
    return res
}
// GetHubSiteUrls gets the hubSiteUrls property value. 
func (m *AssociateWithHubSitesRequestBody) GetHubSiteUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hubSiteUrls
    }
}
// GetPropagateToExistingLists gets the propagateToExistingLists property value. 
func (m *AssociateWithHubSitesRequestBody) GetPropagateToExistingLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateToExistingLists
    }
}
func (m *AssociateWithHubSitesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssociateWithHubSitesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHubSiteUrls() != nil {
        err := writer.WriteCollectionOfStringValues("hubSiteUrls", m.GetHubSiteUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateToExistingLists", m.GetPropagateToExistingLists())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssociateWithHubSitesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHubSiteUrls sets the hubSiteUrls property value. 
func (m *AssociateWithHubSitesRequestBody) SetHubSiteUrls(value []string)() {
    if m != nil {
        m.hubSiteUrls = value
    }
}
// SetPropagateToExistingLists sets the propagateToExistingLists property value. 
func (m *AssociateWithHubSitesRequestBody) SetPropagateToExistingLists(value *bool)() {
    if m != nil {
        m.propagateToExistingLists = value
    }
}
