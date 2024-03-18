module github.com/ooaklee/ghatd-detail-web-demo-landing-dash-and-more

go 1.22.0

require (
	github.com/ooaklee/ghatd v0.1.1-0.20240316161116-dc3d856805a7
	//>ghatd {{ define "WebDetailGoModRequirePackages" }}
	go.uber.org/zap v1.27.0
//>ghatd {{ end }}
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/matoous/go-nanoid/v2 v2.0.0 // indirect
	github.com/ooaklee/reply v1.0.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20200423170343-7949de9c1215 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

require (
	github.com/gorilla/mux v1.8.1
	go.uber.org/multierr v1.10.0 // indirect
)

// 💻 Enable for local development 💻
// replace github.com/ooaklee/ghatd => ../ghatd
