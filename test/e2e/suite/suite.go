// Copyright Jetstack Ltd. See LICENSE for details.
package suite

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"

	"github.com/jetstack/kube-oidc-proxy/test/e2e/framework"
	"github.com/jetstack/kube-oidc-proxy/test/environment"
)

var (
	env *environment.Environment
	cfg = framework.DefaultConfig
)

var _ = SynchronizedBeforeSuite(func() []byte {
	var err error
	env, err = environment.Create(1, 3)
	if err != nil {
		log.Fatalf("Error provisioning environment: %v", err)
	}

	cfg.KubeConfigPath = env.KubeConfigPath()
	cfg.Kubectl = filepath.Join(env.RootPath(), "bin", "kubectl")
	cfg.RepoRoot = env.RootPath()
	cfg.Environment = env

	if err := framework.DefaultConfig.Validate(); err != nil {
		log.Fatalf("Invalid test config: %v", err)
	}

	return nil
}, func([]byte) {
})

var _ = SynchronizedAfterSuite(func() {},
	func() {
		if env != nil {
			if err := env.Destory(); err != nil {
				log.Fatalf("Failed to destory environment: %s", err)
			}
		}
	},
)
