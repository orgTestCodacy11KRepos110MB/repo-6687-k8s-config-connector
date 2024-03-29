/*
Copyright 2018 The Kubernetes Authors.

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

/*
Package cert provides functions to manage certificates for webhookClientConfiguration.

Create a Provisioner with a CertWriter.

	provisioner := Provisioner{
		CertWriter: admission.NewSecretCertWriter(admission.SecretCertWriterOptions{...}),
	}

Provision the certificates for the webhookClientConfig

	err := provisioner.Provision(Options{
		ClientConfig: webhookClientConfig,
		Objects: []runtime.Object{mutatingWebhookConfiguration, validatingWebhookConfiguration}
	})
	if err != nil {
		// handle error
	}

This package and all its subpackages have been forked from sigs.k8s.io/controller-runtime/pkg/webhook/internal/cert
at version 0.1.12
*/
package cert
