package rtsp

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"tailscale.com/net/interfaces"
)

// RequestHandler callback function that gets invoked when a request is received
type RequestHandler func(req *Request, resp *Response, localAddr string, remoteAddr string)

// Server Server for handling Rtsp control requests
type Server struct {
	port     int
	handlers map[Method]RequestHandler
	done     chan bool
	reqChan  chan *Request
	ip       string
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
	// Get the default outbound interface address
	_, myIP, ok := interfaces.LikelyHomeRouterIP()
	if !ok {
		log.Errorf("Error getting local outbound IP address: ok=%v\n", ok)
		return
	}
	r.ip = myIP.String()
	log.Printf("Starting RTSP server on address: %s:%d\n", r.ip, r.port)

	tcpListen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", r.ip, r.port))
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

	requestDumpFile, err := os.OpenFile(".handshakedata-requests", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Panicf("Error opening request dumpfile: %v\n", err)
	}
	defer requestDumpFile.Close()
	requestDumpFile.Truncate(0)
	requestDumpFile.Seek(0, 0)

	responseDumpFile, err := os.OpenFile(".handshakedata-responses", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Panicf("Error opening response dumpfile: %v\n", err)
	}
	defer responseDumpFile.Close()
	responseDumpFile.Truncate(0)
	responseDumpFile.Seek(0, 0)

	combinedFile, err := os.OpenFile(".handshakedata-combined", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Panicf("Error opening combined handshake dumpfile: %v\n", err)
	}
	defer combinedFile.Close()
	combinedFile.Truncate(0)
	combinedFile.Seek(0, 0)

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

		_, _ = requestDumpFile.Write(buildRequestDump(request))

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

		_, _ = responseDumpFile.Write(buildResponseDump(resp))
		_, _ = combinedFile.Write(buildRequestDump(request))
		_, _ = combinedFile.Write(buildResponseDump(resp))

		_, _ = writeResponse(conn, resp)
	}
}

func buildRequestDump(req *Request) []byte {
	buf := new(bytes.Buffer)
	buf.WriteString("--------BEGIN REQUEST DUMP------\n")
	buf.WriteString(fmt.Sprintf("Request URI: %s\n", req.RequestURI))
	buf.WriteString(fmt.Sprintf("Request Method: %s\n", req.Method.String()))
	buf.WriteString(fmt.Sprintf("Headers: %v\n", req.Headers))
	buf.WriteString(fmt.Sprintf("Protocol: %s\n", req.protocol))
	buf.WriteString(fmt.Sprintf("Body: %s\n", string(req.Body)))
	buf.WriteString("------END REQUEST DUMP------\n\n")
	return buf.Bytes()
}

func buildResponseDump(resp *Response) []byte {
	buf := new(bytes.Buffer)
	buf.WriteString("------BEGIN RESPONSE DUMP------\n")
	buf.WriteString(fmt.Sprintf("Status: %s\n", resp.Status.String()))
	buf.WriteString(fmt.Sprintf("Headers: %v\n", resp.Headers))
	buf.WriteString(fmt.Sprintf("Protocol: %s\n", resp.protocol))
	buf.WriteString(fmt.Sprintf("Body: %s\n ", string(resp.Body)))
	buf.WriteString("------END RESPONSE DUMP------\n\n")
	return buf.Bytes()
}
