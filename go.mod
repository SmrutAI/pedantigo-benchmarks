module pedantigo-benchmarks

go 1.24.0

require (
	github.com/SmrutAI/pedantigo v0.0.0
	github.com/danielgtaylor/huma/v2 v2.34.1
	github.com/deepankarm/godantic v0.0.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/go-playground/validator/v10 v10.27.0
	github.com/pasqal-io/godasse v0.12.1
)

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/invopop/jsonschema v0.13.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/SmrutAI/pedantigo => ../Pedantigo

replace github.com/deepankarm/godantic => ../Pedantigo/etc/godantic

replace github.com/pasqal-io/godasse => ../Pedantigo/etc/godasse
