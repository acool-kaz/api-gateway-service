package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/acool-kaz/api-gateway-service/internal/models"
	post_crud_pb "github.com/acool-kaz/api-gateway-service/pkg/post_crud/pb"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := h.postClient.Client.Read(r.Context(), &post_crud_pb.ReadRequest{})
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		h.errorHandler(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.postClient.Client.Create(r.Context(), &post_crud_pb.CreateRequest{Post: models.FromPostToProto(post)})
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getPostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "post_id"), 10, 32)
	if err != nil {
		h.errorHandler(w, http.StatusNotFound, err.Error())
		return
	}

	postId := int32(id)

	resp, err := h.postClient.Client.Read(r.Context(), &post_crud_pb.ReadRequest{Id: &postId})
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) deletePostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "post_id"), 10, 32)
	if err != nil {
		h.errorHandler(w, http.StatusNotFound, err.Error())
		return
	}

	postId := int32(id)

	resp, err := h.postClient.Client.Delete(r.Context(), &post_crud_pb.DeleteRequest{Id: postId})
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) updatePostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "post_id"), 10, 32)
	if err != nil {
		h.errorHandler(w, http.StatusNotFound, err.Error())
		return
	}

	postId := int32(id)

	var update models.Post

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.errorHandler(w, http.StatusBadRequest, err.Error())
		return
	}

	newId := int32(update.Id)
	newUserId := int32(update.UserId)

	req := post_crud_pb.UpdateRequest{
		Id:        postId,
		NewId:     &newId,
		NewUserId: &newUserId,
		NewTitle:  &update.Title,
		NewBody:   &update.Body,
	}

	resp, err := h.postClient.Client.Update(r.Context(), &req)
	if err != nil {
		h.errorHandler(w, http.StatusBadGateway, err.Error())
		return
	}

	if err = json.NewEncoder(w).Encode(&resp); err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}
}
