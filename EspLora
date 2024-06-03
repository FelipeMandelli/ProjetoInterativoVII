#include <SPI.h>
#include <LoRa.h>
#include <WiFi.h>
#include <HTTPClient.h>

const char* ssid = "SSID_nome_da_rede_wifi"; // SSID da rede WiFi
const char* password = "senha_da_rede_wifi"; // Senha da rede WiFi

void setup() {
  Serial.begin(9600);
  LoRa.begin(433E6); // Inicia o LoRa na frequência 433 MHz
  LoRa.setSpreadingFactor(7);
  LoRa.setSignalBandwidth(125E3);
  LoRa.setCodingRate4(5);
  LoRa.setSyncWord(0xF3); // Palavra de sincronização

  // Conecta ao WiFi
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.println("Conectando ao WiFi...");
  }
  Serial.println("Conectado ao WiFi com sucesso");
}

void enviarDadosWeb(String dados) {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    http.begin("http://endereco_da_api.com/receber_dados"); // URL da API
    http.addHeader("Content-Type", "application/x-www-form-urlencoded");

    int httpResponseCode = http.POST("dados=" + dados);

    if (httpResponseCode == 200) {
      String response = http.getString();
      Serial.println("Dados enviados com sucesso: " + response);
    } else {
      Serial.println("Erro ao enviar dados: " + String(httpResponseCode));
    }

    http.end();
  } else {
    Serial.println("Erro: Não conectado ao WiFi");
  }
}

void loop() {
  int tamPacote = LoRa.parsePacket();
  if (tamPacote) {
    String leitura = "";
    while (LoRa.available()) {
      leitura += (char)LoRa.read();
    }
    Serial.print("Status: ");
    Serial.println(leitura);

    // Envia os dados para a web
    enviarDadosWeb(leitura);
  }
}
