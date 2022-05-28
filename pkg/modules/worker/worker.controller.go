package worker

import (
	"github.com/go-chi/chi"
	"github.com/thehaung/cloudflare-worker/pkg/utils"
	"io"
	"net/http"
)

type Controller struct {
	r       *chi.Mux
	service IService
}

func NewWorkerController(r *chi.Mux, service IService) {
	controller := Controller{
		r:       r,
		service: service,
	}

	controller.Init()
}

func (c Controller) Init() {
	c.r.Route("/workers", func(r chi.Router) {
		r.Get("/dns", c.FetchDNS)
		r.Get("/ip", c.GetIPAddress)
		r.Get("/cloudflare", c.GetListDNSCloudFlare)
	})
}

func (c Controller) FetchDNS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	httpUtil := utils.NewHttpUtil(w, r)

	if result, err := c.service.FetchDNS(ctx); err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, err.Error())
		return
	} else {
		httpUtil.WriteJson(http.StatusOK, result)
		return
	}
}

func (c Controller) GetIPAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	httpUtil := utils.NewHttpUtil(w, r)

	if result, err := c.service.GetIPAddress(ctx); err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, err.Error())
		return
	} else {
		httpUtil.WriteJson(http.StatusOK, result)
		return
	}
}

func (c Controller) GetListDNSCloudFlare(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	httpUtil := utils.NewHttpUtil(w, r)

	w.Header().Set("Content-Type", "application/json")

	if result, err := c.service.GetListDNSCloudFlare(ctx); err != nil {
		httpUtil.WriteError(http.StatusInternalServerError, err.Error())
		return
	} else {
		if _, err := io.Copy(w, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
