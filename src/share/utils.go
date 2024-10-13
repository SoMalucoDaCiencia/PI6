package share

import (
	"PI6/models"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func GetTimestamps(year, month, day int) (time.Time, time.Time) {
	start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.Month(month), day, 23, 59, 59, 999999999, time.UTC)
	return start, end
}

func FloatsAsUUID(f1, f2 float64) string {
	var vec []byte
	vec = binary.BigEndian.AppendUint64(vec, math.Float64bits(f1))
	vec = binary.BigEndian.AppendUint64(vec, math.Float64bits(f2))
	return uuid.NewMD5(uuid.NameSpaceDNS, vec[:]).String()
}

func FloatsFromUUID(uuid []byte) (float64, float64) {
	f1 := math.Float64frombits(binary.BigEndian.Uint64(uuid[:8]))
	f2 := math.Float64frombits(binary.BigEndian.Uint64(uuid[8:16]))
	return f1, f2
}

func GetRandomProxy(proxies []models.ProxyObj) *http.Client {
	rand.Seed(time.Now().UnixNano())
	selectedProxy := proxies[rand.Intn(len(proxies))]

	var transport *http.Transport
	urll, err := url.Parse(fmt.Sprintf("http://%s:%d", selectedProxy.Ip, selectedProxy.Port))
	if err != nil {
		panic(err)
	}

	//switch selectedProxy.Protocol { // Seleciona o primeiro protocolo
	//case "http":
	transport = &http.Transport{
		Proxy: http.ProxyURL(urll),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//case "socks4", "socks5":
	//    dialer, _ := proxy.SOCKS5("tcp", selectedProxy.Proxy, nil, proxy.Direct)
	//    transport = &http.Transport{
	//        DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
	//            return dialer.Dial("tcp", addr)
	//        },
	//    }
	//}

	return &http.Client{
		Transport: transport,
	}
}
