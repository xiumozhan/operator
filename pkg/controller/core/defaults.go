// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/intstr"

	operatorv1alpha1 "github.com/tigera/operator/pkg/apis/operator/v1alpha1"
)

// fillDefaults fills in the default values for an instance.
func fillDefaults(instance *operatorv1alpha1.Core) {
	if len(instance.Spec.Version) == 0 {
		instance.Spec.Version = "latest"
	}
	if len(instance.Spec.Datastore.Type) == 0 {
		instance.Spec.Datastore.Type = operatorv1alpha1.Kubernetes
	}
	if len(instance.Spec.Registry) == 0 {
		instance.Spec.Registry = "docker.io/"
	}
	if !strings.HasSuffix(instance.Spec.Registry, "/") {
		instance.Spec.Registry = fmt.Sprintf("%s/", instance.Spec.Registry)
	}
	if len(instance.Spec.Variant) == 0 {
		instance.Spec.Variant = operatorv1alpha1.Calico
	}
	if len(instance.Spec.CNINetDir) == 0 {
		instance.Spec.CNINetDir = "/etc/cni/net.d"
	}
	if len(instance.Spec.CNIBinDir) == 0 {
		instance.Spec.CNIBinDir = "/opt/cni/bin"
	}
	if len(instance.Spec.IPPools) == 0 {
		instance.Spec.IPPools = []operatorv1alpha1.IPPool{
			{CIDR: "192.168.0.0/16"},
		}
	}
	if instance.Spec.Components.KubeProxy.Required {
		if len(instance.Spec.Components.KubeProxy.Image) == 0 {
			// Openshift's latest release uses Kubernetes 1.13. This is the latest stable kube-proxy version under that
			// release as of 5.21.19.
			instance.Spec.Components.KubeProxy.Image = "k8s.gcr.io/kube-proxy:v1.13.6"
		}
	}
	if instance.Spec.Components.Node.MaxUnavailable == nil {
		mu := intstr.FromInt(1)
		instance.Spec.Components.Node.MaxUnavailable = &mu
	}
}