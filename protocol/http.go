package protocol

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/http/label"
	"github.com/motongxue/keyauth-g7/apps/endpoint"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	keyauth_rpc "github.com/motongxue/keyauth-g7/client/rpc"
	keyauth_auth "github.com/motongxue/keyauth-g7/client/rpc/auth"

	"github.com/motongxue/cmdb-g7/conf"
	"github.com/motongxue/cmdb-g7/swagger"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {

	r := restful.DefaultContainer
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	// http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"*"},
		CookiesAllowed: false,
		Container:      r}
	r.Filter(cors.Filter)

	keyauthClient, err := keyauth_rpc.NewClient(conf.C().Mcenter)
	if err != nil {
		panic(err)
	}
	auther := keyauth_auth.NewKeyauthAuther(keyauthClient, "cmdb")
	fmt.Println(auther)
	// 通过使用keyauth提供的中间件实现认证服务
	r.Filter(auther.RestfulAuthHandlerFunc)
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTP.Addr(),
		Handler:           r,
	}

	return &HTTPService{
		kc:     keyauthClient,
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	kc     *keyauth_rpc.ClientSet
	r      *restful.Container
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	app.LoadRESTfulApp(s.PathPrefix(), s.r)

	// API Doc
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: swagger.Docs}
	s.r.Add(restfulspec.NewOpenAPIService(config))
	s.l.Infof("Get the API using http://%s%s", s.c.App.HTTP.Addr(), config.APIPath)

	// 此时所有的webservice已经加载完成
	if err := s.Registry(); err != nil {
		// 注册流程不影响启动流程，不retrun
		s.l.Errorf("registry failed, %s", err)
	}

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}

// Registry 通过Keyauth SDK 注册服务功能
// 什么时候注册? 服务启动的时候? 需要WebService都已经加载完成,才能使用RegisteredWebServices()
// 一定要等到所有WebService已经加载到router后
// TODO 这里实际上每一次启动服务都将所有服务重新注册了一遍，导致数据库中冗余
func (s *HTTPService) Registry() error {
	registeredWebServices := s.r.RegisteredWebServices()
	endpoints := endpoint.EndpointSet{
		Service:   "cmdb",
		Endpoints: []*endpoint.Endpoint{},
	}
	for i := range registeredWebServices {
		// 取出每一个web service路由
		routes := registeredWebServices[i].Routes()
		for _, r := range routes {
			var resource, action string
			if r.Metadata != nil {
				if v, ok := r.Metadata[label.Resource]; ok {
					resource, _ = v.(string)
				}
				if v, ok := r.Metadata[label.Action]; ok {
					action, _ = v.(string)
				}
			}
			endpoints.Endpoints = append(endpoints.Endpoints, &endpoint.Endpoint{
				Resource: resource,
				Action:   action,
				Path:     r.Path,
				Method:   r.Method,
			})
		}
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	registryEndpoint, err := s.kc.Endpoint().RegistryEndpoint(ctx, &endpoints)
	if err != nil {
		return err
	}
	s.l.Debugf("registry response: %s", registryEndpoint)
	return nil
}
