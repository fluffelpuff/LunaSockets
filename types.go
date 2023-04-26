package LunaSockets

type FirstCheck struct {
	Id   string
	Type string
}

type RpcRequest struct {
	ProxyPass *HeaderData   `json:"proxypass"`
	Params    []interface{} `json:"params"`
	JSONRPC   string        `json:"jsonrpc"`
	Method    string        `json:"method"`
	ID        int           `json:"id"`
}

type RpcResponse struct {
	Error  *string
	Result interface{}
}

type IoFlowPackage struct {
	Id   string
	Type string
	Body string
}

type HeaderData struct {
	Accept         string `header:"Accept"`
	AcceptCharset  string `header:"Accept-Charset"`
	AcceptEncoding string `header:"Accept-Encoding"`
	AcceptLanguage string `header:"Accept-Language"`
	Authorization  string `header:"Authorization"`
	CacheControl   string `header:"Cache-Control"`
	Connection     string `header:"Connection"`
	ContentLength  int64  `header:"Content-Length"`
	ContentType    string `header:"Content-Type"`
	Cookie         string `header:"Cookie"`
	Host           string `header:"Host"`
	Referer        string `header:"Referer"`
	UserAgent      string `header:"User-Agent"`
	SourcePort     int64  `header:"Source-Port"`
	SourceIp       string `header:"Source-Ip"`
}

type Request struct {
	ProxyPass     bool
	Header        *HeaderData
	OutPassedArgs []interface{}
}

type LunaRpcResponseFunction func(response RpcResponse)
type LunaServicesMapList map[string]interface{}
type LunaPingResponseFunction func()
