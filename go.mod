module github.com/kubeflow/pipelines

require (
	github.com/Masterminds/squirrel v0.0.0-20190107164353-fa735ea14f09
	github.com/VividCortex/mysqlerr v0.0.0-20170204212430-6c6b55f8796f
	github.com/argoproj/argo v0.0.0-20200506223611-54154c61eb4f
	github.com/cenkalti/backoff v2.0.0+incompatible
	github.com/elazarl/goproxy v0.0.0-20181111060418-2ce16c963a8a // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.4
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.8
	github.com/go-openapi/validate v0.19.5
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/google/addlicense v0.0.0-20200906110928-a0294312aa76 // indirect
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.1
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.8
	github.com/jinzhu/gorm v1.9.12
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/peterhellberg/duration v0.0.0-20191119133758-ec6baeebcd10
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	github.com/tektoncd/pipeline v0.17.3
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	google.golang.org/api v0.31.0
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d
	google.golang.org/grpc v1.31.1
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.18.8
	k8s.io/kubernetes v1.14.7
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19 // indirect
	knative.dev/pkg v0.0.0-20200831162708-14fb2347fb77
	sigs.k8s.io/controller-runtime v0.6.1
)

replace (
	k8s.io/api => k8s.io/api v0.17.9
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.9
	k8s.io/apiserver => k8s.io/apiserver v0.17.9
	k8s.io/client-go => k8s.io/client-go v0.17.9
	k8s.io/code-generator => k8s.io/code-generator v0.17.9
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
)

go 1.13
