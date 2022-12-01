package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Packet struct {
	version     int
	packet_type int
	literal     int
	sub_packets []Packet
}

func (p Packet) String() string {
	if len(p.sub_packets) > 0 {
		operator_str_map := map[int]string{
			0: "+",
			1: "+",
			2: "min",
			3: "max",
			5: "<",
			6: ">",
			7: "==",
		}
		return fmt.Sprintf("(%s: %s)", operator_str_map[p.packet_type], p.sub_packets)
	}
	return fmt.Sprintf("%d", p.literal)
}

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

func hex_to_bin_str(in string) string {
	out := ""
	converter := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	for _, c := range in {
		out += converter[string(c)]
	}

	return out
}

func read_header(stream string) (int, int, string) {
	version64, _ := strconv.ParseInt(stream[0:3], 2, 64)
	stream = stream[3:]
	type64, _ := strconv.ParseInt(stream[0:3], 2, 64)
	stream = stream[3:]
	return int(version64), int(type64), stream
}

func read_literal(stream string) (int, string) {
	literal_bin := ""
	for true {
		stopper := stream[0:1]
		stream = stream[1:]
		segment := stream[0:4]
		stream = stream[4:]
		literal_bin += segment
		if stopper == "0" {
			break
		}
	}
	literal32, _ := strconv.ParseInt(literal_bin, 2, 32)
	return int(literal32), stream
}

func read_operator_length(stream string) (int, int, string) {
	length_type, _ := strconv.ParseInt(stream[0:1], 2, 64)
	stream = stream[1:]
	var length64 int64
	switch length_type {
	case 0:
		length64, _ = strconv.ParseInt(stream[0:15], 2, 64)
		stream = stream[15:]
	case 1:
		length64, _ = strconv.ParseInt(stream[0:11], 2, 64)
		stream = stream[11:]
	}
	return int(length_type), int(length64), stream
}

func parse_stream(stream string, max int) ([]Packet, string) {
	values := []Packet{}
	all_zeros_re := regexp.MustCompile("^0*$")
	count := 0
	for !all_zeros_re.MatchString(stream) && (max == 0 || count < max) {
		var p Packet
		p.version, p.packet_type, stream = read_header(stream)
		switch p.packet_type {
		case 4:
			var value int
			value, stream = read_literal(stream)
			p.literal = value
		default:
			var length_type, length int
			length_type, length, stream = read_operator_length(stream)
			var values []Packet
			switch length_type {
			case 0:
				values, _ = parse_stream(stream[0:length], 0)
				stream = stream[length:]
			case 1:
				values, stream = parse_stream(stream, length)
			}
			p.sub_packets = values
		}
		values = append(values, p)
		count++
	}

	return values, stream
}

func sum_version(in []Packet) int {
	sum := 0
	for _, p := range in {
		sum += p.version
		if len(p.sub_packets) > 0 {
			sum += sum_version(p.sub_packets)
		}
	}
	return sum
}

func calculate(p Packet) int {
	result := 0
	switch p.packet_type {
	case 0:
		for _, sp := range p.sub_packets {
			result += calculate(sp)
		}
	case 1:
		result = 1
		for _, sp := range p.sub_packets {
			result *= calculate(sp)
		}
	case 2:
		result = -1
		for _, sp := range p.sub_packets {
			value := calculate(sp)
			if result == -1 || value < result {
				result = value
			}
		}
	case 3:
		for _, sp := range p.sub_packets {
			value := calculate(sp)
			if value > result {
				result = value
			}
		}
	case 4:
		result = p.literal
	case 5:
		valueA := calculate(p.sub_packets[0])
		valueB := calculate(p.sub_packets[1])
		if valueA > valueB {
			result = 1
		} else {
			result = 0
		}
	case 6:
		valueA := calculate(p.sub_packets[0])
		valueB := calculate(p.sub_packets[1])
		if valueA < valueB {
			result = 1
		} else {
			result = 0
		}
	case 7:
		valueA := calculate(p.sub_packets[0])
		valueB := calculate(p.sub_packets[1])
		if valueA == valueB {
			result = 1
		} else {
			result = 0
		}
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)
	stream := hex_to_bin_str(lines[0])
	fmt.Println(stream)

	packets, stream := parse_stream(stream, 0)
	fmt.Println(packets)
	fmt.Printf("Version total %d\n", sum_version(packets))
	fmt.Printf("Result %d\n", calculate(packets[0]))
}
