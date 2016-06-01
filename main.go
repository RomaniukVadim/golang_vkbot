package main

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"webapp/vkapi"
)

var (
	BOT_ID = "271741813" // ID бота (Panda Do-Do)
	api    = &vkapi.Api{}
)

func main() {
	api.AccessToken = "Put your token here"
	prefixes := []string{"Afina ", "Афина ", "бот ", "Afina, ", "Афина, ", "бот, ", "!"}
	p := make(map[string]string)
	p["count"] = "1"
	p["out"] = "0"
	last_msg := 0
	is_pause := 0
	for {
		m := api.Request("messages.get", p)
		//uid
		uid_regexp, _ := regexp.Compile("\"uid\":([0-9]+),\"read")
		uid := uid_regexp.FindStringSubmatch(m)[1]
		//uid

		//chat_id
		// chat_id_regexp, _ := regexp.Compile("\"chat_id\":([0-9]+)")
		// chat_id := chat_id_regexp.FindStringSubmatch(m)[1]
		// if chat_id != "" {
		// 	uid = chat_id
		// }
		//chat_id

		body := "no"

		if len((vkapi.GetResponse(m, "body").(string))) != 0 {
			body = (vkapi.GetResponse(m, "body").(string))
		}
		mid := int((vkapi.GetResponse(m, "mid")).(float64))
		help_message := `=== Бот Афина ===
		       Доступные комманды:
		      ` + "&#10145;" + `Афина, помощь 
		      ` + "&#10145;" + `Афина, привет
		      ` + "&#10145;" + `Афина, мику
		      ` + "&#10145;" + `Афина, 2Dняшку
		      ` + "&#10145;" + `Афина, 3Dняшку
		      ` + "&#10145;" + `Афина, мегумин
		      ` + "&#10145;" + `Афина, широ
		      ` + "&#10145;" + `Афина, создатель
		       Author: Meryborn(https://vk.com/meryborn), 2016`
		if mid > last_msg && body != "no" {
			if is_pause == 0 {
				for _, e := range prefixes {
					switch body {
					//
					//Commands
					//

					case e + "help", e + "помощь", e + "старт":
						send(help_message, "photo232317814_403576726")
					case e + "привет", e + "Привет":
						send("Здравствуй, Семпай", "photo232317814_403576725")
					case e + "создатель":
						send("Meryborn(https://vk.com/meryborn)", "photo232317814_403576721")
					case e + "3Dняшку":
						random_photo("https://vk.com/album232317814_232070447", uid, "Лови 3D няшку")
					case e + "2Dняшку", e + "няшку":
						random_photo("https://vk.com/album232317814_231917660", uid, "Лови 2D няшку")
					case e + "мику", e + "хатсуне", e + "хатсуне мику":
						random_photo("https://vk.com/album232317814_231918384", uid, "Лови Хатсуне Мику")
					case e + "мегумин", e + "megumin":
						random_photo("https://vk.com/album-54385020_174984625", uid, "Лови Мегумин")
					case e + "shiro", e + "широ", e + "waifu", e + "вайфу":
						random_photo("vk.com/album232317814_229487221", uid, "Лови Широ")
						//
						//Commands
						//
					}
					last_msg = mid
				}
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}

}

func send(msg string, img string) {

	response := api.Request("messages.get", map[string]string{"count": "1", "out": "0"})
	uid_regexp, _ := regexp.Compile("\"uid\":([0-9]+),\"read")
	uid := uid_regexp.FindStringSubmatch(response)[1]
	api.Request("messages.send", map[string]string{"user_id": uid, "message": msg, "attachment": img})
}

func random_photo(url string, to_user string, message string) {
	album_id := strings.Split((strings.Split(url, "/a")[1]), "_")[1]
	user_id := strings.Replace((strings.Split((strings.Split(url, "/a")[1]), "_")[0]), "lbum", "", 1)
	photos_count := api.Request("photos.getAlbums", map[string]string{"owner_id": user_id, "album_ids": album_id})
	size_regexp, _ := regexp.Compile("\"size\":([0-9]+)")
	size, _ := strconv.Atoi(size_regexp.FindStringSubmatch(photos_count)[1])
	offset := randInt(0, size)
	p_id := api.Request("photos.get", map[string]string{"owner_id": user_id, "album_id": album_id, "offset": strconv.Itoa(offset)})
	pid_regexp, _ := regexp.Compile("\"pid\":([0-9]+)")
	photo_id_string := pid_regexp.FindStringSubmatch(p_id)[1]
	//photo_id_int, _ := strconv.Atoi(photo_id_string)
	api.Request("messages.send", map[string]string{"user_id": to_user, "message": message + " №" + strconv.Itoa(offset), "attachment": "photo" + user_id + "_" + photo_id_string})
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
