package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	mb "github.com/letterj/oohhc/proto/account"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// PayLoad ...
type PayLoad struct {
	ID         string
	Name       string
	Token      string
	Status     string
	CreateDate int64
	DeleteDate int64
}

// AccountAPIServer is used to implement grpchello.CfsAdminApiServer
type AccountAPIServer struct {
	acctws *AccountWS
}

// NewAccountAPIServer ...
func NewAccountAPIServer(acctws *AccountWS) *AccountAPIServer {
	s := new(AccountAPIServer)
	s.acctws = acctws
	return s
}

// CreateAcct ...
func (s *AccountAPIServer) CreateAcct(ctx context.Context, r *mb.CreateAcctRequest) (*mb.CreateAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.CreateAcctResponse{Status: "ERROR: Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// validate new account does not exit
	//TODO:

	// create account information
	// Group:		/acct
	// Member:  "(uuid)"
	// Value:   { "id": "uuid", "name": "name", "apikey": "12345",
	//            "status": "active", "createdate": <timestamp>,
	//            "deletedate": <timestamp> }
	var p PayLoad
	g := "/acct"
	m := uuid.NewV4().String()
	// build payload
	p.ID = m
	p.Name = r.Acctname
	p.Token = uuid.NewV4().String()
	p.Status = "active"
	p.CreateDate = time.Now().Unix()
	p.DeleteDate = 0
	d, err := json.Marshal(p)
	if err != nil {
		status = fmt.Sprintf("ERROR: account %s was not created", r.Acctname)
		return &mb.CreateAcctResponse{Status: status}, err
	}
	// write information into the group store
	_, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		status = fmt.Sprintf("ERROR: account %s was not created", r.Acctname)
		return &mb.CreateAcctResponse{Status: status}, err
	}
	status = fmt.Sprintf("SUCCESS: account %s, was created with name %s", m, r.Acctname)
	return &mb.CreateAcctResponse{Status: status}, nil
}

