package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/linimbus/ocr_web/models"
)

// FilesController 文件控制器
type FilesController struct {
}

// Preview 获取并显示文件预览
func (v FilesController) Preview(c buffalo.Context) error {
	// 从URL参数中获取文件ID
	fileID := c.Param("file_id")

	// 创建文件模型实例
	file := &models.File{}

	// 从数据库中查询文件
	err := models.DB.Find(file, fileID)
	if err != nil {
		return c.Error(404, err)
	}

	// 将文件数据传递给模板
	c.Set("file", file)

	return c.Render(http.StatusOK, r.HTML("files/preview.plush.html"))
}

// FilesUpload default implementation.
func FilesUpload(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("files/upload.html"))
}

// FilesPreview default implementation.
func FilesPreview(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("files/preview.html"))
}
