package release

import (
	releaseapi "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	"github.com/caicloud/release-controller/pkg/kube"
	"github.com/caicloud/release-controller/pkg/render"
	"github.com/caicloud/release-controller/pkg/storage"
	"github.com/golang/glog"
)

func (rc *releaseContext) createRelease(backend storage.ReleaseStorage, release *releaseapi.Release) error {
	glog.V(4).Infof("Creating release: %s/%s", release.Namespace, release.Name)
	// Prepare
	_, err := backend.AddCondition(storage.ConditionCreating())
	if err != nil {
		glog.Errorf("Failed to add condition for release %s/%s: %v", release.Namespace, release.Name, err)
		return recordError(backend, err)
	}

	// Render
	carrier, err := rc.render.Render(&render.RenderOptions{
		Namespace: release.Namespace,
		Release:   release.Name,
		Version:   1,
		Template:  release.Spec.Template,
		Config:    release.Spec.Config,
	})
	if err != nil {
		glog.Errorf("Failed to render release %s/%s: %v", release.Namespace, release.Name, err)
		return recordError(backend, err)
	}
	// Create Resources
	resources := carrier.Resources()
	err = rc.client.Create(release.Namespace, resources, kube.CreateOptions{
		OwnerReferences: referencesForRelease(release),
	})
	if err != nil {
		glog.Errorf("Failed to create resources for release %s/%s: %v", release.Namespace, release.Name, err)
		return recordError(backend, err)
	}
	// Record success status
	release.Status.Manifest = render.MergeResources(resources)
	_, err = backend.Update(release)
	if err != nil {
		glog.Errorf("Failed to create release %s/%s: %v", release.Namespace, release.Name, err)
		return recordError(backend, err)
	}
	glog.V(4).Infof("Created release: %s/%s", release.Namespace, release.Name)
	return nil
}
