# quic-c2

---

## Features

* Persistent QUIC connections to multiple devices (`quic-go`).
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
**This project is for _educational purposes only_.**  
Do **not** use this code to access, attack, or interfere with systems without explicit, written permission.  
Use only in isolated lab environments — **use at your own risk**; authors disclaim liability.
