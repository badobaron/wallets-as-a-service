# Gliederung - Inhalt

## 1. Einführung
- Was sind Kryptowährungen
- Wissenschaftlicher Dreiklang
- ?Was ist Blockchain?

### Motivation
- Keine großen Finanzunternehmen hinter den Walletentwicklern
- Wenn nicht open source und von Hand compiliert, kann Schadcode in den Code gelangen, der sensible Daten an Dritte weitergibt
- Nutzer verlieren täglich Geld in Wallets, weil sie ihren privaten Schlüssel vergessen

## 2. Vorgehensweise
- Analyse herkömmlicher Lösungen (Bankkonten + Cryptocoin Wallets)
- Befragungen von Probanden
- Transformation passender Eigenschaften der Analogien auf Crypto Wallets As a Service

## 3. Herkömmliche Bankkonten - Ist Zustand
- Einfache Übersicht der Konten in App
- App Zugriff per Passwort
- Transaktionen per TAN

## 4. Crypto Wallets - Ist Zustand
### 4.1 Lokale Cryptocoin Wallets
- Lokale Wallets werden durch privaten Schlüssel gesichert
- private Schlüssel sehr langer Hashwert
- privater Schlüssel nicht einfach zu merken
- Bei Verlust des Schlüssels ist ein finanzieller Verlust unausweichlich
- Transaktionen per Passwort (kann einfacher sein)
- Kein Nachlassprozess
- GWG wird erst eingehalten bei Umtausch von Cryptowährungen in Fiat-Währungen

### 4.2 Online Cryptocoin Wallets
- Zugriff per Passwort
- Kein Nachlassprozess
- Kein privater Schlüssel
- Bei Hardforks der Währungen muss auf die Entscheidungen der Online Dienstleister vertraut werden
- Bei Hardforks nur eine Währung statt mehrere

## 5. Konzept Cryptocoin Wallets As A Service

### 5.1 Benutzersicht Anforderungen
- Einfacher Zugriff, einfache und schnelle Transaktionen
- Muss sich nur Passwort merken
- Auch bei Passwortverlust kein finanzielles Risiko
- Möchte sich für Cryptocoins nicht umgewöhnen
- Möchte gesicherten Nachlass
- Möchte Wallet als App und Website
- Möchte gesicherte verschlüsselte Kommunikation zu Dienstleister

### 5.2 Dienstleistersicht Anforderungen
- Speichert private Schlüssel der Wallets
- Kann neue Passwörter mittels PK für Besitzer generieren
- Einzelne Personen keinen Zugriff auf private Wallets
- Wallets per PK gesichert
- Möchte Service absichern
- Keine PK und Passwörter in Klartext speichern

## 6. Implementierung

### 6.1 Client

### 6.2 Server

## 7. Evaluation

## 8. Ausblich

## 9. Zusammenfassung
