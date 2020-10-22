package v1helpers

import (
	configv1 "github.com/openshift/api/config/v1"
)

// AddRelatedObject inserts the provided object into the slice of relatedObjects.
func AddRelatedObject(objects []configv1.ObjectReference, newObj configv1.ObjectReference) []configv1.ObjectReference {
	if objects == nil {
		objects = []configv1.ObjectReference{}
	}
	objects = append(objects, newObj)
	return objects
}

// FindRelatedObject returns the ObjectReference that matches the provided name and namespace
func FindRelatedObject(objects []configv1.ObjectReference, name, namespace string) *configv1.ObjectReference {
	for i := range objects {
		if objects[i].Name == name && objects[i].Namespace == namespace {
			return &objects[i]
		}
	}
	return nil
}

// RemoveRelatedObject removes the provided object from the list of relatedObjects, and then
// returns a new slice with the results along with whether it was able to remove the object or not.
func RemoveRelatedObject(objects *[]configv1.ObjectReference, obj configv1.ObjectReference) {
	if objects == nil {
		objects = &[]configv1.ObjectReference{}
	}
	newObjects := []configv1.ObjectReference{}
	for _, value := range *objects {
		if !objectIsEqual(value, obj) {
			newObjects = append(newObjects, value)
		}
	}

	*objects = newObjects
}

func objectIsEqual(oldObj, newObj configv1.ObjectReference) bool {
	if oldObj.Group == newObj.Group &&
		oldObj.Resource == newObj.Resource &&
		oldObj.Namespace == newObj.Namespace &&
		oldObj.Name == newObj.Name {
		return true
	}
	return false
}
