package server

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sskender/jargonlsp/protocol"
	"github.com/sskender/jargonlsp/state"
	"github.com/sskender/jargonlsp/version"
)

type ServerSettings struct {
	DictionaryPath string
	EnableTcp      bool
}

type Server struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func New(settings ServerSettings) *Server {

	// TODO handle TCP instance

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	server := Server{
		Reader: reader,
		Writer: writer,
	}

	log.Printf("Starting %s %s", server.Name(), server.Version())

	gdb := state.GetDatabase()

	err := gdb.Dictionary.Load(settings.DictionaryPath)
	if err != nil {
		log.Printf("warning: loading dictionary failed - dictionary is empty: %v", err)
	}

	return &server
}

func (s *Server) Name() string {
	return version.Name
}

func (s *Server) Version() string {
	return version.Version
}

func (s *Server) RunLoop() {

	// TODO run async

	for {
		err := s.processRequest()
		if err != nil {
			if err == io.EOF {
				// TODO call exit command
				return
			}
			log.Fatalf("ERROR: %v", err)
		}
	}
}

func (s *Server) processRequest() error {
	log.Println("waiting for a new request")

	content, err := s.readRequest()
	if err != nil {
		return err
	}

	if content == nil {
		return errors.New("received content is empty")
	}

	log.Printf("Content: %s", string(content))

	response, err := protocol.HandleClientRequest(content)
	if err != nil {
		return err
	}

	err = s.writeResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) readRequest() ([]byte, error) {

	var contentLength int = 0

	// read the header

	for {
		const headerSeparator uint8 = '\n'

		headerLine, err := s.Reader.ReadString(headerSeparator)
		if err != nil {
			return nil, err
		}

		header := strings.TrimSpace(headerLine)
		if len(header) == 0 {
			break
		}

		headerParts := strings.Split(header, ":")
		if len(headerParts) != 2 {
			return nil, errors.New("header is in invalid format")
		}

		contentLength, err = strconv.Atoi(strings.TrimSpace(headerParts[1]))
		if err != nil {
			return nil, err
		}

		log.Printf("Content length: %d", contentLength)

		if contentLength == 0 {
			return nil, errors.New("content length should not be zero")
		}
	}

	// read the content

	content := make([]byte, contentLength)

	var totalReadBytes int = 0

	for {
		readBytes, err := s.Reader.Read(content[totalReadBytes:])
		if err != nil {
			return nil, err
		}

		if readBytes == 0 {
			break
		}

		totalReadBytes += readBytes
		if totalReadBytes == contentLength {
			break
		}
	}

	return content, nil
}

func (s *Server) writeResponse(response []byte) error {
	if response == nil {
		return nil
	}

	_, err := s.Writer.Write(response)
	if err != nil {
		return err
	}

	err = s.Writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
