#include <SPI.h>
#include <LoRa.h>
#include "max6675.h"
#include <MPU6050_tockn.h>
#include <Wire.h>

// Definições do sensor de temperatura
int soPin = 4;
int csPin = 5;
int sckPin = 6;

MAX6675 robojax(sckPin, csPin, soPin);

// Definições do sensor de corrente
int pinoSensor = A1;
float sensibilidade = 0.066;
float voltsporUnidade = 0.004887586;

// Definições do sensor MPU6050
MPU6050 mpu6050(Wire);

// Definições do sensor de som
int pinSensorSom = A0;

void setup() {
  Serial.begin(9600);
  Wire.begin();
  mpu6050.begin();
  mpu6050.calcGyroOffsets(true);

  pinMode(pinoSensor, INPUT);
  LoRa.begin(433E6);
  LoRa.setSpreadingFactor(7);
  LoRa.setSignalBandwidth(125E3);
  LoRa.setCodingRate4(5);
  LoRa.setSyncWord(0xF3);
}

void loop() {
  float temperaturaC = robojax.readCelsius();
  float temperaturaF = robojax.readFahrenheit();
  
  int sensorValue = analogRead(pinoSensor);
  float corrente = (sensorValue * voltsporUnidade - 2.5) / sensibilidade;

  mpu6050.update();
  float acelX = mpu6050.getAccX();
  float acelY = mpu6050.getAccY();
  float acelZ = mpu6050.getAccZ();

  LoRa.beginPacket();
  LoRa.print("TempC: "); LoRa.print(temperaturaC);
  LoRa.print(" TempF: "); LoRa.print(temperaturaF);
  LoRa.print(" Corrente: "); LoRa.print(corrente);
  LoRa.print(" AcelX: "); LoRa.print(acelX);
  LoRa.print(" AcelY: "); LoRa.print(acelY);
  LoRa.print(" AcelZ: "); LoRa.println(acelZ);
  LoRa.endPacket();

  delay(1500);  
}
