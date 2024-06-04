#include <SPI.h>
#include <LoRa.h>
#include <WiFi.h>
#include <HTTPClient.h>

const char* ssid = "SSID_nome_da_rede_wifi";
const char* password = "senha_da_rede_wifi";

void setup() {
  Serial.begin(9600);
  LoRa.begin(433E6);
  LoRa.setSpreadingFactor(7);
  LoRa.setSignalBandwidth(125E3);
  LoRa.setCodingRate4(5);
  LoRa.setSyncWord(0xF3);

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
    http.begin("http://endereco_da_api.com/receber_dados");
    http.addHeader("Content-Type", "application/x-www-form-urlencoded");

    int httpResponseCode = http.POST("dados=" + dados);

    if (httpResponseCode > 0) {
      String response = http.getString();
      Serial.println("Dados enviados com sucesso: " + response);
    } else {
      Serial.println("Falha na conexão: " + String(httpResponseCode));
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
    Serial.print("Dados recebidos: ");
    Serial.println(leitura);

    enviarDadosWeb(leitura);
  }
}
