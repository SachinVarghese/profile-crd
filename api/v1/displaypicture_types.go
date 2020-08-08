/*

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DisplayPictureSpec defines the desired state of DisplayPicture
type DisplayPictureSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Src is an image source field of DisplayPicture. Edit DisplayPicture_types.go to remove/update
	Src string `json:"src,omitempty"`
}

// DisplayPictureStatus defines the observed state of DisplayPicture
type DisplayPictureStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DisplayPicture is the Schema for the displaypictures API
type DisplayPicture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DisplayPictureSpec   `json:"spec,omitempty"`
	Status DisplayPictureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DisplayPictureList contains a list of DisplayPicture
type DisplayPictureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DisplayPicture `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DisplayPicture{}, &DisplayPictureList{})
}
