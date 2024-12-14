package bolt

import (
	"github.com/semaphoreui/semaphore/db"
)

func (d *BoltDb) CreateTerraformInventoryAlias(alias db.TerraformInventoryAlias) (res db.TerraformInventoryAlias, err error) {
	var newInv interface{}
	newInv, err = d.createObject(alias.ProjectID, db.TerraformInventoryAliasProps, alias)
	if err != nil {
		return
	}
	res = newInv.(db.TerraformInventoryAlias)
	return
}

func (d *BoltDb) UpdateTerraformInventoryAlias(alias db.TerraformInventoryAlias) (err error) {

	err = d.updateObject(alias.ProjectID, db.TemplateProps, alias)

	return
}

func (d *BoltDb) GetTerraformInventoryAliasByAlias(alias string) (res db.TerraformInventoryAlias, err error) {

	err = d.getObject(-1, db.TerraformInventoryStateProps, strObjectID(alias), &res)

	return
}

func (d *BoltDb) GetTerraformInventoryAlias(projectID, inventoryID int, alias string) (res db.TerraformInventoryAlias, err error) {

	al, err := d.GetTerraformInventoryAliasByAlias(alias)

	if err != nil {
		return
	}

	if al.ProjectID != projectID || al.InventoryID != inventoryID {
		err = db.ErrNotFound
		return
	}

	return
}

func (d *BoltDb) GetTerraformInventoryAliases(projectID, inventoryID int) (res []db.TerraformInventoryAlias, err error) {
	err = d.getObjects(projectID, db.TerraformInventoryAliasProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
		alias := i.(db.TerraformInventoryAlias)
		return alias.InventoryID == inventoryID
	}, &res)
	return
}

func (d *BoltDb) DeleteTerraformInventoryAlias(projectID int, inventoryID int, alias string) (err error) {
	err = d.deleteObject(projectID, db.TerraformInventoryAliasProps, strObjectID(alias), nil)
	return
}

func (d *BoltDb) GetTerraformInventoryStates(projectID, inventoryID int) (res []db.TerraformInventoryState, err error) {
	err = d.getObjects(projectID, db.TerraformInventoryStateProps, db.RetrieveQueryParams{}, func(i interface{}) bool {
		state := i.(db.TerraformInventoryState)
		return state.InventoryID == inventoryID
	}, &res)
	return
}
