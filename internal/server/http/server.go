package http

import (
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/service"
	"admin/pkg/conf"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
)

type Server struct {
	service service.Service
	mongo   dao.Mongo
}

type Router struct {
	server Server
}

type count struct {
	Answered   int `json:"answered"`
	Unanswered int `json:"unanswered"`
}

func New(service *service.Service, mongo *dao.Mongo, yaml *conf.Yaml) (*iris.Application, error) {
	app := iris.New()
	app.Use(logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  []string{"User-Agent"},
		LogFunc:            nil,
		Skippers:           nil,
	}))

	myRouter := &Router{server: Server{service: *service, mongo: *mongo}}
	app.Get("/count", myRouter.count)
	app.Get("/getLog", myRouter.getLog)
	app.PartyFunc("/questions", myRouter.registerQuestionsRoutes)
	app.PartyFunc("/question", myRouter.registerQuestionRoutes)
	if err := app.Run(iris.Addr(fmt.Sprintf(":%s", yaml.Http.Port))); err != nil {
		return nil, err
	}
	return app, nil
}

func before(ctx iris.Context) {
	tokenString := ctx.Request().Header.Get("Authorization")
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("mySecret"), nil
	})
	claims, _ = token.Claims.(jwt.MapClaims)
	ctx.Values().Set("name", claims["sub"])
	ctx.Next()
}

func (r *Router) count(ctx iris.Context) {
	_ = success(ctx, count{
		Answered:   r.server.service.QuestionsService.Count(),
		Unanswered: r.server.service.QuestionService.Count(),
	})
}

func (r *Router) getLog(ctx iris.Context) {
	logs, err := r.server.mongo.FindLogs()
	if err != nil {
		_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
	} else {
		_ = success(ctx, logs)
	}
}

func (r *Router) registerQuestionsRoutes(router iris.Party) {
	router.Post("/add", before, r.server.add)
	router.Post("/delete", before, r.server.delete)
	router.Post("/update", before, r.server.update)
	router.Get("/get_all", r.server.getAll)
	router.Post("/importExcel", r.server.importExcel)
}

func (s *Server) add(ctx iris.Context) {
	var questions model.Questions
	if err := ctx.ReadJSON(&questions); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if questions.Question == "" || questions.Answer == "" {
		_ = failed(ctx, iris.StatusBadRequest, "问题或答案不能为空")
	} else {
		if err := s.service.QuestionsService.Add(questions); err == nil {
			defer s.mongo.InsertLog(questions, ctx.Values().GetString("name"), "add")
			_ = success(ctx, questions)
		} else {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		}
	}
}

func (s *Server) delete(ctx iris.Context) {
	var questions model.Questions
	if err := ctx.ReadJSON(&questions); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if questions.Id == 0 {
		_ = failed(ctx, iris.StatusBadRequest, "id不能为空")
	} else {
		if err := s.service.QuestionsService.Delete(questions); err != nil {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		} else {
			defer s.mongo.InsertLog(questions, ctx.Values().GetString("name"), "delete")
			_ = success(ctx, "")
		}
	}
}

func (s *Server) update(ctx iris.Context) {
	var questions model.Questions
	if err := ctx.ReadJSON(&questions); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if questions.Id == 0 {
		_ = failed(ctx, iris.StatusBadRequest, "id不能为空")
	} else {
		s.mongo.InsertLog(questions, ctx.Values().GetString("name"), "update")
		if err := s.service.QuestionsService.Update(questions); err != nil {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		} else {
			_ = success(ctx, "")
		}
	}
}

func (s *Server) getAll(ctx iris.Context) {
	if result, err := s.service.QuestionsService.GetAll(); err == nil {
		_ = success(ctx, result)
	} else {
		_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
	}
}

func (s *Server) importExcel(ctx iris.Context) {
	var questions []model.Questions
	if err := ctx.ReadJSON(&questions); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	} else {
		if err := s.service.QuestionsService.AddBatch(questions); err != nil {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		} else {
			_ = success(ctx, "")
		}
	}
}

func (r *Router) registerQuestionRoutes(router iris.Party) {
	router.Post("/add", r.server.addQuestion)
	router.Post("/delete", r.server.deleteQuestion)
	router.Post("/update", r.server.updateQuestion)
	router.Get("/get_all", r.server.getAllQuestion)
}

func (s *Server) addQuestion(ctx iris.Context) {
	var question model.Question
	if err := ctx.ReadJSON(&question); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if question.Question == "" {
		_ = failed(ctx, iris.StatusBadRequest, "问题或答案不能为空")
	} else {
		if err := s.service.QuestionService.Add(question); err == nil {
			_ = success(ctx, question)
		} else {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		}
	}
}

func (s *Server) deleteQuestion(ctx iris.Context) {
	var question model.Question
	if err := ctx.ReadJSON(&question); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if question.Id == 0 {
		_ = failed(ctx, iris.StatusBadRequest, "id不能为空")
	} else {
		if err := s.service.QuestionService.Delete(question); err != nil {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		} else {
			_ = success(ctx, "")
		}
	}
}

func (s *Server) updateQuestion(ctx iris.Context) {
	var question model.Question
	if err := ctx.ReadJSON(&question); err != nil {
		_ = failed(ctx, iris.StatusBadRequest, "解析错误：%v", err)
	}
	if question.Id == 0 {
		_ = failed(ctx, iris.StatusBadRequest, "id不能为空")
	} else {
		if err := s.service.QuestionService.Update(question); err != nil {
			_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
		} else {
			_ = success(ctx, "")
		}
	}
}

func (s *Server) getAllQuestion(ctx iris.Context) {
	if result, err := s.service.QuestionService.GetAll(); err == nil {
		_ = success(ctx, result)
	} else {
		_ = failed(ctx, iris.StatusInternalServerError, "服务器错误：%v", err)
	}
}

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func success(ctx context.Context, data interface{}) (err error) {
	res := result{
		Code:    iris.StatusOK,
		Message: "操作成功",
		Data:    data,
	}
	_, err = ctx.JSON(res)
	return err
}

func failed(ctx context.Context, statusCode int, format string, a ...interface{}) (err error) {
	res := result{
		Code:    statusCode,
		Message: fmt.Sprintf(format, a...),
		Data:    nil,
	}
	if statusCode >= 500 {
		ctx.Application().Logger().Error(err)
	}
	_, err = ctx.JSON(res)
	return err
}

func validateFailed(ctx context.Context, message string) (err error) {
	res := result{
		Code:    iris.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
	_, err = ctx.JSON(res)
	return err
}

func forbidden(ctx context.Context) (err error) {
	res := result{
		Code:    iris.StatusForbidden,
		Message: "没有相关权限",
		Data:    nil,
	}
	_, err = ctx.JSON(res)
	return err
}

func unauthorized(ctx context.Context) (err error) {
	res := result{
		Code:    iris.StatusUnauthorized,
		Message: "",
		Data:    nil,
	}
	_, err = ctx.JSON(res)
	return err
}
