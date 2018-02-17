package requester

type EndpointHandler interface {
	Get(url string)
	//Post(url string)
}
