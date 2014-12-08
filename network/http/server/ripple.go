//Forked from github.com/laurent22/ripple
//Changed by @Weiran YUAN 2013
//1. make request parameters as form of "?p1=v1&p2=v2...." available in this framework
//

package server

import (
	"encoding/base64"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// A context holds information about the current request and response.
// An object of type Context is passed to methods of a controller.
type Context struct {
	// The parameters matched in the current URL. A parameter is defined
	// in a route by prefixing it with a ":". For example, ":id", ":name".
	Params map[string]string
	// The actual HTTP request.
	Request *http.Request
	// The response object.
	Response *Response
}

// Build a new context object.
func NewContext() *Context {
	output := new(Context)
	output.Params = make(map[string]string)
	output.Response = NewResponse()
	return output
}

// A Ripple application. Use NewApplication() to build it.
type Application struct {
	controllers map[string]interface{}
	routes      []Route
	//contentType   string
	baseUrl       string
	parsedBaseUrl *url.URL
}

// Build a new application object.
func NewApplication() *Application {
	output := new(Application)
	output.controllers = make(map[string]interface{})
	//output.contentType = "application/json"
	output.SetBaseUrl("/")
	return output
}

// A route maps a URL pattern to a controller and action.
// See README.md for more information on the format of the pattern.
type Route struct {
	Pattern    string
	Controller string
	Action     string
}

// Holds information about the HTTP response.
type Response struct {
	//
	ContentType string
	// HTTP code (200, 404, etc.). Use the constants of the http package.
	Status int
	// The response body. It will be serialized automatically
	// by the Ripple application before being sent to the client.
	Body interface{}
}

// Build a new response object.
func NewResponse() *Response {
	output := new(Response)
	output.Body = nil
	output.ContentType = "application/json"
	return output
}

// Helper struct used by `prepareServeHttpResponseData()`
type serveHttpResponseData struct {
	ContentType string
	Status      int
	Body        string
}

func defaultHttpStatus(method string) int {
	output := http.StatusOK
	if method == "POST" {
		output = http.StatusCreated
	}
	return output
}

func (this *Application) Start(port string) {
	this.SetBaseUrl("/")
	http.HandleFunc("/", this.ServeHTTP)
	log.Println("Starting server...")
	http.ListenAndServe(port, nil)
}

// Sets the base URL (default to "/"). The base URL will
// be stripped from the beginning of the request URL. For
// instance if the base URL is "/api/" and the client does
// a request on "/api/images/1", the application will dispatch
// "images/1". Specifying the full URL (with domain, etc.) is
// not necessary.
func (this *Application) SetBaseUrl(v string) {
	this.baseUrl = v
	var err error
	this.parsedBaseUrl, err = url.Parse(this.baseUrl)
	if err != nil {
		log.Panicf("Invalid base URL: %s", this.baseUrl)
	}
}

// Returns the base URL.
func (this *Application) BaseUrl() string {
	return this.baseUrl
}

// Helper function to prepare the response writter data for `ServeHTTP()`
func (this *Application) prepareServeHttpResponseData(context *Context) serveHttpResponseData {
	var statusCode int
	var body string
	var err error
	contentType := "application/json"
	//
	if context == nil {
		statusCode = http.StatusNotFound
	} else {
		statusCode = context.Response.Status
	}
	if context != nil {
		body, err = this.serializeResponseBody(context.Response.Body)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}

		log.Printf("response: %#v", context.Response)
		if context.Response.ContentType != "" {
			contentType = context.Response.ContentType
		}
	}

	var output serveHttpResponseData

	output.ContentType = contentType
	output.Status = statusCode
	output.Body = body
	return output
}

// Serves an HTTP request - implementation of net.http.ServeHTTP
func (this *Application) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	////////////////////////////////////////////

	auth := request.Header.Get("Authorization")
	if auth == "" {
		writter.Header().Set("WWW-Authenticate", `Basic realm="User Login"`)
		// writter.Header().Set("Access-Control-Allow-Origin", "*")
		// writter.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		// writter.Header().Set("content-type", "application/json")
		writter.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println(auth)

	auths := strings.SplitN(auth, " ", 2)
	if len(auths) != 2 {
		log.Println("error")
		return
	}

	authMethod := auths[0]
	authB64 := auths[1]

	switch authMethod {
	case "Basic":
		authstr, err := base64.StdEncoding.DecodeString(authB64)
		if err != nil {
			log.Println(err)
			io.WriteString(writter, "Unauthorized!\n")
			return
		}
		log.Println(string(authstr))

		userPwd := strings.SplitN(string(authstr), ":", 2)
		if len(userPwd) != 2 {
			log.Println("error")
			return
		}

		username := userPwd[0]
		password := userPwd[1]

		log.Println("Username:", username)
		log.Println("Password:", password)

		if username != password {
			writter.Header().Set("WWW-Authenticate", `Basic realm="User Login"`)
			writter.WriteHeader(http.StatusUnauthorized)
			return
		}

	default:
		log.Println("error")
		return
	}

	////////////////////////////////////////////
	context := this.Dispatch(request)
	r := this.prepareServeHttpResponseData(context)

	writter.Header().Set("Content-Type", r.ContentType)
	writter.WriteHeader(r.Status)
	if r.ContentType == "image/jpeg" {
		img := context.Response.Body.(*image.RGBA)
		jpeg.Encode(writter, img, &jpeg.Options{95})
	} else {
		writter.Write([]byte(r.Body))
	}
}