// ListAcct ...
func (s *AccountAPIServer) ListAcct(ctx context.Context, r *mb.ListAcctRequest) (*mb.ListAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.ListAcctResponse{Payload: "", Status: "ERROR: Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// build the group store request
	g := "/acct"

	// try and get account details form the group store
	data, err := s.acctws.lookupGStore(g)
	if err != nil {
		status = fmt.Sprintf("ERROR: Problem looking up account group %s", "/acct")
		return &mb.ListAcctResponse{Payload: "", Status: status}, err
	}
	status = fmt.Sprintf("SUCCESS: Complete listing for group %s", g)
	return &mb.ListAcctResponse{Payload: data, Status: status}, nil
}

// ShowAcct ...
func (s *AccountAPIServer) ShowAcct(ctx context.Context, r *mb.ShowAcctRequest) (*mb.ShowAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.ShowAcctResponse{Payload: "", Status: "ERROR: Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string

	// build the group store request
	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	data, err := s.acctws.getGStore(g, m)
	if err != nil {
		status = fmt.Sprintf("ERROR: Problem looking up %s.", r.Acctnum)
		return &mb.ShowAcctResponse{Payload: "", Status: status}, err
	}
	if data == "" {
		status = fmt.Sprintf("ERROR: Account %s not found", r.Acctnum)
		return &mb.ShowAcctResponse{Payload: "", Status: status}, errors.New("Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(data), &p)
	if err != nil {
		status = fmt.Sprintf("ERROR: Details parsing error for %s", r.Acctnum)
		return &mb.ShowAcctResponse{Payload: "", Status: status}, err
	}
	status = fmt.Sprintf("SUCCESS: account %s was found with id %s", p.Name, r.Acctnum)
	return &mb.ShowAcctResponse{Payload: data, Status: status}, nil
}

// DeleteAcct ...
func (s *AccountAPIServer) DeleteAcct(ctx context.Context, r *mb.DeleteAcctRequest) (*mb.DeleteAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.DeleteAcctResponse{Status: "ERROR: Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	// get information from the group store
	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	result, err := s.acctws.getGStore(g, m)
	if err != nil {
		status = fmt.Sprintf("ERROR: Problem looking up %s.", r.Acctnum)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	if result == "" {
		status = fmt.Sprintf("ERROR: Account %s not found", r.Acctnum)
		return &mb.DeleteAcctResponse{Status: status}, errors.New("Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		status = fmt.Sprintf("ERROR: Details parsing error for %s", r.Acctnum)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	// only active accounts can be marked as deleted
	if p.Status != "active" || p.DeleteDate != 0 {
		status = fmt.Sprintf("ERROR: Account %s not in an active state. State: %s, Deleted Date: %d", r.Acctnum, p.Status, p.DeleteDate)
		return &mb.DeleteAcctResponse{Status: status}, errors.New("Incorrect account status.")
	}
	// send delete to the group store
	p.Status = "deleted"
	p.DeleteDate = time.Now().Unix()
	d, err := json.Marshal(p)
	if err != nil {
		status = fmt.Sprintf("ERROR: Account %s was not parsed", r.Acctnum)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	// write updated information into the group store
	result, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		status = fmt.Sprintf("ERROR: Account %s was not deleted", r.Acctnum)
		return &mb.DeleteAcctResponse{Status: status}, err
	}
	log.Println(result)
	status = fmt.Sprintf("SUCCESS: Account %s was deleted", r.Acctnum)
	return &mb.DeleteAcctResponse{Status: status}, nil
}

// UpdateAcct ...
func (s *AccountAPIServer) UpdateAcct(ctx context.Context, r *mb.UpdateAcctRequest) (*mb.UpdateAcctResponse, error) {
	var status string
	// validate superapikey
	if r.Superkey != s.acctws.superKey {
		return &mb.UpdateAcctResponse{Payload: "", Status: "ERROR: Invalid Credintials"}, errors.New("Permission Denied")
	}
	// validate account string
	g := "/acct"
	m := r.Acctnum

	// try and get account details form the group store
	result, err := s.acctws.getGStore(g, m)
	if err != nil {
		status = fmt.Sprintf("ERROR: Problem looking up %s.", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, err
	}
	if result == "" {
		status = fmt.Sprintf("ERROR: Account %s not found", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, errors.New("Not Found")
	}
	var p PayLoad
	err = json.Unmarshal([]byte(result), &p)
	if err != nil {
		status = fmt.Sprintf("ERROR: Details parsing error for %s", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, err
	}
	// update account information
	if r.ModAcct.Name != "" {
		p.Name = r.ModAcct.Name
	}
	if r.ModAcct.Status != "" {
		if p.Status == "deleted" && r.ModAcct.Status != "deleted" {
			p.DeleteDate = 0
		}
		p.Status = r.ModAcct.Status
	}
	if r.ModAcct.Token != "" {
		p.Token = r.ModAcct.Token
	}
	// write new information to the group store
	d, err := json.Marshal(p)
	if err != nil {
		status = fmt.Sprintf("ERROR: Account %s was not parsed", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, err
	}
	// write information into the group store
	_, err = s.acctws.writeGStore(g, m, d)
	if err != nil {
		status = fmt.Sprintf("ERROR: Account %s was not updated", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, err
	}
	// Pull updated data
	uresult, err := s.acctws.getGStore(g, m)
	if err != nil {
		status = fmt.Sprintf("ERROR: Problem looking up updated account %s.", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, err
	}
	if uresult == "" {
		status = fmt.Sprintf("ERROR: Updated account %s not found", r.Acctnum)
		return &mb.UpdateAcctResponse{Payload: "", Status: status}, errors.New("Server Error")
	}
	// Good request return
	status = fmt.Sprintf("SUCCESS: Account %s with id %s was updated", p.Name, r.Acctnum)
	return &mb.UpdateAcctResponse{Payload: uresult, Status: status}, nil
}
