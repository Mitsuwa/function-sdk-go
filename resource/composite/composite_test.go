/*
Copyright 2023 The Crossplane Authors.

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

package composite

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func Example() {

	manifest := []byte(`
apiVersion: example.org/v1
kind: CoolCompositeResource
metadata:
  name: my-cool-xr
spec:
  coolness: 9001
`)

	// Create a new, empty XR.
	xr := New()

	// Unmarshal our manifest into the XR. This is just for illustration
	// purposes - the SDK functions like GetObservedCompositeResource do this
	// for you.
	_ = yaml.Unmarshal(manifest, xr)

	// Get the spec.coolness field of the XR.
	coolness, _ := xr.GetInteger("spec.coolness")
	fmt.Println(coolness)
	// Output: 9001
}
