package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"app/internal/common/errors"
	"app/internal/common/metrics"
	"app/internal/common/response"
	"app/internal/post/domain"
	"app/internal/post/service"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type PostHandler struct {
	postService service.PostService
	logger      *zap.Logger
}

func NewPostHandler(postService service.PostService, logger *zap.Logger) *PostHandler {
	return &PostHandler{postService: postService, logger: logger}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	}()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.handleError(w, r, errors.NewAppError("INVALID_ID", "Invalid post ID", err))
		return
	}

	post, err := h.postService.GetPost(r.Context(), id)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondWithJSON(w, r, http.StatusOK, "Post retrieved successfully", post)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	}()

	posts, err := h.postService.GetAllPosts(r.Context())
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondWithJSON(w, r, http.StatusOK, "Posts retrieved successfully", posts)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	}()

	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		h.handleError(w, r, errors.NewAppError("INVALID_INPUT", "Invalid input data", err))
		return
	}

	err = h.postService.CreatePost(r.Context(), &post)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondWithJSON(w, r, http.StatusCreated, "Post created successfully", post)
}

func (h *PostHandler) DeletePosts(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	}()

	idStr := r.URL.Query().Get("ids")
	if idStr == "" {
		h.handleError(w, r, errors.NewAppError("MISSING_IDS", "No IDs provided", nil))
		return
	}

	idStrings := strings.Split(idStr, ",")
	ids := make([]int, 0, len(idStrings))

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.handleError(w, r, errors.NewAppError("INVALID_ID", "Invalid post ID", err))
			return
		}
		ids = append(ids, id)
	}

	err := h.postService.DeletePosts(r.Context(), ids)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.respondWithJSON(w, r, http.StatusNoContent, "Posts deleted successfully", nil)
}

func (h *PostHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	var statusCode int
	var message string

	if appErr, ok := err.(*errors.AppError); ok {
		switch appErr.Code {
		case "INVALID_ID", "INVALID_INPUT", "MISSING_IDS":
			statusCode = http.StatusBadRequest
		case "NOT_FOUND":
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
		message = appErr.Message
	} else {
		statusCode = http.StatusInternalServerError
		message = "An unexpected error occurred"
	}

	h.logger.Error("Request error",
		zap.Error(err),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.Int("status", statusCode),
	)

	metrics.RequestsTotal.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(statusCode)).Inc()
	response.JSON(w, statusCode, response.StandardResponse{Status: "error", Message: message})
}

func (h *PostHandler) respondWithJSON(w http.ResponseWriter, r *http.Request, statusCode int, message string, data interface{}) {
	metrics.RequestsTotal.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(statusCode)).Inc()
	response.JSON(w, statusCode, response.StandardResponse{Status: "success", Message: message, Data: data})
}
