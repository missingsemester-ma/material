#include <iostream>
#include <cstdlib>
#include <cstring>
#include <arpa/inet.h>
#include <unistd.h>

using namespace std;

int main() {
  // Create a socket
  int clientSocket = socket(AF_INET, SOCK_STREAM, 0);
  if (clientSocket == -1) {
      cerr << "Error creating socket" << endl;
      return -1;
  }

  // Connect to the server
  sockaddr_in serverAddr;
  serverAddr.sin_family = AF_INET;
  serverAddr.sin_port = htons(3030);
  serverAddr.sin_addr.s_addr = inet_addr("127.0.0.1"); 

  if (connect(clientSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr)) == -1) {
      cerr << "Error connecting" << endl;
      return -1;
  }

  int numbers[2];
  cin >> numbers[0] >> numbers[1];
  send(clientSocket, numbers, sizeof(numbers), 0);

  int sum;
  recv(clientSocket, &sum, sizeof(sum), 0);
  cout << "Sum received from server: " << sum << endl;

  close(clientSocket);

  return 0;
}

