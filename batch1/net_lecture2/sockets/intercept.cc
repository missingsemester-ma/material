#include <pcap.h>
#include <iostream>

using namespace std;

void packet_handler(const u_char* packet, const struct pcap_pkthdr* header) {
    // Print the packet data in hexadecimal format
    for (int i = 0; i < header->len; i++) {
        printf("%02x ", packet[i]);
    }
    printf("\n");
}


int main() {
    char errbuf[PCAP_ERRBUF_SIZE];
    pcap_t* handle;
    const char* dev = "lo"; // Replace with your network interface name

    handle = pcap_open_live(dev, BUFSIZ, 1, 1000, errbuf);

    if (handle == nullptr) {
        fprintf(stderr, "Could not open device %s: %s\n", dev, errbuf);
        return 1;
    }

    struct pcap_pkthdr header;
    const u_char* packet;

    cout << "Waiting for packets" << endl;
    while (1) {
        packet = pcap_next(handle, &header);
        // Process the captured packet here
        packet_handler(packet, &header);
    }

    pcap_close(handle);
    return 0;
}
