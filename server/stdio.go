package server

import (
	"bufio"
	"errors"
	"io"
	"jargonlsp/protocol"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO maybe store documents, maybe not??
type LanguageServer struct {
	Reader *bufio.Reader // TODO maybe depends on tcp maybe not
	Writer *bufio.Writer // TODO maybe depends on tcp maybe not
}

func New() *LanguageServer {
	log.Println("Creating new language server")

	// TODO handle tcp input
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	server := LanguageServer{
		Reader: reader,
		Writer: writer,
	}

	return &server
}

func (s *LanguageServer) Run() {
	for {
		// TODO dont do anythign with writing responses here !!!
		// TODO just log error if it happens, ignore and continue running
		err := s.processRequest()
		if err != nil {
			if err == io.EOF {
				log.Println("EOF reached")
				return
			} else {
				log.Fatalf("ERROR: %v", err)
			}
		}
	}
}

func (s *LanguageServer) processRequest() error {
	log.Println("call process request")

	const headerSeparator uint8 = '\n'

	var contentLength int = 0

	// read the header

	for {
		line, err := s.Reader.ReadString(headerSeparator)
		if err != nil {
			return err
		}

		log.Println("GOT:")
		log.Println(line)
		log.Println("END GOT")

		// TODO ovo mogu maknut skroz
		header := strings.TrimSpace(line)
		if len(header) == 0 {
			log.Println("sad je ovaj prazan")
			break
		}

		headerParts := strings.Split(strings.TrimSpace(header), ":")
		if len(headerParts) != 2 {
			log.Println("smth is fucked with header parts:", headerParts)
			return errors.New("header parts not two wtf")
		}
		log.Println(headerParts)

		// TODO verify header parts 0

		contentLength, err = strconv.Atoi(strings.TrimSpace(headerParts[1]))
		if err != nil {
			return err
		}

		log.Printf("Got size: %d", contentLength)

		if contentLength == 0 {
			return errors.New("wtf how is content len zero")
		}
	}

	// read the content

	content := make([]byte, contentLength)
	var totalReadBytes int = 0

	for {
		readBytes, err := s.Reader.Read(content[totalReadBytes:])
		if err != nil {
			return err
		}

		log.Printf("READ: %d EXPECTED: %d", readBytes, contentLength)

		log.Println("CONTENT START:")
		log.Println(string(content))
		log.Println("CONTENT END:")

		if readBytes == 0 {
			break
		}

		totalReadBytes += readBytes
		if totalReadBytes == contentLength {
			break
		}
	}

	//
	// log.Println("ready for parse", string(content))
	//
	// log.Println("parsing json")
	// log.Println(string(content))
	// message, err := protocol.DecodeMessage(content)
	// if err != nil {
	// 	return err
	// }
	// log.Println(*message)
	// log.Println("parsing json done")
	//
	// TODO pass everything to protocol handler - based on method

	// TODO wrap entire run in  go

	response, err := protocol.HandleClientRequest(content)
	if err != nil {
		log.Println(err)
	}

	// TOOD write response
	log.Println("RESPONSE")
	log.Println(response)

	if response == nil {
		log.Println("empty response - nothing to write")
		return nil
	}

	// TODO send back to server with some internal io component

	wroteBytes, err := s.Writer.Write(response)
	if err != nil {
		return err
	}

	log.Println("WROTE:")
	log.Println(wroteBytes)

	err = s.Writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
