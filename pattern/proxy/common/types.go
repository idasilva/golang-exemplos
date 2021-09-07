package common


type server interface {
	HandlerRequest(string, string) (int,string)
}
