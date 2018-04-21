# clarch

## Gettign Started

### Installation

```
go get github.com/locona/clarch
```

### New Project
```
clarch new sample-project
```

if already repository exist.

```
clarch init
```

### Create New Service

```
clarch add book

// stdout
--------------------------------------------------

// import books
bookHandler "github.com/locona/clarch/book/handler/http"
bookUC "github.com/locona/clarch/book/usecase"
bookRepo "github.com/locona/clarch/book/repository"

// book
{
	bookRepo := bookRepo.NewBookRepository(infra.Mysql)
	bookUC := bookUC.NewBookUsecase(bookRepo)
	bookHandler.NewBookHttpHandler(api, bookUC)
}
--------------------------------------------------
```

And you should register generated services.

```diff
// cmd/api.go

import (
    "github.com/locona/clarch-sample/infra"
    "github.com/locona/clarch-sample/project"
    "github.com/locona/clarch-sample/project/logger"
    "github.com/locona/clarch-sample/project/middleware"
    "github.com/locona/clarch-sample/project/validator"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cobra"

+   // import books
+   bookHandler "github.com/locona/clarch-sample/book/handler/http"
+   bookUC "github.com/locona/clarch-sample/book/usecase"
+   bookRepo "github.com/locona/clarch-sample/book/repository"
)

func ListenApi() {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// middleware
	{
		middle := middleware.NewMiddleware()
		api.Use(middle.Cors())
		api.Use(middle.Logging())
	}

+	// book
+	{
+		bookRepo := bookRepo.NewBookRepository(infra.Mysql)
+		bookUC := bookUC.NewBookUsecase(bookRepo)
+		bookHandler.NewBookHttpHandler(api, bookUC)
+	}

	infra.Mysql.LogMode(true)
	api.Run(":" + project.Config.API.Port)
}


```
