package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReviewReviewerScope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The query specifying who will be the reviewer. See table for examples.
    query *string;
    // In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
    queryRoot *string;
    // The type of query. Examples include MicrosoftGraph and ARM.
    queryType *string;
}
// Instantiates a new accessReviewReviewerScope and sets the default values.
func NewAccessReviewReviewerScope()(*AccessReviewReviewerScope) {
    m := &AccessReviewReviewerScope{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewReviewerScope) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the query property value. The query specifying who will be the reviewer. See table for examples.
func (m *AccessReviewReviewerScope) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
// Gets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
func (m *AccessReviewReviewerScope) GetQueryRoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryRoot
    }
}
// Gets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
func (m *AccessReviewReviewerScope) GetQueryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryType
    }
}
// The deserialization information for the current model
func (m *AccessReviewReviewerScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQuery(val)
        return nil
    }
    res["queryRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryRoot(val)
        return nil
    }
    res["queryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryType(val)
        return nil
    }
    return res
}
func (m *AccessReviewReviewerScope) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReviewReviewerScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryRoot", m.GetQueryRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryType", m.GetQueryType())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AccessReviewReviewerScope) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the query property value. The query specifying who will be the reviewer. See table for examples.
// Parameters:
//  - value : Value to set for the query property.
func (m *AccessReviewReviewerScope) SetQuery(value *string)() {
    m.query = value
}
// Sets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
// Parameters:
//  - value : Value to set for the queryRoot property.
func (m *AccessReviewReviewerScope) SetQueryRoot(value *string)() {
    m.queryRoot = value
}
// Sets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
// Parameters:
//  - value : Value to set for the queryType property.
func (m *AccessReviewReviewerScope) SetQueryType(value *string)() {
    m.queryType = value
}
