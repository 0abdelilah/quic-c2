Here’s a clean, professional README for your project **“quic-c2”**:

---

# quic-c2



---

## Features

* Persistent QUIC connections to multiple devices.
* Terminal-based UI using **TUI** (`tview`).
* Send commands to selected devices in real-time.
* Continuous output streaming from each device.
* Thread-safe device management.

---

## Installation

### Start the server
```bash
git clone https://github.com/0abdelilah/quic-c2.git
cd quic-c2/server/
go run .
```

2. Connect a client.
```bash
cd quic-c2/client/
go run .
```


3. Use the TUI to select a device and enter commands:

* Navigate devices with arrow keys.
* Press `Enter` to select a device.
* Type commands in the input field and press `Enter` to send.
* Press `Left Arrow` to return to the device list.
* Press `Ctrl+C` to exit.

---

## Project Structure

```
quic-c2/
├─ models/          # Device and manager structs
├─ server/          # QUIC server
├─ tui/             # Terminal UI and command handling
```

---

## Contributing
Open for contributions: anyone can add new features or add more client-side commands.

⚠️ **Disclaimer:**  
This project is for **educational purposes only**. It is **not intended for production use**.  
The code may contain **incomplete, untested, or unsafe implementations**.  
Use it at your own risk and **do not deploy it in a live environment**.  
