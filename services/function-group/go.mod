module gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group

go 1.19

require (
	github.com/gofiber/fiber/v2 v2.41.0
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.28.0
	github.com/samber/lo v1.37.0
	github.com/uptrace/bun v1.1.10
	github.com/uptrace/bun/dialect/pgdialect v1.1.10
	github.com/uptrace/bun/driver/pgdriver v1.1.10
	gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models v0.0.0
	gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.43.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/sys v0.4.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)

replace gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token v0.0.0 => ../../libs/token

replace gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/models v0.0.0 => ../../libs/models

replace gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway v0.0.0 => ../api-gateway
