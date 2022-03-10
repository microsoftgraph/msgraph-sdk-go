package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Configuration provides operations to manage the collection of externalConnection entities.
type Configuration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of application IDs for registered Azure Active Directory apps that are allowed to manage the externalConnection and to index content in the externalConnection.
    authorizedAppIds []string;
}
// NewConfiguration instantiates a new configuration and sets the default values.
func NewConfiguration()(*Configuration) {
    m := &Configuration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Configuration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthorizedAppIds gets the authorizedAppIds property value. A collection of application IDs for registered Azure Active Directory apps that are allowed to manage the externalConnection and to index content in the externalConnection.
func (m *Configuration) GetAuthorizedAppIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.authorizedAppIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Configuration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authorizedAppIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAuthorizedAppIds(res)
        }
        return nil
    }
    return res
}
func (m *Configuration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Configuration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAuthorizedAppIds() != nil {
        err := writer.WriteCollectionOfStringValues("authorizedAppIds", m.GetAuthorizedAppIds())
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
func (m *Configuration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthorizedAppIds sets the authorizedAppIds property value. A collection of application IDs for registered Azure Active Directory apps that are allowed to manage the externalConnection and to index content in the externalConnection.
func (m *Configuration) SetAuthorizedAppIds(value []string)() {
    if m != nil {
        m.authorizedAppIds = value
    }
}
