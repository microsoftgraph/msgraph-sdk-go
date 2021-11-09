package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type DataSourceContainer struct {
    Entity
    // Created date and time of the dataSourceContainer entity.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Display name of the dataSourceContainer entity.
    displayName *string;
    // 
    lastIndexOperation *CaseIndexOperation;
    // Last modified date and time of the dataSourceContainer.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date and time that the dataSourceContainer was released from the case.
    releasedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Latest status of the dataSourceContainer. Possible values are: Active, Released.
    status *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceContainerStatus;
}
// Instantiates a new dataSourceContainer and sets the default values.
func NewDataSourceContainer()(*DataSourceContainer) {
    m := &DataSourceContainer{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastIndexOperation property value. 
func (m *DataSourceContainer) GetLastIndexOperation()(*CaseIndexOperation) {
    if m == nil {
        return nil
    } else {
        return m.lastIndexOperation
    }
}
// Gets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) GetReleasedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.releasedDateTime
    }
}
// Gets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) GetStatus()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceContainerStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *DataSourceContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastIndexOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseIndexOperation() })
        if err != nil {
            return err
        }
        m.SetLastIndexOperation(val.(*CaseIndexOperation))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["releasedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReleasedDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseDataSourceContainerStatus)
        if err != nil {
            return err
        }
        cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceContainerStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *DataSourceContainer) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DataSourceContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    {
        err = writer.WriteObjectValue("lastIndexOperation", m.GetLastIndexOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releasedDateTime", m.GetReleasedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DataSourceContainer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. Display name of the dataSourceContainer entity.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DataSourceContainer) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastIndexOperation property value. 
// Parameters:
//  - value : Value to set for the lastIndexOperation property.
func (m *DataSourceContainer) SetLastIndexOperation(value *CaseIndexOperation)() {
    m.lastIndexOperation = value
}
// Sets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DataSourceContainer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
// Parameters:
//  - value : Value to set for the releasedDateTime property.
func (m *DataSourceContainer) SetReleasedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releasedDateTime = value
}
// Sets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
// Parameters:
//  - value : Value to set for the status property.
func (m *DataSourceContainer) SetStatus(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceContainerStatus)() {
    m.status = value
}
