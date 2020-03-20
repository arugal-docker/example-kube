package main

import (
	"encoding/json"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Resource struct {
	Metadata   metav1.ObjectMeta `json:"metadata"`
	Kind       string            `json:"kind"`
	ApiVersion string            `json:"api_version"`
}

// Add a label {"added-label": "yes"} to the object
func addLabel(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	log.Infof("calling add-label %v", string(ar.Request.Object.Raw))
	obj := Resource{}
	raw := ar.Request.Object.Raw
	err := json.Unmarshal(raw, &obj)
	if err != nil {
		log.Error(err)
		return toAdmissionResponse(err)
	}

	reviewResponse := v1beta1.AdmissionResponse{}
	reviewResponse.Allowed = true

	log.Infof("Resource: %v", obj)

	var patches []patchOperation
	if len(obj.Metadata.Labels) == 0 {
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
