package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SchedulingGroup 
type SchedulingGroup struct {
    ChangeTrackedEntity
    // The display name for the schedulingGroup. Required.
    displayName *string;
    // Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
    isActive *bool;
    // The list of user IDs that are a member of the schedulingGroup. Required.
    userIds []string;
}
// NewSchedulingGroup instantiates a new schedulingGroup and sets the default values.
func NewSchedulingGroup()(*SchedulingGroup) {
    m := &SchedulingGroup{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// CreateSchedulingGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSchedulingGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSchedulingGroup(), nil
}
// GetDisplayName gets the displayName property value. The display name for the schedulingGroup. Required.
func (m *SchedulingGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SchedulingGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
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
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["userIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserIds(res)
        }
        return nil
    }
    return res
}
// GetIsActive gets the isActive property value. Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
func (m *SchedulingGroup) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetUserIds gets the userIds property value. The list of user IDs that are a member of the schedulingGroup. Required.
func (m *SchedulingGroup) GetUserIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userIds
    }
}
// Serialize serializes information the current object
func (m *SchedulingGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    if m.GetUserIds() != nil {
        err = writer.WriteCollectionOfStringValues("userIds", m.GetUserIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the schedulingGroup. Required.
func (m *SchedulingGroup) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsActive sets the isActive property value. Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
func (m *SchedulingGroup) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetUserIds sets the userIds property value. The list of user IDs that are a member of the schedulingGroup. Required.
func (m *SchedulingGroup) SetUserIds(value []string)() {
    if m != nil {
        m.userIds = value
    }
}
