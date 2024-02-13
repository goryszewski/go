package internal_client

type Autok8sClient struct {
	URL   string
	TOKEN string
}

type Ip struct {
	Private string
	Public  string
}

type Worker struct {
	Name string
	IP   Ip
	Type string
}

type LB struct {
	Id      int
	Ip      string
	service string
}
