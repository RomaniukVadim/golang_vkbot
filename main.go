package main

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"test/vkapi"
	"time"
)

var (
	BOT_ID = "271741813" // ID бота (Panda Do-Do)
	api    = &vkapi.Api{}
)

func main() {
	api.AccessToken = " PUT YOUR TOKEN HERE"
	prefixes := []string{"Afina ", "Афина ", "бот ", "Afina, ", "Афина, ", "бот, ", "!"}
	p := make(map[string]string)
	p["count"] = "1"
	p["out"] = "0"
	last_msg := 0
	is_pause := 0
	for {
		m := api.Request("messages.get", p)

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
						random_photo("https://vk.com/album232317814_232070447", "Лови 3D няшку")
					case e + "2Dняшку", e + "няшку":
						random_photo("https://vk.com/album232317814_231917660", "Лови 2D няшку")
					case e + "мику", e + "хатсуне", e + "хатсуне мику":
						random_photo("https://vk.com/album232317814_231918384", "Лови Хатсуне Мику")
					case e + "мегумин", e + "megumin":
						random_photo("https://vk.com/album-54385020_174984625", "Лови Мегумин")
					case e + "shiro", e + "широ", e + "waifu", e + "вайфу":
						random_photo("vk.com/album232317814_229487221", "Лови Широ")
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
	k, v := GetUid()
	api.Request("messages.send", map[string]string{k: v, "message": msg, "attachment": img})
}

func random_photo(url string, message string) {
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
	k, v := GetUid()
	api.Request("messages.send", map[string]string{"message": message + " №" + strconv.Itoa(offset), k: v, "attachment": "photo" + user_id + "_" + photo_id_string})
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func GetUid() (string, string) {
	p := make(map[string]string)
	p["count"] = "1"
	p["out"] = "0"
	m := api.Request("messages.get", p)
	//uid
	uid_regexp, _ := regexp.Compile("\"uid\":([0-9]+),\"read")
	uid := uid_regexp.FindStringSubmatch(m)[1]
	//uid
	//chat_id

	title := vkapi.GetResponse(m, "title")
	if title != " ... " {
		chat_id_regexp, _ := regexp.Compile("\"chat_id\":([0-9]+)")
		chat_id := chat_id_regexp.FindStringSubmatch(m)[1]
		return "chat_id", chat_id

	} else {
		return "user_id", uid
	}

}

func getRandUser() string {
	params := make(map[string]string)
	k, v := GetUid()
	params[k] = v

	chatUsers := api.Request("messages.getChatUsers", params)

	users_regexp, _ := regexp.Compile("([0-9]+)")
	users := users_regexp.FindAllStringSubmatch(chatUsers, 50)
	return users[rand.Intn(len(users))][0]
}
