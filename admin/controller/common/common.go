package common

import (
	"b5gocmf/admin/lib"
	"b5gocmf/utils/upload"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommonsController struct {
	lib.Controller
}

func (c *CommonsController) Route(engine *gin.Engine, group *gin.RouterGroup) {
	group.GET(c.Dispatch("cropper", false, c.Cropper))
	group.POST(c.Dispatch("upload_img", false, c.UploadImg))
	group.POST(c.Dispatch("upload_file", false, c.UploadFile))
	group.POST(c.Dispatch("upload_video", false, c.UploadVideo))
}

func NewCommonController() *CommonsController {
	c := &CommonsController{}
	c.Id = "common"
	return c
}

func (c *CommonsController) Cropper(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	cat := ctx.DefaultQuery("cat", "")
	c.Render(ctx, "cropper", gin.H{"id": id, "cat": cat})
}

// UploadImg 上传图片
func (c *CommonsController) UploadImg(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")
	widthParams := ctx.DefaultPostForm("width", "0")
	heightParams := ctx.DefaultPostForm("height", "0")
	width, _ := strconv.Atoi(widthParams)
	height, _ := strconv.Atoi(heightParams)

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("img", upload.UploaderWithCat(catParams), upload.UploaderWithWidthHeight(width, height))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}

// UploadFile 上传文件
func (c *CommonsController) UploadFile(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("file", upload.UploaderWithCat(catParams))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}

// UploadVideo 上传视频
func (c *CommonsController) UploadVideo(ctx *gin.Context) {
	//获取参数
	catParams := ctx.DefaultPostForm("cat", "")

	//实例化上传方法
	action, errAction := upload.NewUploaderAction("video", upload.UploaderWithCat(catParams))
	if errAction != nil {
		c.Error(ctx, "上传失败："+errAction.Error())
		return
	}
	//上传
	result, errUpload := action.Upload(ctx)
	if errUpload != nil {
		c.Error(ctx, errUpload.Error())
		return
	}
	c.Success(ctx, "上传成功", result)
}
