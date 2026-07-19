package collab

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SessionSync struct {
	SessionID   string        `json:"session_id"`
	SessionName string        `json:"session_name"`
	Messages    []SyncMessage `json:"messages"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type SyncMessage struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Reasoning string    `json:"reasoning,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type Server struct {
	addr        string
	mu          sync.RWMutex
	current     *SessionSync
	subscribers map[chan SessionSync]struct{}
	httpSrv     *http.Server
}

func NewServer(addr string) *Server {
	if addr == "" {
		addr = ":9876"
	}
	return &Server{
		addr:        addr,
		subscribers: make(map[chan SessionSync]struct{}),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/sync", s.handleSync)
	mux.HandleFunc("/subscribe", s.handleSubscribe)
	mux.HandleFunc("/push", s.handlePush)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	s.httpSrv = &http.Server{Addr: s.addr, Handler: mux}
	return s.httpSrv.ListenAndServe()
}

func (s *Server) Stop() error {
	if s.httpSrv != nil {
		return s.httpSrv.Close()
	}
	return nil
}

func (s *Server) Push(sync *SessionSync) {
	s.mu.Lock()
	s.current = sync
	s.current.UpdatedAt = time.Now()
	subs := make([]chan SessionSync, 0, len(s.subscribers))
	for ch := range s.subscribers {
		subs = append(subs, ch)
	}
	s.mu.Unlock()

	for _, ch := range subs {
		select {
		case ch <- *sync:
		default:
		}
	}
}

func (s *Server) handleSync(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.current == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "no session"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.current)
}

func (s *Server) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	ch := make(chan SessionSync, 4)
	s.mu.Lock()
	s.subscribers[ch] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.subscribers, ch)
		s.mu.Unlock()
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Send initial state
	if s.current != nil {
		data, _ := json.Marshal(s.current)
		w.Write([]byte("data: " + string(data) + "\n\n"))
		flusher.Flush()
	}

	for {
		select {
		case sync := <-ch:
			data, _ := json.Marshal(sync)
			w.Write([]byte("data: " + string(data) + "\n\n"))
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func (s *Server) handlePush(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	var sync SessionSync
	if err := json.NewDecoder(r.Body).Decode(&sync); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.Push(&sync)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type Client struct {
	serverAddr string
	httpClient http.Client
}

func NewClient(addr string) *Client {
	if addr == "" {
		addr = "http://localhost:9876"
	}
	return &Client{
		serverAddr: addr,
		httpClient: http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) Push(sync *SessionSync) error {
	body, err := json.Marshal(sync)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Post(c.serverAddr+"/push", "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) Get() (*SessionSync, error) {
	resp, err := c.httpClient.Get(c.serverAddr + "/sync")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var sync SessionSync
	if err := json.NewDecoder(resp.Body).Decode(&sync); err != nil {
		return nil, err
	}
	return &sync, nil
}

func (c *Client) Subscribe() (<-chan SessionSync, error) {
	resp, err := c.httpClient.Get(c.serverAddr + "/subscribe")
	if err != nil {
		return nil, err
	}
	ch := make(chan SessionSync, 4)
	go func() {
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		for {
			var sync SessionSync
			if err := dec.Decode(&sync); err != nil {
				close(ch)
				return
			}
			ch <- sync
		}
	}()
	return ch, nil
}