func (this *Application) serializeResponseBody(body interface{}) (string, error) {
	if body == nil {
		return "", nil
	}

	var output string
	var err error
	err = nil

	switch body.(type) {

	case string:

		output = body.(string)

	case int, int8, int16, int32, int64:

		output = strconv.Itoa(body.(int))

	case uint, uint8, uint16, uint32, uint64:

		output = strconv.FormatUint(body.(uint64), 10)

	case float32, float64:

		output = strconv.FormatFloat(body.(float64), 'f', -1, 64)

	case bool:

		if body.(bool) {
			output = "true"
		} else {
			output = "false"
		}

	case *image.RGBA:
		output = ""

	default:

		//contentType := this.contentType
		//if contentType != "application/json" { // Currently, only JSON is supported
		//	log.Printf("Unsupported content type: %s! Defaulting to application/json.", this.contentType)
		//	contentType = "application/json"
		//}

		//if contentType == "application/json" {
		var b []byte
		b, err = json.Marshal(body)
		output = string(b)
		//}

	}

	return output, err
}

func (this *Application) checkRoute(route Route) {
	if route.Controller != "" {
		_, exists := this.controllers[route.Controller]
		if !exists {
			log.Panicf("\"%s\" controller does not exist.\n", route.Controller)
		}

	}
}

// Registers a controller. The name should be the same as in the URL path. For example
// if the URL is "users/1", the name should be "users". The controller itself can be
// any struct that implements HTTP method handlers. See README.md and the demo for more
// details on the structure of a controller.
func (this *Application) RegisterController(name string, controller interface{}) {
	this.controllers[name] = controller
}

// Add a route to the application.
func (this *Application) AddRoute(route Route) {
	log.Printf("AddRoute: %s", route.Pattern)
	this.checkRoute(route)
	this.routes = append(this.routes, route)
	//log.Printf("done")
}

func splitPath(path string) []string {
	var output []string
	if len(path) == 0 {
		return output
	}
	if path[0] == '/' {
		path = path[1:]
	}
	pathTokens := strings.Split(path, "/")
	for i := 0; i < len(pathTokens); i++ {
		e := pathTokens[i]
		if len(e) > 0 {
			output = append(output, e)
		}
	}
	return output
}

func makeMethodName(requestMethod string, actionName string) string {
	return strings.Title(strings.ToLower(requestMethod)) + strings.Title(actionName)
}

// Provided for debugging/testing purposes only.
type MatchRequestResult struct {
	Success          bool
	ControllerName   string
	ActionName       string
	ControllerValue  reflect.Value
	ControllerMethod reflect.Value
	MatchedRoute     Route
	Params           map[string]string
}

func (this *Application) matchRequest(request *http.Request) MatchRequestResult {
	log.Printf("matchRequest")
	var output MatchRequestResult
	output.Success = false

	log.Printf("url: %s | %s\n", request.URL, request.URL.RawQuery)
	parametermap, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		log.Printf("error: ", err)
	}
	log.Printf("parsed queries: ", parametermap)

	path := request.URL.Path
	log.Printf("origin path: %s\n", path)
	path = path[len(this.parsedBaseUrl.Path):len(path)]
	log.Printf("path: %s\n", path)
	pathTokens := splitPath(path)
	log.Printf("pathTokens: %s\n", pathTokens)

	for routeIndex := 0; routeIndex < len(this.routes); routeIndex++ {
		route := this.routes[routeIndex]
		patternTokens := splitPath(route.Pattern)

		if len(patternTokens) != len(pathTokens) {
			continue
		}

		var controller interface{}
		var exists bool

		controllerName := ""
		actionName := ""
		notMatched := false
		params := make(map[string]string)
		for i := 0; i < len(patternTokens); i++ {
			patternToken := patternTokens[i]
			pathToken := pathTokens[i]
			if patternToken == ":_controller" {

				//log.Printf("pathToken: %s\n", pathToken)
				controllerName = pathToken

			} else if patternToken == ":_action" {
				actionName = pathToken
			} else if patternToken == pathToken {

			} else if patternToken[0] == ':' {
				params[patternToken[1:]] = pathToken
			} else {
				notMatched = true
				break
			}
		}

		//traite the condition of ?para1=value1&para2=value2...
		for key, value := range parametermap {
			params[key] = value[0]
		}
		//

		if notMatched {
			continue
		}

		if controllerName == "" {
			controllerName = route.Controller
		}

		if actionName == "" {
			actionName = route.Action
		}

		log.Printf("parse request: controller='%s' action='%s'", controllerName, actionName)

		controller, exists = this.controllers[controllerName]
		if !exists {
			continue
			log.Printf("no matched controller")
		}

		log.Printf("matched controller: %s", controller)

		methodName := makeMethodName(request.Method, actionName)
		controllerVal := reflect.ValueOf(controller)
		log.Printf("matched controller method: %s.%s", controllerVal, methodName)

		controllerMethod := controllerVal.MethodByName(methodName)
		if !controllerMethod.IsValid() {
			log.Printf("controller is not valid")
			continue
		}

		output.Success = true
		output.ControllerName = controllerName
		output.ActionName = actionName
		output.ControllerValue = controllerVal
		output.ControllerMethod = controllerMethod
		output.MatchedRoute = route
		output.Params = params
	}

	return output
}

// Provided for debugging/testing purposes only.
func (this *Application) Dispatch(request *http.Request) *Context {
	r := this.matchRequest(request)
	if !r.Success {
		log.Printf("No match for: %s %s\n", request.Method, request.URL)
		return nil
	}

	ctx := NewContext()
	ctx.Request = request
	ctx.Params = r.Params
	ctx.Response.Status = defaultHttpStatus(request.Method)
	var args []reflect.Value
	args = append(args, reflect.ValueOf(ctx))

	r.ControllerMethod.Call(args)
	return ctx
}
