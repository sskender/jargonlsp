package server

import (
	"bufio"
	"errors"
	"io"
	"jargonlsp/protocol"
	"jargonlsp/protocol/base"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanguageServer struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func New() *LanguageServer {

	// TODO handle tcp input

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	server := LanguageServer{
		Reader: reader,
		Writer: writer,
	}

	log.Printf("%s v%s", server.Name(), server.Version())

	return &server
}

func (s *LanguageServer) Name() string {
	return base.LSP_SERVER_NAME
}

func (s *LanguageServer) Version() string {
	return base.LSP_SERVER_VERSION
}

func (s *LanguageServer) Run() {

	// TODO run async

	for {
		err := s.processRequest()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("ERROR: %v", err)
		}
	}
}

func (s *LanguageServer) processRequest() error {
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

func (s *LanguageServer) readRequest() ([]byte, error) {

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

func (s *LanguageServer) writeResponse(response []byte) error {
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
