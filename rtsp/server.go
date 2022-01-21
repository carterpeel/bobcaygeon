package rtsp

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

// RequestHandler callback function that gets invoked when a request is received
type RequestHandler func(req *Request, resp *Response, localAddr string, remoteAddr string)

// Server Server for handling Rtsp control requests
type Server struct {
	port     int
	handlers map[Method]RequestHandler
	done     chan bool
	reqChan  chan *Request
}

// NewServer instantiates a new RtspServer
func NewServer(port int) *Server {
	return &Server{
		port:     port,
		done:     make(chan bool),
		handlers: make(map[Method]RequestHandler),
		reqChan:  make(chan *Request),
	}
}

// AddHandler registers a handler for a given RTSP method
func (r *Server) AddHandler(m Method, rh RequestHandler) {
	r.handlers[m] = rh
}

// Stop stops the RTSP server
func (r *Server) Stop() {
	log.Println("Stopping RTSP server")
	r.done <- true
}

// Start creates listening socket for the RTSP connection
func (r *Server) Start(verbose bool) {
	log.Printf("Starting RTSP server on port: %d\n", r.port)
	tcpListen, err := net.Listen("tcp", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Errorln("Error listening:", err.Error())
		return
	}

	defer tcpListen.Close()

	// Handle TCP connections.
	go func() {
		for {
			// Listen for an incoming connection.
			conn, err := tcpListen.Accept()
			if err != nil {
				log.Errorln("Error accepting: ", err.Error())
				return
			}
			go r.read(conn, r.handlers, verbose)
		}
	}()

	<-r.done
}

func (r *Server) read(conn net.Conn, handlers map[Method]RequestHandler, verbose bool) {
	defer conn.Close()
	for {
		request, err := readRequest(conn)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Client closed connection")
			} else {
				log.Println("Error reading data: ", err.Error())
			}
			return
		}

		if verbose {
			log.Println("Received Request")
			log.Println(request.String())
		}

		handler, exists := handlers[request.Method]
		if !exists {
			log.Printf("Method: %s does not have a handler. Skipping", request.Method)
			continue
		}
		resp := NewResponse()
		// for now we just stick in the protocol (protocol/version) from the request
		resp.protocol = request.protocol
		// same with CSeq
		resp.Headers["CSeq"] = request.Headers["CSeq"]
		// invokes the client specified handler to build the response
		localAddr := conn.LocalAddr().(*net.TCPAddr).IP
		remoteAddr := conn.RemoteAddr().(*net.TCPAddr).IP
		handler(request, resp, localAddr.String(), remoteAddr.String())
		if verbose {
			log.Println("Outbound Response")
			log.Println(resp.String())
		}
		_, _ = writeResponse(conn, resp)
	}
}
