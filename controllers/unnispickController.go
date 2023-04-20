package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"unnispick/entities"
	"unnispick/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UnnispickController struct {
	DB               *gorm.DB
	UnnispickService services.UnnispickService
}

type RequestMember struct {
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	SkinType  string `json:"skin_type"`
	SkinColor string `json:"skin_color"`
}

type RequestLike struct {
	ReviewProductID uint `json:"review_product_id"`
	MemberID        uint `json:"member_id"`
}

type json map[string]interface{}

func (uc *UnnispickController) CreateMember(ctx echo.Context) error {
	var requestMember RequestMember
	if err := ctx.Bind(&requestMember); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	member := entities.Member{
		Username:  requestMember.Username,
		Gender:    requestMember.Gender,
		Skintype:  requestMember.SkinType,
		Skincolor: requestMember.SkinColor,
	}

	result, err := uc.UnnispickService.CreateMember(member)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status":       "success",
		"created_data": result,
	})
}

func (uc *UnnispickController) UpdateMember(ctx echo.Context) error {
	id := ctx.Param("id")
	MemberId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%s not an integer", id),
		})
	}

	var requestMember RequestMember
	if err := ctx.Bind(&requestMember); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	member := entities.Member{
		ID:        uint(MemberId),
		Username:  requestMember.Username,
		Gender:    requestMember.Gender,
		Skintype:  requestMember.SkinType,
		Skincolor: requestMember.SkinColor,
	}

	result, err := uc.UnnispickService.UpdateMember(member)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status":       "success",
		"updated_data": result,
	})
}

func (uc *UnnispickController) DeleteMember(ctx echo.Context) error {
	id := ctx.Param("id")
	MemberId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%s not an integer", id),
		})
	}

	err = uc.UnnispickService.DeleteMember(uint(MemberId))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
	})
}

func (uc *UnnispickController) GetAllMember(ctx echo.Context) error {
	result, err := uc.UnnispickService.GetAllMember()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   result,
	})
}

func (uc *UnnispickController) GetProductById(ctx echo.Context) error {
	id := ctx.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%s not an integer", id),
		})
	}

	result, err := uc.UnnispickService.GetProductById(uint(productId))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   result,
	})
}

func (uc *UnnispickController) Like(ctx echo.Context) error {
	var requestLike RequestLike
	if err := ctx.Bind(&requestLike); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	like := entities.LikeReview{
		ReviewProductID: requestLike.ReviewProductID,
		MemberID:        requestLike.MemberID,
	}

	exist := uc.UnnispickService.GetLikeReview(like)

	if exist {
		err := uc.UnnispickService.DeleteLike(like)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, json{
				"status":  "failed",
				"message": err.Error(),
			})
		}

		return ctx.JSON(http.StatusOK, json{
			"status": "success dislike",
		})
	}
	err := uc.UnnispickService.CreateLike(like)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success like",
	})

}
