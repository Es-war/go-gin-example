package routers

import (
	"net/http"

	_ "github.com/Es-war/go-gin-example/docs"
	"github.com/Es-war/go-gin-example/middleware/jwt"
	"github.com/Es-war/go-gin-example/pkg/setting"
	"github.com/Es-war/go-gin-example/pkg/upload"
	"github.com/Es-war/go-gin-example/routers/api"
	"github.com/Es-war/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 当访问 $HOST/upload/images 时，
	// 将会读取到 $GOPATH/src/github.com/Es-war/go-gin-example/runtime/upload/images 下的文件
	// http.Dir() 会将这个GetImageFullPath相对路径拼接到工作目录上，从而形成一个绝对路径
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)		//获取 token
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)


		//获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)
        //新建文章
        apiv1.POST("/articles", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
