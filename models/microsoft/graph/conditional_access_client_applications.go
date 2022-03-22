package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessClientApplications 
type ConditionalAccessClientApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Service principal IDs excluded from the policy scope.
    excludeServicePrincipals []string;
    // Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
    includeServicePrincipals []string;
}
// NewConditionalAccessClientApplications instantiates a new conditionalAccessClientApplications and sets the default values.
func NewConditionalAccessClientApplications()(*ConditionalAccessClientApplications) {
    m := &ConditionalAccessClientApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessClientApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessClientApplicationsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConditionalAccessClientApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessClientApplications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeServicePrincipals gets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
func (m *ConditionalAccessClientApplications) GetExcludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeServicePrincipals
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessClientApplications) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeServicePrincipals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeServicePrincipals(res)
        }
        return nil
    }
    res["includeServicePrincipals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeServicePrincipals(res)
        }
        return nil
    }
    return res
}
// GetIncludeServicePrincipals gets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
func (m *ConditionalAccessClientApplications) GetIncludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeServicePrincipals
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessClientApplications) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExcludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("excludeServicePrincipals", m.GetExcludeServicePrincipals())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("includeServicePrincipals", m.GetIncludeServicePrincipals())
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
func (m *ConditionalAccessClientApplications) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeServicePrincipals sets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
func (m *ConditionalAccessClientApplications) SetExcludeServicePrincipals(value []string)() {
    if m != nil {
        m.excludeServicePrincipals = value
    }
}
// SetIncludeServicePrincipals sets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
func (m *ConditionalAccessClientApplications) SetIncludeServicePrincipals(value []string)() {
    if m != nil {
        m.includeServicePrincipals = value
    }
}
