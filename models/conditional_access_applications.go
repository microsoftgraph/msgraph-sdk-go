package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessApplications 
type ConditionalAccessApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of application IDs explicitly excluded from the policy.
    excludeApplications []string;
    // The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
    includeApplications []string;
    // Authentication context class references include. Supported values are c1 through c25.
    includeAuthenticationContextClassReferences []string;
    // User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
    includeUserActions []string;
}
// NewConditionalAccessApplications instantiates a new conditionalAccessApplications and sets the default values.
func NewConditionalAccessApplications()(*ConditionalAccessApplications) {
    m := &ConditionalAccessApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessApplications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeApplications gets the excludeApplications property value. The list of application IDs explicitly excluded from the policy.
func (m *ConditionalAccessApplications) GetExcludeApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeApplications
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessApplications) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeApplications"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeApplications(res)
        }
        return nil
    }
    res["includeApplications"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeApplications(res)
        }
        return nil
    }
    res["includeAuthenticationContextClassReferences"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeAuthenticationContextClassReferences(res)
        }
        return nil
    }
    res["includeUserActions"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeUserActions(res)
        }
        return nil
    }
    return res
}
// GetIncludeApplications gets the includeApplications property value. The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
func (m *ConditionalAccessApplications) GetIncludeApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeApplications
    }
}
// GetIncludeAuthenticationContextClassReferences gets the includeAuthenticationContextClassReferences property value. Authentication context class references include. Supported values are c1 through c25.
func (m *ConditionalAccessApplications) GetIncludeAuthenticationContextClassReferences()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeAuthenticationContextClassReferences
    }
}
// GetIncludeUserActions gets the includeUserActions property value. User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
func (m *ConditionalAccessApplications) GetIncludeUserActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeUserActions
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessApplications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeApplications() != nil {
        err := writer.WriteCollectionOfStringValues("excludeApplications", m.GetExcludeApplications())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeApplications() != nil {
        err := writer.WriteCollectionOfStringValues("includeApplications", m.GetIncludeApplications())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeAuthenticationContextClassReferences() != nil {
        err := writer.WriteCollectionOfStringValues("includeAuthenticationContextClassReferences", m.GetIncludeAuthenticationContextClassReferences())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeUserActions() != nil {
        err := writer.WriteCollectionOfStringValues("includeUserActions", m.GetIncludeUserActions())
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
func (m *ConditionalAccessApplications) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeApplications sets the excludeApplications property value. The list of application IDs explicitly excluded from the policy.
func (m *ConditionalAccessApplications) SetExcludeApplications(value []string)() {
    if m != nil {
        m.excludeApplications = value
    }
}
// SetIncludeApplications sets the includeApplications property value. The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
func (m *ConditionalAccessApplications) SetIncludeApplications(value []string)() {
    if m != nil {
        m.includeApplications = value
    }
}
// SetIncludeAuthenticationContextClassReferences sets the includeAuthenticationContextClassReferences property value. Authentication context class references include. Supported values are c1 through c25.
func (m *ConditionalAccessApplications) SetIncludeAuthenticationContextClassReferences(value []string)() {
    if m != nil {
        m.includeAuthenticationContextClassReferences = value
    }
}
// SetIncludeUserActions sets the includeUserActions property value. User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
func (m *ConditionalAccessApplications) SetIncludeUserActions(value []string)() {
    if m != nil {
        m.includeUserActions = value
    }
}
