package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// getFullfillmentStatusHandler simulates actual fulfillment by supplying bogus values for QuantityInStock
// and ShipsWithin for any given SKU. Used to demonstrate a backing service supporting
// a primary service.
func getFullfillmentStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		sku := vars["sku"]
		formatter.JSON(w, http.StatusOK, fulfillmentStatus{
			SKU:             sku,
			ShipsWithin:     14,
			QuantityInStock: 100,
		})
	}
}

func rootHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.Text(w, http.StatusOK, "Fulfillment Service, see http://github.com/cloudnativego/backing-fulfillment for API.")
	}
}
