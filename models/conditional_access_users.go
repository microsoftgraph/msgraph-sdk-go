package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessUsers 
type ConditionalAccessUsers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Group IDs excluded from scope of policy.
    excludeGroups []string;
    // Role IDs excluded from scope of policy.
    excludeRoles []string;
    // User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
    excludeUsers []string;
    // Group IDs in scope of policy unless explicitly excluded, or All.
    includeGroups []string;
    // Role IDs in scope of policy unless explicitly excluded, or All.
    includeRoles []string;
    // User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
    includeUsers []string;
}
// NewConditionalAccessUsers instantiates a new conditionalAccessUsers and sets the default values.
func NewConditionalAccessUsers()(*ConditionalAccessUsers) {
    m := &ConditionalAccessUsers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessUsersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessUsersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessUsers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessUsers) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeGroups gets the excludeGroups property value. Group IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) GetExcludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroups
    }
}
// GetExcludeRoles gets the excludeRoles property value. Role IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) GetExcludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeRoles
    }
}
// GetExcludeUsers gets the excludeUsers property value. User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) GetExcludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessUsers) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeGroups"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["excludeRoles"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeRoles(res)
        }
        return nil
    }
    res["excludeUsers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeUsers(res)
        }
        return nil
    }
    res["includeGroups"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeGroups(res)
        }
        return nil
    }
    res["includeRoles"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeRoles(res)
        }
        return nil
    }
    res["includeUsers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeUsers(res)
        }
        return nil
    }
    return res
}
// GetIncludeGroups gets the includeGroups property value. Group IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) GetIncludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeGroups
    }
}
// GetIncludeRoles gets the includeRoles property value. Role IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) GetIncludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeRoles
    }
}
// GetIncludeUsers gets the includeUsers property value. User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) GetIncludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeUsers
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessUsers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeRoles() != nil {
        err := writer.WriteCollectionOfStringValues("excludeRoles", m.GetExcludeRoles())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeUsers() != nil {
        err := writer.WriteCollectionOfStringValues("excludeUsers", m.GetExcludeUsers())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeRoles() != nil {
        err := writer.WriteCollectionOfStringValues("includeRoles", m.GetIncludeRoles())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeUsers() != nil {
        err := writer.WriteCollectionOfStringValues("includeUsers", m.GetIncludeUsers())
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
func (m *ConditionalAccessUsers) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeGroups sets the excludeGroups property value. Group IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) SetExcludeGroups(value []string)() {
    if m != nil {
        m.excludeGroups = value
    }
}
// SetExcludeRoles sets the excludeRoles property value. Role IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) SetExcludeRoles(value []string)() {
    if m != nil {
        m.excludeRoles = value
    }
}
// SetExcludeUsers sets the excludeUsers property value. User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) SetExcludeUsers(value []string)() {
    if m != nil {
        m.excludeUsers = value
    }
}
// SetIncludeGroups sets the includeGroups property value. Group IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) SetIncludeGroups(value []string)() {
    if m != nil {
        m.includeGroups = value
    }
}
// SetIncludeRoles sets the includeRoles property value. Role IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) SetIncludeRoles(value []string)() {
    if m != nil {
        m.includeRoles = value
    }
}
// SetIncludeUsers sets the includeUsers property value. User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) SetIncludeUsers(value []string)() {
    if m != nil {
        m.includeUsers = value
    }
}
