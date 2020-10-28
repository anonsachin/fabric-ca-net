package rpc

import(
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"addorg/proto/server"
	"addorg/generate"
	"log"
	"context"
)

// NewOrg the grpc server implementation
type NewOrg struct{
	log *log.Logger
	server.UnimplementedNewOrgServer
}

func NewOrgServer(log *log.Logger) *NewOrg{
	return &NewOrg{log: log}
}

func (n *NewOrg) CreateMSP(ctx context.Context, m *server.Msp) (*server.Void, error){
	if m.GetOutDir() == "" || m.GetOrg() == "" || m.GetVaultHost() == ""{
		return nil, status.Errorf(codes.FailedPrecondition,"The arguments are not proper %#v",m)
	} 
	msp := generate.NewMSP(m.GetVaultHost(),m.GetOrg(),m.GetOutDir())

	err := msp.CreateMSP()

	if err != nil{
		return nil, status.Errorf(codes.Internal,"The internal error of %s", err.Error())
	}

	return &server.Void{}, nil
}