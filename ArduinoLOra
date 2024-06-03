#include <SPI.h>
#include <LoRa.h>
#include "max6675.h" // max6675.h file is part of the library that you should download from Robojax.com

const int sensorPin = 7; // Pino digital ao qual o sensor de vibração está conectado
int contagemZero = 0; // Variável para contar quantas vezes o valor 0 foi detectado
int contagemUm = 0; // Variável para contar quantas vezes o valor 1 foi detectado
int contadorAmostraVib = 0;
String resultadoLeitura = "";
int sensorValue = 0;

int pinoSensorCorrente = A1;
int sensorValue_aux = 0;
float valorSensor = 0;
float valorCorrente = 0;
float voltsporUnidade = 0.004887586; // 5%1023
float sensibilidade = 0.066; // Sensibilidade do sensor ACS712 de 5A

int pinSensorSom = A0;
const int amostrasPorSegundo = 60;
int contadorAmostrasSom = 0;
long somaAmostrasSom = 0;
float mediaAmostrasSom = 0.0;

// Variáveis para cálculo da média da temperatura
float somaTemperaturas = 0.0;
int contadorAmostrasTemperatura = 0;
float mediaAmostrasTemperatura = 0.0;
unsigned long ultimoTempo = 0; // Marca o tempo da última exibição
const unsigned long intervaloExibicao = 60000;
bool imprimirResultado = false;

int soPin = 4; // SO=Serial Out
int csPin = 5; // CS = chip select CS pin
int sckPin = 6; // SCK = Serial Clock pin

MAX6675 robojax(sckPin, csPin, soPin); // create instance object of MAX6675

void setup() {
  Serial.begin(9600); // initialize serial monitor with 9600 baud
  Serial.println("PI VII - Monitoramento Motor Monofásico");
  pinMode(pinoSensorCorrente, INPUT);
  LoRa.begin(433E6); // Inicia o LoRa na frequência 433 MHz
  LoRa.setSpreadingFactor(7);
  LoRa.setSignalBandwidth(125E3);
  LoRa.setCodingRate4(5);
  LoRa.setSyncWord(0xF3); // Palavra de sincronização
}

void loop() {
  // Realiza 60 medições e armazena os valores lidos
  for (int i = 0; i < 60; i++) {
    contadorAmostraVib++;
    sensorValue = digitalRead(sensorPin);
    if ((sensorValue) == HIGH) {
      contagemUm++;
    } else {
      contagemZero++;
    }
  }

  // Verifica qual valor foi mais frequente
  if (contagemUm > contagemZero) {
    resultadoLeitura = "Está Vibrando!";
  } else {
    resultadoLeitura = "Não Está Vibranado!";
  }

  for (int i = 60000; i > 0; i--) {
    sensorValue_aux = (analogRead(pinoSensorCorrente) - 510);
    valorSensor += pow(sensorValue_aux, 2);
    delay(1);
  }

  valorSensor = (sqrt(valorSensor / 10000)) * voltsporUnidade;
  valorCorrente = (valorSensor / sensibilidade);

  valorSensor = 0;

  // Leitura do sensor de som
  int valorSensorSom = analogRead(pinSensorSom);
  somaAmostrasSom += valorSensorSom;
  contadorAmostrasSom++;

  // Leitura do sensor de temperatura
  float temperaturaCelsius = robojax.readCelsius();
  somaTemperaturas += temperaturaCelsius;
  contadorAmostrasTemperatura++;

  if (contadorAmostrasSom >= amostrasPorSegundo && contadorAmostrasTemperatura >= amostrasPorSegundo && contadorAmostraVib >= amostrasPorSegundo) {
    mediaAmostrasSom = somaAmostrasSom / (float)amostrasPorSegundo;
    mediaAmostrasTemperatura = somaTemperaturas / (float)amostrasPorSegundo;

    // Envia os dados via LoRa
    LoRa.beginPacket();
    LoRa.print("Média Som: "); LoRa.print(mediaAmostrasSom);
    LoRa.print(" db Média Temperatura: "); LoRa.print(mediaAmostrasTemperatura);
    LoRa.print(" C° Média Corrente: "); LoRa.print(valorCorrente, 3);
    LoRa.print(" A Vibração: "); LoRa.println(resultadoLeitura);
    LoRa.endPacket();

    // Reinicia os contadores e acumuladores
    contagemZero = 0;
    contagemUm = 0;
    contadorAmostrasSom = 0;
    somaAmostrasSom = 0;
    contadorAmostrasTemperatura = 0;
    somaTemperaturas = 0.0;

    // Aguarda 60 segundos antes de coletar novamente
    delay(intervaloExibicao);
  }

  delay(1000 / amostrasPorSegundo);
}
