package main

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"log"
	"time"
)

const LinkLength = 7

func hashMD5(link, clientIp string) []byte {
	timestamp := time.Now().String()
	h := md5.New()
	io.WriteString(h, clientIp)
	io.WriteString(h, timestamp)
	io.WriteString(h, link)
	return h.Sum(nil)
}

func hashBase64(md5Hash []byte) string {
	return base64.URLEncoding.EncodeToString(md5Hash)
}

func calcHash(link, clientIp string) string {
	md5Hash := hashMD5(clientIp, link)
	base64Hash := hashBase64(md5Hash)
	return base64Hash[:LinkLength]
}

type LinkService struct {
	config         Config
	linkRepository LinkRepository
}

func (us *LinkService) ShortenLink(link, clientIp string) string {
	hash := calcHash(link, clientIp)
	shortLink := us.config.RemoteAddr + "/" + hash
	us.linkRepository.SaveLink(hash, link)
	return shortLink
}

func (us *LinkService) FindLink(hash string) string {
	link, err := us.linkRepository.FindLink(hash)
	if err != nil {
		log.Fatalln(err)
	}
	return link
}
