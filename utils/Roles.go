package utils

import (
	"errors"
)

//Here you can set the permissions By Role
/*var Roles = models.Role{
	Permissions: []string{"BasicUserPermission", "Mentor", "AdminMaster", "TL"},
}*/

//GetPermissions returns the permissions for a given Role
//The parameter in this function is not case sensitive
/*func GetPermissions(name string) ([]string,int) {
	ret := []string{""}
	var aux int
	for i := 0; i < len(Roles.Permissions); i++ {
		if strings.EqualFold(name, Roles.Permissions[i]) {
			ret = Roles.Permissions
			aux=i
			break
		}
	}
	return ret,aux
}*/

//IsAllowed returns true if a Role has the wanted permission
//The parameters in this function are not case sensitive
func IsAllowed(role string) (bool, error) {
	if role == "AdminMaster" {
		return true, nil
	}
	return false, errors.New("Action not allowed")
}
