package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"app/internal/post/domain"
	"app/internal/post/service"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetPost(r.Context(), id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.postService.CreatePost(r.Context(), &post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) DeletePosts(w http.ResponseWriter, r *http.Request) {
	// Get ids from query params
	idStr := r.URL.Query().Get("ids")
	if idStr == "" {
		http.Error(w, "No IDs provided", http.StatusBadRequest)
		return
	}

	idStrings := strings.Split(idStr, ",")
	ids := make([]int, 0, len(idStrings))

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid post ID: %s", idStr), http.StatusBadRequest)
			return
		}
		ids = append(ids, id)
	}

	err := h.postService.DeletePosts(r.Context(), ids)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting posts: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
