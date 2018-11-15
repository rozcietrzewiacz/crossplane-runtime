/*
Copyright 2018 The Conductor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GoogleBucketSpec defines the desired state of S3Bucket
type GoogleBucketSpec struct {
	Name string `json:"name,omitempty"`

	// Location - See authoritative list https://developers.google.com/storage/docs/bucket-locations
	// Which you use is dependent on whether it's multi_region or not.
	Location string `json:"region,omitempty"`

	// PredefinedACL is OneOf - private, authenticatedRead, projectPrivate, publicRead, publicReadWrite
	PredefinedACL *string `json:"cannedACL,omitempty"`

	//StorageClass one of MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when the bucket is created, it will default to STANDARD
	StorageClass *string `json:"storageClass,omitempty"`
	Versioning   bool    `json:"versioning,omitempty"`
}

// S3BucketStatus defines the observed state of RDSInstance
type S3BucketStatus struct {
	State      string `json:"state,omitempty"`
	Message    string `json:"message,omitempty"`
	ProviderID string `json:"providerID,omitempty"` // the external ID to identify this resource in the cloud provider
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBucket is the Schema for the instances API
// +k8s:openapi-gen=true
type GoogleBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleBucketSpec   `json:"spec,omitempty"`
	Status S3BucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBucketList contains a list of S3Buckets
type GoogleBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GoogleBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GoogleBucket{}, &GoogleBucketList{})
}
