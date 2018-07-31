package muxhandler

import ("time"
	"net/http"
		"encoding/json"
		"github.com/gorilla/mux"
	"strconv"
	"fmt"
)





type Topic struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type TopicHandler struct {
	topicCacheMap map[int]Topic
}

func NewTopicHandler() (*TopicHandler){
	return &TopicHandler{topicCacheMap: map[int]Topic{}}
}

func(t *TopicHandler) HandleAdd(w http.ResponseWriter, req *http.Request) {
	topic := Topic{}
	if err := json.NewDecoder(req.Body).Decode(&topic); err != nil {
		http.Error(w,"Decode json error", 200)
	}
	t.topicCacheMap[topic.Id] = topic
	fmt.Println(t.topicCacheMap)
	w.Write([]byte("HandleAdd Success."))
}

func(t *TopicHandler) HandleGet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		http.Error(w, "id must be integer", 200)
	}
	topic := t.topicCacheMap[id]
	//fmt.Println(t.topicCacheMap)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	topicBytes, _ := json.Marshal(topic)
	//fmt.Println(string(topicBytes))
	w.Write(topicBytes)
	//json.NewEncoder(w).Encode(topic)
}

func(t * TopicHandler) HandleGetAll(w http.ResponseWriter, req *http.Request){
	var topicSlice []Topic
	for _, value := range t.topicCacheMap{
		topicSlice = append(topicSlice, value)
	}
	json.NewEncoder(w).Encode(topicSlice)
}

func(t *TopicHandler) HandleModify(w http.ResponseWriter, req *http.Request) {
	topic := Topic{}
	if err := json.NewDecoder(req.Body).Decode(&topic); err != nil {
		http.Error(w, "Decode json error", 200)
	}
	t.topicCacheMap[topic.Id] = topic
	fmt.Println(t.topicCacheMap)
	w.Write([]byte("HandleModify Success."))
}

func(t *TopicHandler) HandleDelete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		http.Error(w, "id must be integer", 200)
	}
	delete(t.topicCacheMap, id)
	fmt.Println(t.topicCacheMap)
	w.Write([]byte("HandleDelete Success."))
}