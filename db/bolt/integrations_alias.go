package bolt

import (
	"errors"
	"fmt"
	"github.com/semaphoreui/semaphore/db"
	"reflect"
)

var integrationAliasProps = db.ObjectProps{
	TableName:         "integration_alias",
	Type:              reflect.TypeOf(db.IntegrationAlias{}),
	PrimaryColumnName: "alias",
}

func (d *BoltDb) GetIntegrationAliases(projectID int, integrationID *int) (res []db.IntegrationAlias, err error) {

	err = d.getObjects(projectID, db.IntegrationAliasProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
		alias := i.(db.IntegrationAlias)
		if alias.IntegrationID == nil && integrationID == nil {
			return true
		} else if alias.IntegrationID != nil && integrationID != nil {
			return *alias.IntegrationID == *integrationID
		}
		return false
	}, &res)

	return
}

func (d *BoltDb) GetIntegrationsByAlias(alias string) (res []db.Integration, err error) {

	var aliasObj db.IntegrationAlias
	err = d.getObject(-1, integrationAliasProps, strObjectID(alias), &aliasObj)

	if err != nil {
		return
	}

	if aliasObj.IntegrationID == nil {
		err = d.getObjects(aliasObj.ProjectID, db.IntegrationProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
			integration := i.(db.Integration)
			return integration.Searchable
		}, &res)

		if err != nil {
			return
		}

	} else {
		var integration db.Integration
		integration, err = d.GetIntegration(aliasObj.ProjectID, *aliasObj.IntegrationID)
		if err != nil {
			return
		}
		res = append(res, integration)
	}

	return
}

func (d *BoltDb) CreateIntegrationAlias(alias db.IntegrationAlias) (res db.IntegrationAlias, err error) {

	_, err = d.GetIntegrationsByAlias(alias.Alias)

	if err == nil {
		err = fmt.Errorf("alias already exists")
	}

	if !errors.Is(err, db.ErrNotFound) {
		return
	}

	newAlias, err := d.createObject(alias.ProjectID, db.IntegrationAliasProps, alias)

	if err != nil {
		return
	}

	res = newAlias.(db.IntegrationAlias)

	_, err = d.createObject(-1, integrationAliasProps, alias)

	if err != nil {
		_ = d.DeleteIntegrationAlias(alias.ProjectID, alias.ID)
		return
	}

	return
}

func (d *BoltDb) DeleteIntegrationAlias(projectID int, aliasID int) (err error) {

	var alias db.IntegrationAlias
	err = d.getObject(projectID, db.IntegrationAliasProps, intObjectID(aliasID), &alias)
	if err != nil {
		return
	}

	err = d.deleteObject(projectID, db.IntegrationAliasProps, intObjectID(aliasID), nil)
	if err != nil {
		return
	}

	err = d.deleteObject(-1, integrationAliasProps, strObjectID(alias.Alias), nil)
	if err != nil {
		return
	}

	return
}
