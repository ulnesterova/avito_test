package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	avitotest "github.com/ulnesterova/avito_test"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input avitotest.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segment.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	slug := c.Param("slug")

	err := h.services.Segment.Delete(slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type slugForAdd struct {
	Slug string `json:"slug" binding:"required"`
}

func (h *Handler) addUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	fmt.Print("HELP")

	var slug slugForAdd

	if err := c.BindJSON(&slug); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Print("HELP")
	err = h.services.Segment.AddSegmentToUser(userId, slug.Slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	slug := c.Param("slug")
	fmt.Print("SLUG=" + slug + "||||||")

	err = h.services.Segment.DeleteSegmentFromUser(userId, slug)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getAllListsResponse struct {
	Data []string `json:"segments"`
}

func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	usersSegments, err := h.services.Segment.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: usersSegments,
	})
}
