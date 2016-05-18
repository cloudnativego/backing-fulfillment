# Fulfillment Service
This is the fulfillment service that is used as part of the backing services chapter.

# Service API
This is a fake service, so it has no real data. It just returns manufactured payloads based on what you supply.

| Resource | Method | Description |
|---|---|---|
| skus/{sku} | GET | Fabricates a stock quantity and shipping time for an arbitrary SKU. This SKU can be any string |
