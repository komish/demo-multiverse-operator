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

package v1alpha1

import (
	"strconv"

	"github.com/komish/demo-multiverse-operator/api/conversionutil"
	demov1beta1 "github.com/komish/demo-multiverse-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo partially implements the conversion.Convertible interface
// indicates that this version is a spoke.
func (src *MyCustomKind) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*demov1beta1.MyCustomKind) // assert the type because we have the raw
	srcSpec := src.Spec

	// Handle Foo
	dst.Spec.NewFoo = src.Spec.Foo

	// Handle Zap
	if srcSpec.Zap > 0 {
		dstAnnotations := dst.GetAnnotations()

		dstAnnotations[conversionutil.ZapAnnotation] = strconv.Itoa(int(srcSpec.Zap))
		dst.SetAnnotations(dstAnnotations)
	}

	// Handle Bar
	srcAnnotations := src.GetAnnotations()
	if barVal, exists := srcAnnotations[conversionutil.BarAnnotation]; exists {
		dst.Spec.Bar = barVal
	}

	return nil
}

// ConvertFrom partially implements the conversion.Convertible interface
// indicates that this version is a spoke.
func (dst *MyCustomKind) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*demov1beta1.MyCustomKind) // assert the type because we have the raw

	// Handle Foo
	dst.Spec.Foo = src.Spec.NewFoo

	// Handle Zap
	srcAnnotations := src.GetAnnotations()
	if zapValStr, exists := srcAnnotations[conversionutil.ZapAnnotation]; exists {
		zapValInt64, err := strconv.ParseInt(zapValStr, 10, 32)
		if err != nil {
			return err
		}

		zapValInt32 := int32(zapValInt64)
		dst.Spec.Zap = zapValInt32
	}

	// Handle Bar
	if len(src.Spec.Bar) > 0 {
		dstAnnotations := dst.GetAnnotations()

		dstAnnotations[conversionutil.BarAnnotation] = src.Spec.Bar
		dst.SetAnnotations(dstAnnotations)
	}

	return nil
}
