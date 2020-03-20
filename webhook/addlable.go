package main

import (
	"encoding/json"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add a label {"added-label": "yes"} to the object
func addLabel(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	log.Info("calling add-label %v", string(ar.Request.Object.Raw))
	obj := struct {
		metav1.ObjectMeta
		Data map[string]string
	}{}
	raw := ar.Request.Object.Raw
	err := json.Unmarshal(raw, &obj)
	if err != nil {
		log.Error(err)
		return toAdmissionResponse(err)
	}

	reviewResponse := v1beta1.AdmissionResponse{}
	reviewResponse.Allowed = true

	var patches []patchOperation
	log.Infof("Object: %v", obj)
	log.Infof("Labels: %v", obj.Labels)
	log.Infof("ObjectMeta Labels: %v", obj.Labels)
	if len(obj.ObjectMeta.Labels) == 0 {
		labels := make(map[string]string)
		labels["added-label"] = "first-label"
		patches = append(patches, patchOperation{
			Op:    "add",
			Path:  "/metadata/labels",
			Value: labels,
		})
	} else {
		patches = append(patches, patchOperation{
			Op:    "add",
			Path:  "/metadata/labels/added-label",
			Value: "yes",
		})
	}
	patch, err := json.Marshal(patches)
	if err != nil {
		log.Errorf("patches marshal err: %v", err)
	}
	reviewResponse.Patch = patch
	pt := v1beta1.PatchTypeJSONPatch
	reviewResponse.PatchType = &pt
	return &reviewResponse
}
