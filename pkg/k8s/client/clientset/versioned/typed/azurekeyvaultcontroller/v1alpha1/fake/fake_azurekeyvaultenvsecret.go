/*
Copyright Sparebanken Vest

Based on the Kubernetes controller example at
https://github.com/kubernetes/sample-controller

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/SparebankenVest/azure-key-vault-to-kubernetes/pkg/apis/azurekeyvaultcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureKeyVaultEnvSecrets implements AzureKeyVaultEnvSecretInterface
type FakeAzureKeyVaultEnvSecrets struct {
	Fake *FakeAzurekeyvaultcontrollerV1alpha1
	ns   string
}

var azurekeyvaultenvsecretsResource = schema.GroupVersionResource{Group: "azurekeyvaultcontroller.spv.no", Version: "v1alpha1", Resource: "azurekeyvaultenvsecrets"}

var azurekeyvaultenvsecretsKind = schema.GroupVersionKind{Group: "azurekeyvaultcontroller.spv.no", Version: "v1alpha1", Kind: "AzureKeyVaultEnvSecret"}

// Get takes name of the azureKeyVaultEnvSecret, and returns the corresponding azureKeyVaultEnvSecret object, and an error if there is any.
func (c *FakeAzureKeyVaultEnvSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureKeyVaultEnvSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurekeyvaultenvsecretsResource, c.ns, name), &v1alpha1.AzureKeyVaultEnvSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultEnvSecret), err
}

// List takes label and field selectors, and returns the list of AzureKeyVaultEnvSecrets that match those selectors.
func (c *FakeAzureKeyVaultEnvSecrets) List(opts v1.ListOptions) (result *v1alpha1.AzureKeyVaultEnvSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurekeyvaultenvsecretsResource, azurekeyvaultenvsecretsKind, c.ns, opts), &v1alpha1.AzureKeyVaultEnvSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureKeyVaultEnvSecretList{ListMeta: obj.(*v1alpha1.AzureKeyVaultEnvSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureKeyVaultEnvSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureKeyVaultEnvSecrets.
func (c *FakeAzureKeyVaultEnvSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurekeyvaultenvsecretsResource, c.ns, opts))

}

// Create takes the representation of a azureKeyVaultEnvSecret and creates it.  Returns the server's representation of the azureKeyVaultEnvSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultEnvSecrets) Create(azureKeyVaultEnvSecret *v1alpha1.AzureKeyVaultEnvSecret) (result *v1alpha1.AzureKeyVaultEnvSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurekeyvaultenvsecretsResource, c.ns, azureKeyVaultEnvSecret), &v1alpha1.AzureKeyVaultEnvSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultEnvSecret), err
}

// Update takes the representation of a azureKeyVaultEnvSecret and updates it. Returns the server's representation of the azureKeyVaultEnvSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultEnvSecrets) Update(azureKeyVaultEnvSecret *v1alpha1.AzureKeyVaultEnvSecret) (result *v1alpha1.AzureKeyVaultEnvSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurekeyvaultenvsecretsResource, c.ns, azureKeyVaultEnvSecret), &v1alpha1.AzureKeyVaultEnvSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultEnvSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureKeyVaultEnvSecrets) UpdateStatus(azureKeyVaultEnvSecret *v1alpha1.AzureKeyVaultEnvSecret) (*v1alpha1.AzureKeyVaultEnvSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurekeyvaultenvsecretsResource, "status", c.ns, azureKeyVaultEnvSecret), &v1alpha1.AzureKeyVaultEnvSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultEnvSecret), err
}

// Delete takes name of the azureKeyVaultEnvSecret and deletes it. Returns an error if one occurs.
func (c *FakeAzureKeyVaultEnvSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurekeyvaultenvsecretsResource, c.ns, name), &v1alpha1.AzureKeyVaultEnvSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureKeyVaultEnvSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurekeyvaultenvsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureKeyVaultEnvSecretList{})
	return err
}

// Patch applies the patch and returns the patched azureKeyVaultEnvSecret.
func (c *FakeAzureKeyVaultEnvSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureKeyVaultEnvSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurekeyvaultenvsecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureKeyVaultEnvSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultEnvSecret), err
}
