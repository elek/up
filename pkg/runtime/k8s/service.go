// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package k8s

import (
	_ "embed"
	"github.com/zeebo/errs/v2"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
	"os"
	"path/filepath"
	"storj.io/storj-up/pkg/recipe"
	"storj.io/storj-up/pkg/runtime/runtime"
)

//go:embed "deployment.yaml"
var deployment []byte

//go:embed "config.yaml"
var config []byte

type service struct {
	id         runtime.ServiceInstance
	deployment *appsv1.Deployment
	render     func(string) (string, error)
	config     map[string]string
	env        map[string]string
}

func NewService(id runtime.ServiceInstance, r recipe.Service, render func(string) (string, error)) (*service, error) {

	o, err := k8sruntime.Decode(serializer.NewCodecFactory(scheme.Scheme).UniversalDeserializer(), deployment)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	deployment := o.(*appsv1.Deployment)
	name := id.Name
	deployment.Name = name
	deployment.ObjectMeta.Name = name
	deployment.ObjectMeta.Labels["app"] = name
	deployment.Spec.Selector.MatchLabels["app"] = name
	deployment.Spec.Template.Spec.Containers[0].Name = name
	return &service{
		id:         id,
		deployment: deployment,
		render:     render,
		config:     make(map[string]string),
	}, nil
}

func (s *service) UseFile(path string, name string, data string) error {
	// TODO: files are not yet extracted, but we accept recipes with files
	return nil
}

func (s *service) Labels() []string {
	return []string{}
}

func (s *service) RemoveFlag(flag string) error {
	return nil
}

func (s *service) Persist(dir string) error {
	return nil
}

func (s *service) ChangeImage(change func(string) string) error {
	return nil
}

func (s *service) AddPortForward(runtime.PortMap) error {
	return nil
}

func (s *service) ID() runtime.ServiceInstance {
	return s.id
}

func (s *service) AddConfig(key string, value string) error {
	return s.AddEnvironment(key, value)
	//rendered, err := s.render(value)
	//s.config[key] = rendered
	//return err
}

func (s *service) AddFlag(flag string) error {
	return nil
}

func (s *service) AddEnvironment(key string, value string) error {
	rendered, err := s.render(value)
	s.deployment.Spec.Template.Spec.Containers[0].Env = append(s.deployment.Spec.Template.Spec.Containers[0].Env, corev1.EnvVar{
		Name:  key,
		Value: rendered,
	})
	return err
}

func (s *service) write(dir string) error {
	yamlSerializer := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)

	w, err := os.Create(filepath.Join(dir, s.id.Name+".yaml"))
	if err != nil {
		return errs.Wrap(err)
	}
	defer w.Close()
	err = yamlSerializer.Encode(s.deployment, w)
	if err != nil {
		return errs.Wrap(err)
	}

	cfg, err := k8sruntime.Decode(serializer.NewCodecFactory(scheme.Scheme).UniversalDeserializer(), config)
	if err != nil {
		return errs.Wrap(err)
	}
	c := cfg.(*corev1.ConfigMap)
	c.Name = s.ID().Name
	raw, err := yaml.Marshal(s.config)
	if err != nil {
		return errs.Wrap(err)
	}
	c.Data["config.yaml"] = string(raw)

	cw, err := os.Create(filepath.Join(dir, s.id.Name+"-config.yaml"))
	if err != nil {
		return errs.Wrap(err)
	}
	defer cw.Close()
	err = yamlSerializer.Encode(cfg, cw)
	if err != nil {
		return errs.Wrap(err)
	}
	return nil
}
