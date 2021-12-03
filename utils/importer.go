package utils

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSessionKey() string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("AOC_SESSION_KEY")
}

func ImportData(day int) (string, error) {
	fileName := fmt.Sprintf("data/day%d.txt", day)
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File already exists, using file %s\n", fileName)
		return fileName, nil
	}

	u := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)
	urlObj, _ := url.ParseRequestURI(u)
	jar, _ := cookiejar.New(&cookiejar.Options{})
	sessionKey := getSessionKey()
	if sessionKey == "" {
		panic(fmt.Sprintf("No file for day %d present and "+
			"no session key present in env variables to load one."+
			"\n Copy file yourself from here %s", day, u))
	}
	cookie := http.Cookie{
		Name:   "session",
		Value:  sessionKey,
		MaxAge: 360,
	}
	jar.SetCookies(urlObj, []*http.Cookie{&cookie})
	client := &http.Client{Jar: jar}

	resp, err := client.Get(u)
	check(err)
	defer resp.Body.Close()

	println(resp.Status)
	if resp.StatusCode != 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		check(err)
		return "", fmt.Errorf("error loading dataset: %s", string(bodyBytes))
	}
	f, err := os.Create(fileName)
	check(err)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		_, err := fmt.Fprintln(f, scanner.Text())
		check(err)
	}
	return fileName, nil
}

func PrepareData(day int) *[]string {
	fileName, err := ImportData(day)
	if err != nil {
		panic(err)
	}
	return ReadLines(fileName)
}
