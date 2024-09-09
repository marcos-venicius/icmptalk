# ICMPTalk

Chat with someone via ICMP.

./demos/demo-01.mp4

### Example

Listen for someone to connect to you!

```bash
go run . --iface 10.0.0.x --listen
```

Connect to someone that is listening:

```bash
go run . --iface 10.0.0.x --target 10.0.0.x
```

### Using

Just write something and press enter!

### The handshake

Who is listening expects the sender to send 4 packets in that order:

1. A number
2. A number
3. A number
4. The sum of all three previous numbers

**The sum of the three first numbers should be odd.**
**These numbers should be sent in the following format "|number|"**

Then, if everything is correct, the listener sends the (4th number times 2)

Then, the sender checks if the number sent is equal to the sum of three previous number times two.

If this requirements is filled, then the sender confirms the handshake sending an "|OK|".
If it fails, the sender sends "|FAIL|"

### Disclaimer

1. **This is not "production ready"!**
2. **All the traffic is completely clean!**
3. **I'v tested this only in local network (between two machines), but should work on external too**
4. **The handshake is not intended to be secure nor some kind of authentication, it's intended to avoid some random ping to turn into a chat automatically**

### Missing features

- Detect when one of the connected users is disconnected, then disconnect both and close the app
- Encrypt messages between peers
- Avoid random ping to be interpreted as a message
