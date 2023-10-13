#include <iostream>
#include <cstdlib>
#include <cstring>
#include <arpa/inet.h>
#include <unistd.h>

using namespace std;

void handleCLient(int clientSocket) {
  int numbers[2];

  if (recv(clientSocket, numbers, sizeof(numbers), 0) == -1) {
    cerr << "Error receiving data from client" << endl;
    close(clientSocket);
    return ;
  }
  
  int sum = numbers[0] + numbers[1];

  if (send(clientSocket, &sum, sizeof(sum), 0) == -1) {
    cerr << "Error sending data to client" << endl;
  }

  close(clientSocket);
}

int main() {
  // Create a socket
  int serverSocket = socket(AF_INET, SOCK_STREAM, 0);
  if (serverSocket == -1) {
    cerr << "Error while creating socket" << endl;
    return -1;
  }

  sleep(10);
  
  sockaddr_in serverAddr;
  serverAddr.sin_family = AF_INET;
  serverAddr.sin_port = htons(3030);
  serverAddr.sin_addr.s_addr = INADDR_ANY;

  if (bind(serverSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr)) == -1) {
    cerr << "Error binding" << endl;
    return -1;
  }

  sleep(10);

  if (listen(serverSocket, 1) == -1) {
    cerr << "Error listening" << endl;
    return -1;
  }

  sleep(10);

  cout << "Server is listening at port 3030" << endl; 

  while (true) {
    sockaddr_in clientAddr;
    socklen_t clientAddrSize = sizeof(clientAddr);
    int clientSocket = accept(serverSocket, (struct sockaddr*)&clientAddr, &clientAddrSize);
    if (clientSocket == -1) {
      cerr << "Error accepting connection" << endl;
      continue;
    }

    int numbers[2];
    recv(clientSocket, numbers, sizeof(numbers), 0);
    sleep(2);
    int sum = numbers[0] + numbers[1];
    send(clientSocket, &sum, sizeof(sum), 0);
    sleep(2);

    close(clientSocket);
  }

  close(serverSocket);

  return 0;
}
