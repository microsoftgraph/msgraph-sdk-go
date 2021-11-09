package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AddToReviewSetOperation struct {
    CaseOperation
    // The review set to which items matching the source collection query are added to.
    reviewSet *ReviewSet;
    // The sourceCollection that items are being added from.
    sourceCollection *SourceCollection;
}
// Instantiates a new addToReviewSetOperation and sets the default values.
func NewAddToReviewSetOperation()(*AddToReviewSetOperation) {
    m := &AddToReviewSetOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// Gets the reviewSet property value. The review set to which items matching the source collection query are added to.
func (m *AddToReviewSetOperation) GetReviewSet()(*ReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.reviewSet
    }
}
// Gets the sourceCollection property value. The sourceCollection that items are being added from.
func (m *AddToReviewSetOperation) GetSourceCollection()(*SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollection
    }
}
// The deserialization information for the current model
func (m *AddToReviewSetOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["reviewSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReviewSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(*ReviewSet))
        }
        return nil
    }
    res["sourceCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(*SourceCollection))
        }
        return nil
    }
    return res
}
func (m *AddToReviewSetOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AddToReviewSetOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the reviewSet property value. The review set to which items matching the source collection query are added to.
// Parameters:
//  - value : Value to set for the reviewSet property.
func (m *AddToReviewSetOperation) SetReviewSet(value *ReviewSet)() {
    m.reviewSet = value
}
// Sets the sourceCollection property value. The sourceCollection that items are being added from.
// Parameters:
//  - value : Value to set for the sourceCollection property.
func (m *AddToReviewSetOperation) SetSourceCollection(value *SourceCollection)() {
    m.sourceCollection = value
}
