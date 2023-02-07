module gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway

go 1.19

replace gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token v0.0.0 => ../../libs/token

replace gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models v0.0.0 => ../../libs/models

require (
	github.com/99designs/gqlgen v0.17.22
	github.com/go-chi/chi v1.5.4
	github.com/imroc/req/v3 v3.30.0
	github.com/vektah/gqlparser/v2 v2.5.1
	gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models v0.0.0
	gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token v0.0.0
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/marten-seemann/qpack v0.2.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.3 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.8.1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220906165146-f3363e06e74c // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
