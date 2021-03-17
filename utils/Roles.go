package utils

import(
	"errors"
	"strings"
)

type Role struct {
	Name   string
	Permissions []string
}

//Here you can set the permissions By Role
var Roles = []Role{
	Role {
		Name : "User",
		Permissions: []string{"BasicUserPermission"},
	},
	Role {
		Name : "AdminMaster",
		Permissions: []string{"BasicUserPermission"},
	},
	Role {
		Name : "Mentor",
		Permissions: []string{"BasicUserPermission","validateSkill", "createSkill", "removeSkill"},
	},
	Role {
		Name : "TL",
		Permissions: []string{"BasicUserPermission"},
	},
	Role {
		Name : "PM",
		Permissions: []string{"BasicUserPermission"},
	},
	Role {
		Name : "HeadOfTechnology",
		Permissions: []string{"BasicUserPermission"},
	},
}

//GetPermissions returns the permissions for a given Role
//The parameter in this function is not case sensitive
func GetPermissions(name string) []string{
	ret := []string{""}
	for i:=0; i< len(Roles); i++{
		if  strings.EqualFold(name, Roles[i].Name) {
			ret = Roles[i].Permissions
			break
		}
	}
	return ret
}

//IsAllowed returns true if a Role has the wanted permission
//The parameters in this function are not case sensitive
func IsAllowed(role string, permission string) (bool, error){
	if strings.EqualFold(role, "admin"){
		return true, nil
	}else{
		per:= GetPermissions(role)	
		for i:=0; i< len(per); i++{
			if  strings.EqualFold(permission, per[i]){
				return true, nil
			}
		}		
	}
	return false, errors.New("Action not allowed")
}