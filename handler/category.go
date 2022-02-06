package handler

import (
	"context"
	"github.com/2559065/category/common"
	"github.com/2559065/category/domain/model"
	"github.com/2559065/category/domain/service"
	"github.com/micro/go-micro/debug/log"

	category "github.com/2559065/category/proto/category"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

func (c *Category)CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	// 赋值
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil
}

// 提供分类更新服务
func (c *Category)UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}
// 提供分类删除
func (c *Category)DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

// 根据分类名称查找分类
func (c *Category)FindCategoryByName(ctx context.Context, req *category.FindByNameRequest, res *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, res)
}
// 根据分类id查找
func (c *Category)FindCategoryByID(ctx context.Context, req *category.FindByIdRequest, res *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(req.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, res)
}
func (c *Category)FindCatehoryByLevel(ctx context.Context, req *category.FindByLevelRequest, res *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(req.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) error {
	for _,cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
	return nil
}
func (c *Category)FindCategoryByParent(ctx context.Context, req *category.FindByParentRequest, res *category.FindAllResponse) error{
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(req.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil
}
func (c *Category)FindAllCategory(ctx context.Context, req *category.FindAllRequest, res *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, res)
	return nil
}
