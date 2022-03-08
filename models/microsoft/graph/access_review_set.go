package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewSet provides operations to manage the identityGovernance singleton.
type AccessReviewSet struct {
    Entity
    // Represents the template and scheduling for an access review.
    definitions []AccessReviewScheduleDefinitionable;
    // Represents a collection of access review history data and the scopes used to collect that data.
    historyDefinitions []AccessReviewHistoryDefinitionable;
}
// NewAccessReviewSet instantiates a new accessReviewSet and sets the default values.
func NewAccessReviewSet()(*AccessReviewSet) {
    m := &AccessReviewSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessReviewSet(), nil
}
// GetDefinitions gets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) GetDefinitions()([]AccessReviewScheduleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScheduleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScheduleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewScheduleDefinitionable)
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    res["historyDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewHistoryDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewHistoryDefinitionable)
            }
            m.SetHistoryDefinitions(res)
        }
        return nil
    }
    return res
}
// GetHistoryDefinitions gets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) GetHistoryDefinitions()([]AccessReviewHistoryDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.historyDefinitions
    }
}
func (m *AccessReviewSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistoryDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistoryDefinitions()))
        for i, v := range m.GetHistoryDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("historyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitions sets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) SetDefinitions(value []AccessReviewScheduleDefinitionable)() {
    if m != nil {
        m.definitions = value
    }
}
// SetHistoryDefinitions sets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)() {
    if m != nil {
        m.historyDefinitions = value
    }
}
