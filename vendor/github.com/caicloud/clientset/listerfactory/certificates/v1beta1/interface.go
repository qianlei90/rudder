/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/certificates/v1beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // CertificateSigningRequests returns a CertificateSigningRequestLister
	CertificateSigningRequests() v1beta1.CertificateSigningRequestLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// CertificateSigningRequests returns a CertificateSigningRequestLister.
func (v *version) CertificateSigningRequests() v1beta1.CertificateSigningRequestLister {
	return &certificateSigningRequestLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// CertificateSigningRequests returns a CertificateSigningRequestLister.
func (v *infromerVersion) CertificateSigningRequests() v1beta1.CertificateSigningRequestLister {
	return v.factory.Certificates().V1beta1().CertificateSigningRequests().Lister()
}