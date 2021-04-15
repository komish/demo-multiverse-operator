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

package v1beta1

import (
	"strconv"

	apiconversion "github.com/komish/demo-multiverse-operator/api/conversion"

	demov1alpha1 "github.com/komish/demo-multiverse-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

/* Note:
You always need to set the src object meta to the dst in order
for the conversion logic to confirm that the two objects are
the same.

How this is handled could be improved, but for now we deep copy
the source ObjectMeta into the destination, make edits to that,
and then set the dst ObjectMeta as the last action before returning.
*/

// BUG(): This convertible always persists configuration as annotations
// regardless of whether not the key needed to be persisted in the annotations
// due to it being removed from the spec.

// ConvertTo partially implements the conversion.Convertible interface
// indicates that this version is a spoke.
func (src *MyCustomKind) ConvertTo(dstRaw conversion.Hub) error {
	// Hub is v1alpha
	dst := dstRaw.(*demov1alpha1.MyCustomKind) // assert the type because we have the raw
	dstNewObjectMeta := src.DeepCopy().ObjectMeta

	// Handle Foo - renamed in Hub
	dst.Spec.Foo = src.Spec.NewFoo

	// Handle Zap - exists in hub spec but not in src
	zapValStr, found := apiconversion.GetZapAnnotationValue(src.Annotations)
	if found {
		zapValInt64, err := strconv.ParseInt(zapValStr, 10, 32)
		if err != nil {
			return err
		}

		dst.Spec.Zap = int32(zapValInt64)
	}

	// Handle Bar -- exists in dst spec but not in hub
	barValStr := src.Spec.Bar
	if barValStr != "" { // empty check
		dstNewAnnotations := apiconversion.GetExistingOrEmptyAnnotations(dstNewObjectMeta)
		dstNewAnnotations[apiconversion.BarAnnotation] = barValStr
		dstNewObjectMeta.SetAnnotations(dstNewAnnotations)
	}

	dst.ObjectMeta = dstNewObjectMeta

	return nil

}

// ConvertFrom partially implements the conversion.Convertible interface
// indicates that this version is a spoke.
func (dst *MyCustomKind) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*demov1alpha1.MyCustomKind) // assert the type because we have the raw
	dstNewObjectMeta := src.DeepCopy().ObjectMeta

	// Handle Foo -
	dst.Spec.NewFoo = src.Spec.Foo

	// Handle Zap - move to dst annotations from src spec
	zapVal := src.Spec.Zap
	if zapVal != 0 { // empty check
		dstNewAnnotations := apiconversion.GetExistingOrEmptyAnnotations(dstNewObjectMeta)
		dstNewAnnotations[apiconversion.ZapAnnotation] = strconv.FormatInt(int64(zapVal), 10)
		dstNewObjectMeta.SetAnnotations(dstNewAnnotations)
	}

	// Handle Bar - move to dst spec from src annotations
	barVal, found := apiconversion.GetBarAnnotationValue(
		apiconversion.GetExistingOrEmptyAnnotations(src.ObjectMeta),
	)

	if found {
		dst.Spec.Bar = barVal
	}

	dst.ObjectMeta = dstNewObjectMeta

	return nil
}
