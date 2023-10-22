#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define PORT 12345
#define MAX_MESSAGE_SIZE 1024

void error(const char *msg) {
    perror(msg);
    exit(1);
}

int main() {
    int sockfd, newsockfd;
    struct sockaddr_in server, client;
    socklen_t client_len = sizeof(client);
    char message[MAX_MESSAGE_SIZE];

    // Create a TCP socket
    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) {
        error("Error opening socket");
    }

    // Configure the server address structure
    server.sin_family = AF_INET;
    server.sin_addr.s_addr = INADDR_ANY;
    server.sin_port = htons(PORT);

    // Bind the socket to the server address
    if (bind(sockfd, (struct sockaddr *)&server, sizeof(server)) < 0) {
        error("Binding failed");
    }

    // Listen for incoming connections
    listen(sockfd, 5);

    printf("Waiting for a connection...\n");

    // Accept the incoming connection
    newsockfd = accept(sockfd, (struct sockaddr *)&client, &client_len);
    if (newsockfd < 0) {
        error("Error accepting connection");
    }

    printf("Connection established.\n");

    while (1) {
        // Receive a message from the client
        if (recv(newsockfd, message, MAX_MESSAGE_SIZE, 0) < 0) {
            error("Error receiving message");
        }
        printf("Client: %s", message);

        // Send a response back to the client
        printf("You: ");
        fgets(message, MAX_MESSAGE_SIZE, stdin);
        send(newsockfd, message, strlen(message), 0);
    }

    close(newsockfd);
    close(sockfd);
    return 0;
}
