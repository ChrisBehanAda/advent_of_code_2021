import math
import os
import functools


def get_input(path):
    fs_content = []
    with open(path) as file:
        for line in file:
            if not line:
                continue
            fs_content.append(line.rstrip("\n"))

    print(f"Lines ingested: {len(fs_content)}")  # Validate input file was read properly
    return fs_content


def convert_to_value(bin_string):
    sum = 0
    for idx in range(len(bin_string)):
        bit_flag = bin_string[len(bin_string) - 1 - idx]
        sum += (math.pow(2, idx)) * int(bit_flag)
    return int(sum)


def convert_to_binary(value):
    return bin(value)[2:]


def parse_literal(bin_string):
    packet = list(bin_string)
    packet_bin_string = ""
    while len(packet) > 0:
        packet_prefix = packet.pop(0)
        group = packet[0:4]
        packet_bin_string += "".join(group)
        packet = packet[4:]
        if packet_prefix == "0":
            break

    return packet_bin_string, packet


def parse_operator(bin_string):
    length_id = bin_string[0]
    agg_subpacket_version = 0

    if length_id == "0":
        # Next 15 bits are total length
        length = convert_to_value(bin_string[1:16])
        packets = bin_string[16 : (16 + length)]
        parsed_packets = []

        while packets != len(packets) > 0:  # Check if its just trailing 0s
            subpacket_version, packet_bin_string, packets = parse_bin_string(packets)
            if packet_bin_string:
                parsed_packets.append(packet_bin_string)
                agg_subpacket_version += subpacket_version
        return agg_subpacket_version, parsed_packets, bin_string[16 + length :]

    else:

        # Next 11 are number of packets
        num_packets = convert_to_value(bin_string[1:12])
        packets = bin_string[12:]
        parsed_packets = []
        for x in range(num_packets):
            subpacket_version, packet_bin_string, packets = parse_bin_string(packets)
            if packet_bin_string:
                parsed_packets.append(packet_bin_string)
                agg_subpacket_version += subpacket_version
        return agg_subpacket_version, parsed_packets, packets


def parse_bin_string(bin_string):
    version = convert_to_value(bin_string[:3])
    type = bin_string[3:6]

    if convert_to_value(type) == 4:
        packet_bin_string, packet = parse_literal(bin_string[6:])
        return version, packet_bin_string, packet

    else:
        agg_version, parsed_packets, remaining_bin_string = parse_operator(
            bin_string[6:]
        )
        aggregate = 0
        if convert_to_value(type) == 0:
            aggregate = functools.reduce(
                lambda x, y: convert_to_binary(
                    convert_to_value(x) + convert_to_value(y)
                ),
                parsed_packets,
            )
        elif convert_to_value(type) == 1:
            aggregate = functools.reduce(
                lambda x, y: convert_to_binary(
                    convert_to_value(x) * convert_to_value(y)
                ),
                parsed_packets,
            )
        elif convert_to_value(type) == 2:
            aggregate = convert_to_binary(
                min(list(map(lambda x: convert_to_value(x), parsed_packets)))
            )
        elif convert_to_value(type) == 3:
            aggregate = convert_to_binary(
                max(list(map(lambda x: convert_to_value(x), parsed_packets)))
            )
        elif convert_to_value(type) == 5:
            aggregate = convert_to_binary(
                1
                if convert_to_value(parsed_packets[0])
                > convert_to_value(parsed_packets[1])
                else 0
            )
        elif convert_to_value(type) == 6:
            aggregate = convert_to_binary(
                1
                if convert_to_value(parsed_packets[0])
                < convert_to_value(parsed_packets[1])
                else 0
            )
        else:
            aggregate = convert_to_binary(
                1 if parsed_packets[0] == parsed_packets[1] else 0
            )

        return version + agg_version, aggregate, remaining_bin_string


def get_solution(content):
    BITS_transmission = content[0]
    hex_map = {
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

    bin_string = ""
    for char in BITS_transmission:
        bin_string += hex_map[char]

    version, agg_value, *args = parse_bin_string(bin_string)

    return convert_to_value(agg_value)


if __name__ == "__main__":
    test_solution = get_solution(get_input(path=os.getcwd() + "/day16/input.txt"))
    if test_solution == 1:
        print("TEST PASSED")
        answer = 1673210814091
        if get_solution(get_input(path=os.getcwd() + "/input.txt")) == answer:
            print("Still works")

    else:
        print(test_solution)
        print("Test failed")
