[![wercker status](https://app.wercker.com/status/bd3db7ca7efa8babc97cb84a73622bc6/m "wercker status")](https://app.wercker.com/project/bykey/bd3db7ca7efa8babc97cb84a73622bc6)

# Fulfillment Service
This is the fulfillment service that is used as part of the backing services chapter.

This application's dependencies are glide managed. Once you clone this repo:

* `glide install` - Fetches dependencies and populates the `vendor/` directory.
* `go build` - Creates the application binary, honoring the locally-vendored dependencies from the **vendor** directory.
* Run the application. By default, it will start on port *3001*, the port on which the default local copy of the [Catalog Service](https://github.com/cloudnativego/backing-catalog) runs.

# Service API
This is a fake service, so it has no real data. It just returns manufactured payloads based on what you supply.

| Resource | Method | Description |
|---|---|---|
| skus/{sku} | GET | Fabricates a stock quantity and shipping time for an arbitrary SKU. This SKU can be any string |
