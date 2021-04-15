/*
Copyright 2021.

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

package conversion

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ZapAnnotation is the annotation representation of Zap for specs that don't
	// contain this value.
	ZapAnnotation = "conversion.mycustomkind.demo.example.com/zap"
	// ZapAnnotation is the annotation representation of Zap for specs that don't
	// contain this value.
	BarAnnotation = "conversion.mycustomkind.demo.example.com/bar"
)

// GetExistingOrEmptyAnnotations grabs annotations from an object's metadata
// and if the object has no annotations (represented by nil), then we create a
// a new map and return it instead.
func GetExistingOrEmptyAnnotations(metadata metav1.ObjectMeta) map[string]string {
	annotations := metadata.GetAnnotations()

	if annotations == nil {
		annotations = make(map[string]string)
	}

	return annotations
}

// GetZarAnnotationValue returns the value of the bar annotation from the annotations
// map. Error will be returned if the key did not exist.
func GetZapAnnotationValue(annotations map[string]string) (string, bool) {
	return getAnnotationValueIfExists(ZapAnnotation, annotations)
}

// GetBarAnnotationValue returns the value of the bar annotation from the annotations
// map. Error will be returned if the key did not exist.
func GetBarAnnotationValue(annotations map[string]string) (string, bool) {
	return getAnnotationValueIfExists(BarAnnotation, annotations)
}

func getAnnotationValueIfExists(key string, annotations map[string]string) (string, bool) {
	var value string
	if key == "" {
		// empty keys are a no-op
		return "", false
	}

	if annotations == nil {
		// empty annotations is a no-op
		return "", false
	}

	value, exists := annotations[key]
	if !exists {
		return value, false
	}

	return value, true
}
