package sn

import (
	"bytes"
	"net"
	"sync"
)

type SNClient interface {
	AddrString() string
}

type Client struct {
	sync.RWMutex
	ClientId         string
	Conn             *net.UDPConn
	Address          *net.UDPAddr
	registeredTopics map[uint16]string
	pendingMessages  map[uint16]*PublishMessage
}

func NewClient(ClientId string, Conn *net.UDPConn, Address *net.UDPAddr) *Client {
	INFO.Printf("NewClient, id: \"%s\"\n", ClientId)
	return &Client{
		sync.RWMutex{},
		ClientId,
		Conn,
		Address,
		make(map[uint16]string),
		make(map[uint16]*PublishMessage),
	}
}

func (c *Client) Write(m Message) error {
	var buf bytes.Buffer
	m.Write(&buf)
	_, e := c.Conn.WriteToUDP(buf.Bytes(), c.Address)
	return e
}

func (c *Client) Register(topicId uint16, topic string) {
	defer c.Unlock()
	c.Lock()
	INFO.Printf("client %s registered topicId %d\n", c.ClientId, topicId)
	c.registeredTopics[topicId] = topic
}

func (c *Client) Registered(topicId uint16) bool {
	defer c.RUnlock()
	c.RLock()
	_, ok := c.registeredTopics[topicId]
	return ok
}

func (c *Client) AddPendingMessage(p *PublishMessage) {
	defer c.Unlock()
	c.Lock()
	c.pendingMessages[p.TopicId] = p
}

func (c *Client) FetchPendingMessage(topicId uint16) *PublishMessage {
	defer c.Unlock()
	c.Lock()
	pm := c.pendingMessages[topicId]
	delete(c.pendingMessages, topicId)
	return pm
}

func (c *Client) AddrString() string {
	return c.Address.String()
}

func (c *Client) String() string {
	return c.ClientId
}

type Clients struct {
	sync.RWMutex
	// indexed by "address:port" => StorableClient
	clients map[string]SNClient
}

func (c *Clients) GetClient(addr *net.UDPAddr) SNClient {
	defer c.RUnlock()
	c.RLock()
	return c.clients[addr.String()]
}

// Return true if this is a new client, false otherwise
// Clients are indexed by their address:port b/c
// that's the only indentifying information we have
// outside of a CONNECT packet
func (c *Clients) AddClient(client SNClient) bool {
	defer c.Unlock()
	c.Lock()
	addr := client.AddrString()
	INFO.Println("AddClient(%s - %s)", client, addr)
	isNew := false
	if c.clients[addr] == nil {
		isNew = true
	}
	//todo: what to do if clientid is in use?
	//     is there some cleanup involved in topictree?
	c.clients[addr] = client
	return isNew
}

func (c *Clients) RemoveClient(id string) {
	defer c.Unlock()
	c.Lock()
	INFO.Println("RemoveClient(%s)", id)
	delete(c.clients, id)
}
