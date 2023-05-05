package operator

import (
	log "github.com/go-logr/logr"
	nacosgroupv1alpha1 "github.com/kevinstudy2021/nacosOperator/api/v1alpha1"
	"github.com/kevinstudy2021/nacosOperator/pkg/service/k8s"
)

type IHealClient interface {
}

type HealClient struct {
}

func NewHealClient(logger log.Logger, k8sService k8s.Services) *HealClient {
	return &HealClient{}
}

func (c HealClient) MakeHeal(nacos *nacosgroupv1alpha1.Nacos) {
	//TODO
}
