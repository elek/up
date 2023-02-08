// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package k8s

import (
	"github.com/zeebo/errs/v2"
	"storj.io/storj-up/pkg/common"
	"storj.io/storj-up/pkg/recipe"
	"storj.io/storj-up/pkg/runtime/runtime"
	"strconv"
)

type Kubernetes struct {
	dir       string
	services  []*service
	variables map[string]map[string]string
}

func NewKubernetes(dir string) (*Kubernetes, error) {
	return &Kubernetes{
		dir:      dir,
		services: make([]*service, 0),
		variables: map[string]map[string]string{
			"cockroach": {
				"main":     "cockroach://root@cockroach:26257/master?sslmode=disable",
				"metainfo": "cockroach://root@cockroach:26257/metainfo?sslmode=disable",
				"dir":      "/tmp/cockroach",
			},
			"storagenode": {
				"identityDir": "/var/lib/storj/.local/share/storj/identity/storagenode/",
				"staticDir":   "/var/lib/storj/web/storagenode",
			},
			"redis": {
				"url": "redis://redis:6379",
			},
			"satellite-api": {
				"mailTemplateDir": "/var/lib/storj/storj/web/satellite/static/emails/",
				"staticDir":       "/var/lib/storj/storj/web/satellite/",
				"identityDir":     "/var/lib/storj/.local/share/storj/identity/satellite-api/",
				"identity":        common.Satellite0Identity,
			},
			"satellite-core": {
				"mailTemplateDir": "/var/lib/storj/storj/web/satellite/static/emails/",
				"identityDir":     "/var/lib/storj/.local/share/storj/identity/satellite-api/",
			},
			"satellite-admin": {
				"staticDir":   "/var/lib/storj/storj/satellite/admin/ui/build",
				"identityDir": "/var/lib/storj/.local/share/storj/identity/satellite-api/",
			},
			"satellite-gc": {
				"identityDir": "/var/lib/storj/.local/share/storj/identity/satellite-api/",
			},
			"satellite-bf": {
				"identityDir": "/var/lib/storj/.local/share/storj/identity/satellite-api/",
			},
			"satellite-rangedloop": {
				"identityDir": "/var/lib/storj/.local/share/storj/identity/satellite-api/",
			},
			"linksharing": {
				"webDir":    "/var/lib/storj/pkg/linksharing/web/",
				"staticDir": "/var/lib/storj/pkg/linksharing/web/static",
			},
		},
	}, nil
}

func (k *Kubernetes) GetHost(service runtime.ServiceInstance, hostType string) string {
	switch hostType {
	case "listen":
		return "0.0.0.0"
	case "internal":
		if service.Name == "storagenode" {
			return service.Name + strconv.Itoa(service.Instance+1)
		}
		return service.Name
	case "external":
		return "localhost"
	}
	return "???"
}

func (k *Kubernetes) GetPort(service runtime.ServiceInstance, portType string) runtime.PortMap {
	if portType == "debug" {
		return runtime.PortMap{Internal: 11111, External: 11111}
	}
	switch service.Name {
	case "satellite-api":
		switch portType {
		case "public":
			return runtime.PortMap{Internal: 7777, External: 7777}
		case "console":
			return runtime.PortMap{Internal: 10000, External: 10000}
		}
	case "storagenode":
		p, _ := runtime.PortConvention(service, portType)
		return runtime.PortMap{Internal: p, External: p}
	case "gateway-mt":
		if portType == "public" {
			return runtime.PortMap{Internal: 9999, External: 9999}
		}
	case "authservice":
		if portType == "public" {
			return runtime.PortMap{Internal: 8888, External: 8888}
		}
	case "linksharing":
		if portType == "public" {
			return runtime.PortMap{Internal: 9090, External: 9090}
		}
	case "satellite-admin":
		if portType == "console" {
			return runtime.PortMap{Internal: 8080, External: 9080}
		}
	}

	return runtime.PortMap{Internal: -1, External: -1}
}

func (k *Kubernetes) Get(s runtime.ServiceInstance, name string) string {
	return k.variables[s.Name][name]
}

func (k *Kubernetes) AddService(r recipe.Service) (runtime.Service, error) {
	id := runtime.NewServiceInstance(r.Name, 0)
	s, err := NewService(id, r, func(s string) (string, error) {
		return runtime.Render(k, id, s)
	})
	if err != nil {
		return s, errs.Wrap(err)
	}
	err = runtime.InitFromRecipe(s, r)
	if err != nil {
		return s, errs.Wrap(err)
	}
	k.services = append(k.services, s)
	return s, nil
}

func (k *Kubernetes) Write() error {
	for _, s := range k.services {
		err := s.write(k.dir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Kubernetes) GetServices() []runtime.Service {
	//TODO implement me
	panic("implement me")
}

func (k *Kubernetes) Reload(stack recipe.Stack) error {
	//TODO implement me
	panic("implement me")
}

var _ runtime.Runtime = &Kubernetes{}
