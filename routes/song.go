package routes

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"

  "github.com/YaminLi/Jukebox_test/models"

)

var idAssignmentChan = make(chan string)

func addSongRoutes(r *mux.Router) {
  r.HandleFunc("/songs", getSongsHandler).Methods("GET")
  r.HandleFunc("/songs", createSongHandler).Methods("POST")
  r.HandleFunc("/songs/{song_id}/upvote", upvoteSongHandler).Methods("POST") // TODO
  r.HandleFunc("/songs/{song_id}", deleteSongHandler).Methods("DELETE") // TODO
}

func getSongsHandler(w http.ResponseWriter, req *http.Request) {
  // Get songs
  var songs []models.Song
  models.DB.Find(&songs)

  // Create response
  w.Header().Set("Content-Type", "application/json")

  data := make(map[string]interface{})
  data["songs"] = songs
  json.NewEncoder(w).Encode(&data)
}

func createSongHandler(w http.ResponseWriter, req *http.Request) {

  // Get JSON request data
  decoder := json.NewDecoder(req.Body)
  var song models.Song
  err := decoder.Decode(&song)
  if err != nil {
    panic(err)
    return
  }

  // Initialize votes value
  song.Votes = 1

  // Save song
  models.DB.Create(&song)

  // Create response
  json.NewEncoder(w).Encode(&song)
}

// TODO
func deleteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  // Get song by id
  // var songs []models.Song
  // models.DB.Find(&songs)
  // var song Song
  // for elem := range songs{
  //   if elem.id == song_id{
  //       song = elem
  //       break
  //   }
  // }
  //
  // // Delete song
  // models.DB.Delete(&song)   //????????????
  //
  // // Return response
  // json.NewEncoder(w).Encode(&song)
}

// TODO
func upvoteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  // // Get song by id specified
  // var songs []models.Song
  // models.DB.Find(&songs)
  // var song Song
  // for elem := range songs{
  //   if elem.id == song_id{
  //       song = elem
  //       break
  //   }
  // }
  //
  // // +1 to score
  // song.Votes += 1
  //
  // // Save song
  // models.DB.Save(&song)
  // // models.DB.Update(&song) //?????????
  //
  // // Return song in response
  // json.NewEncoder(w).Encode(&song)
}
