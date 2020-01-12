package handler

import (
	"errors"
	"net/http"

	"l-semi-chat/pkg/domain"
	"l-semi-chat/pkg/domain/logger"
	"l-semi-chat/pkg/interface/server/middleware"
	"l-semi-chat/pkg/interface/server/response"
	"l-semi-chat/pkg/service/interactor"
	"l-semi-chat/pkg/service/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
	AuthHandler    AuthHandler
	TagHandler     TagHandler
}

// AppHandler ApplicationHandler
type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc
	// ManageAccountTags() http.HandlerFunc
	// ManageAccountTag() http.HandlerFunc

	// auth
	Login() http.HandlerFunc
	Logout() http.HandlerFunc

	// tag
	ManageTags() http.HandlerFunc
	ManageTag() http.HandlerFunc
}

// NewAppHandler create application handler
func NewAppHandler(sh repository.SQLHandler, ph interactor.PasswordHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(sh, ph),
		AuthHandler:    NewAuthHandler(sh, ph),
		TagHandler:     NewTagHandler(sh),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.AccountHandler.GetAccount).ServeHTTP(w, r)
		case http.MethodPost:
			ah.AccountHandler.CreateAccount(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.AccountHandler.UpdateAccount).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.AccountHandler.DeleteAccount).ServeHTTP(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ah.AuthHandler.Login(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			ah.AuthHandler.Logout(w, r)
		default:
			logger.Warn("request method not allowed")
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageTags() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.TagHandler.GetTags(w, r)
		case http.MethodPost:
			middleware.Authorized(ah.TagHandler.CreateTag).ServeHTTP(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.TagHandler.GetTagByTagID(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

// TODO: けして
type getCategoryResponse struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}
